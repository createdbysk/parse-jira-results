package operator

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/sheets/v4"
)

// JWTConfigFactory is the type for the function that returns a *jwt.Config.
type JWTConfigFactory func(credentials []byte, scope ...string) (*jwt.Config, error)

// HttpClientFactoryGetter is the type for a function thta returns a function that returns a *http.Client
type HttpClientFactoryGetter func(config *jwt.Config) HttpClientFactory

// HttpClientFactory is the type for a function that returns *http.Client.
type HttpClientFactory func(ctx context.Context) *http.Client

// SheetsServiceFactory is the type for a function that creates the sheets.Service.
type SheetsServiceFactory func(client *http.Client) (*sheets.Service, error)

// GoogleContext is the a context abstraction for Google APIs.
type GoogleContext struct {
	ConfigFactory        JWTConfigFactory
	GetHttpClientFactory HttpClientFactoryGetter
	ServiceFactory       SheetsServiceFactory
	Context              context.Context
}

func NewGoogleContext() *GoogleContext {
	return &GoogleContext{
		ConfigFactory:        google.JWTConfigFromJSON,
		GetHttpClientFactory: getHttpClientFactory,
		ServiceFactory:       sheets.New,
		Context:              context.Background(),
	}
}

type connection struct {
	srv *sheets.Service
}

func NewGoogleSheetsConnection(googleCtx *GoogleContext, credentials []byte, scope ...string) (Connection, error) {
	config, err := googleCtx.ConfigFactory(credentials, scope...)
	if err != nil {
		return nil, err
	}
	httpClientFactory := googleCtx.GetHttpClientFactory(config)
	httpClient := httpClientFactory(googleCtx.Context)
	srv, err := googleCtx.ServiceFactory(httpClient)
	if err != nil {
		return nil, err
	}
	c := &connection{srv}
	return c, nil
}

func (c *connection) Get(impl interface{}) {
	*(impl.(**sheets.Service)) = c.srv
}

func getHttpClientFactory(config *jwt.Config) HttpClientFactory {
	return config.Client
}

type googleSheetOutput struct {
	spreadsheetId string
	sheetTitle    string
	columnIndex   int64
	rowIndex      int64
}

func NewGoogleSheetOutput(spreadsheetId string, cellRef string) Output {
	pattern := `(\w+)[!](\w)(\d+)`
	re := regexp.MustCompile(pattern)
	submatches := re.FindAllStringSubmatch(cellRef, -1)[0]
	sheetTitle := submatches[1]
	columnIndex := int(submatches[2][0]) - int('A')
	rowIndex, _ := strconv.Atoi(submatches[3])
	rowIndex -= 1
	return &googleSheetOutput{
		spreadsheetId: spreadsheetId,
		sheetTitle:    sheetTitle,
		rowIndex:      int64(rowIndex),
		columnIndex:   int64(columnIndex),
	}
}

func (o *googleSheetOutput) Write(c Connection, data interface{}) error {
	var srv *sheets.Service
	c.Get(&srv)
	spreadsheetsGetCall := srv.Spreadsheets.Get(o.spreadsheetId)
	spreadsheet, err := spreadsheetsGetCall.Context(context.Background()).Do()
	if err != nil {
		return err
	}
	var sheetId int64
	sheetId = -1
	for _, s := range spreadsheet.Sheets {
		if s.Properties.Title == o.sheetTitle {
			sheetId = s.Properties.SheetId
			break
		}
	}
	if sheetId == -1 {
		return fmt.Errorf("sheet %v not found", o.sheetTitle)
	}
	gridCoordinate := &sheets.GridCoordinate{
		SheetId:     sheetId,
		RowIndex:    o.rowIndex,
		ColumnIndex: o.columnIndex,
	}
	d := data.(string)
	pasteDataRequest := sheets.PasteDataRequest{
		Coordinate: gridCoordinate,
		Data:       d,
		Delimiter:  "|",
		Type:       "PASTE_NORMAL",
	}
	request := sheets.Request{
		PasteData: &pasteDataRequest,
	}

	batchUpdateSpreadsheetRequest := sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			&request,
		},
	}
	call := srv.Spreadsheets.BatchUpdate(
		o.spreadsheetId,
		&batchUpdateSpreadsheetRequest,
	)
	_, err = call.Context(context.Background()).Do()
	if err != nil {
		return err
	}
	return nil
}
