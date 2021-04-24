package operator

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"google.golang.org/api/sheets/v4"
	"local.dev/sheetsLoader/internal/config"
)

type connection struct {
	srv *sheets.Service
}

func NewGoogleSheetsConnection(googleCtx *config.GoogleContext, credentialsFilePath string, scope ...string) (Connection, error) {
	optionWithCredentialsFile := googleCtx.OptionWithCredentialsFileFactory(credentialsFilePath)
	optionWithScopes := googleCtx.OptionWithScopesFactory(scope...)
	srv, err := googleCtx.SheetsServiceFactory(googleCtx.Context, optionWithCredentialsFile, optionWithScopes)
	if err != nil {
		return nil, err
	}
	c := &connection{srv}
	return c, nil
}

func (c *connection) Get(impl interface{}) {
	*(impl.(**sheets.Service)) = c.srv
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

func (o *googleSheetOutput) Write(c Connection, it Iterator) error {
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
	var data string
	it.Next(&data)
	pasteDataRequest := sheets.PasteDataRequest{
		Coordinate: gridCoordinate,
		Data:       data,
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
