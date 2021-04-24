package operator

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type mockFactory struct {
	credentialsFilePath           string
	scope                         []string
	option                        []option.ClientOption
	ctx                           context.Context
	mockSheetsService             *sheets.Service
	mockOptionWithCredentialsFile option.ClientOption
	mockOptionWithScopes          option.ClientOption
	createServiceError            error
}

func (c *mockFactory) createOptionWithCredentialsFile(credentialsFilePath string) option.ClientOption {
	c.credentialsFilePath = credentialsFilePath
	return c.mockOptionWithCredentialsFile
}

func (c *mockFactory) createOptionWithScopes(scope ...string) option.ClientOption {
	c.scope = scope
	return c.mockOptionWithScopes
}

func (c *mockFactory) createService(ctx context.Context, option ...option.ClientOption) (*sheets.Service, error) {
	c.ctx = ctx
	c.option = option
	return c.mockSheetsService, c.createServiceError
}

func TestNewGoogleSheetsConnection(t *testing.T) {
	// GIVEN
	credentialsFilePath := "/path/to/creds"
	scope := []string{
		"http://www.testscope.com/test1",
		"http://www.testscope.com/test2",
	}
	sheetsService := &sheets.Service{}
	ctx := context.TODO()
	factory := &mockFactory{
		mockOptionWithCredentialsFile: option.WithCredentialsFile(credentialsFilePath),
		mockOptionWithScopes:          option.WithScopes(scope...),
		mockSheetsService:             sheetsService,
	}
	context := GoogleContext{
		OptionWithCredentialsFileFactory: factory.createOptionWithCredentialsFile,
		OptionWithScopesFactory:          factory.createOptionWithScopes,
		SheetsServiceFactory:             factory.createService,
		Context:                          ctx,
	}

	expected := map[string]interface{}{
		"credentialsFilePath": credentialsFilePath,
		"scope":               scope,
		"ctx":                 ctx,
		"sheets.Service":      sheetsService,
	}

	// WHEN
	var srv *sheets.Service
	cn, err := NewGoogleSheetsConnection(&context, credentialsFilePath, scope...)
	cn.Get(&srv)
	if err != nil {
		t.Fatalf("TestGoogleDataAccess: Error %v", err)
	}
	actual := map[string]interface{}{
		"credentialsFilePath": factory.credentialsFilePath,
		"scope":               factory.scope,
		"ctx":                 factory.ctx,
		"sheets.Service":      srv,
	}
	// THEN
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"TestGoogleDataAccess: expected %v, actual %v",
			expected,
			actual,
		)
	}
}

// func TestGoogleSheetConnectionFailures(t *testing.T) {
// 	// GIVEN
// 	testcases := []struct {
// 		testcase           string
// 		createConfigError  error
// 		createServiceError error
// 	}{
// 		{
// 			testcase:          "createConfigError",
// 			createConfigError: errors.New("create config failure"),
// 		},
// 		{
// 			testcase:           "createServiceError",
// 			createServiceError: errors.New("create service failure"),
// 		},
// 	}

// 	for _, tt := range testcases {
// 		t.Run(
// 			tt.testcase,
// 			func(t *testing.T) {
// 				// GIVEN
// 				credentialsJSON := `{"fake": "Credentials"}`
// 				credentials := bytes.NewBufferString(credentialsJSON).Bytes()
// 				scope := []string{
// 					"http://www.testscope.com/test1",
// 					"http://www.testscope.com/test2",
// 				}
// 				jwtConfig := jwt.Config{}
// 				httpClient := &http.Client{}
// 				sheetsService := &sheets.Service{}
// 				ctx := context.TODO()
// 				factory := &mockFactory{
// 					mockJWTConfig:      &jwtConfig,
// 					mockHttpClient:     httpClient,
// 					mockSheetsService:  sheetsService,
// 					createConfigError:  tt.createConfigError,
// 					createServiceError: tt.createServiceError,
// 				}
// 				context := GoogleContext{
// 					ConfigFactory:        factory.createConfig,
// 					GetHttpClientFactory: factory.getHttpClientFactory,
// 					ServiceFactory:       factory.createService,
// 					Context:              ctx,
// 				}

// 				// WHEN
// 				_, err := NewGoogleSheetsConnection(&context, credentials, scope...)

// 				// THEN
// 				if err == nil {
// 					t.Errorf("%v: Expected error. None returned.", tt.testcase)
// 				}
// 			},
// 		)
// 	}
// }

func spreadsheetIdFixture() string {
	return "spreadsheetId"
}

func badSpreadsheetIdFixture() string {
	return "badSpreadsheetId"
}

func sheetTitleFixture(index int64) string {
	return fmt.Sprintf("TAB%d", index)
}

func badSheetTitleFixture() string {
	return fmt.Sprintf("ImBad")
}

func sheetIdFixture(index int64) int64 {
	return index
}

func dataFixture() string {
	return "Hello|World"
}

func iteratorFixture() Iterator {
	it := &mockIterator{data: dataFixture()}
	return it
}

func sheetPropertiesFixture(index int64) *sheets.SheetProperties {
	properties := sheets.SheetProperties{
		Title:   sheetTitleFixture(index),
		SheetId: sheetIdFixture(index),
	}
	return &properties
}

func sheetsFixture() []*sheets.Sheet {
	return []*sheets.Sheet{
		&sheets.Sheet{
			Properties: sheetPropertiesFixture(1),
		},
		&sheets.Sheet{
			Properties: sheetPropertiesFixture(2),
		},
	}
}

func spreadsheetFixture() *sheets.Spreadsheet {
	shts := sheetsFixture()
	spreadsheet := &sheets.Spreadsheet{
		Sheets: shts,
	}
	return spreadsheet
}

func gridCoordinateFixture(index int64) *sheets.GridCoordinate {
	sheetId := sheetIdFixture(index)
	return &sheets.GridCoordinate{
		SheetId:     sheetId,
		ColumnIndex: 0,
		RowIndex:    0,
	}
}

type mockSpreadsheetsHandler struct{}

func (h *mockSpreadsheetsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// https://sheets.googleapis.com/v4/spreadsheets/{spreadsheetId}
	vars := mux.Vars(r)
	spreadsheetId := vars["spreadsheetId"]
	if spreadsheetId == spreadsheetIdFixture() {
		spreadsheet := spreadsheetFixture()
		json.NewEncoder(w).Encode(spreadsheet)
	}
}

type mockBatchUpdateHandler struct {
	err     error
	request sheets.BatchUpdateSpreadsheetRequest
}

func (h *mockBatchUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// If force error, then just return an error.
	if h.err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		// https://sheets.googleapis.com/v4/spreadsheets/{spreadsheetId}:batchUpda
		vars := mux.Vars(r)
		spreadsheetId := vars["spreadsheetId"]
		if spreadsheetId == spreadsheetIdFixture() {
			err := json.NewDecoder(r.Body).Decode(&h.request)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				batchUpdateSpreadsheetResponse := sheets.BatchUpdateSpreadsheetResponse{
					SpreadsheetId: spreadsheetId,
				}
				json.NewEncoder(w).Encode(batchUpdateSpreadsheetResponse)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func mockSpreadsheetsHandlerFixture() *mockSpreadsheetsHandler {
	return &mockSpreadsheetsHandler{}
}

func mockBatchUpdateHandlerFixture(err error) *mockBatchUpdateHandler {
	return &mockBatchUpdateHandler{err: err}
}

func mockServerFixture(
	spreadsheetsHandler *mockSpreadsheetsHandler,
	batchUpdateHandler *mockBatchUpdateHandler,
) *httptest.Server {
	router := mux.NewRouter()
	// https://sheets.googleapis.com/v4/spreadsheets/{spreadsheetId}
	router.Handle(
		"/v4/spreadsheets/{spreadsheetId}", spreadsheetsHandler,
	).Methods("GET")
	// https://sheets.googleapis.com/v4/spreadsheets/{spreadsheetId}:batchUpdate
	router.Handle(
		"/v4/spreadsheets/{spreadsheetId}:batchUpdate", batchUpdateHandler,
	).Methods("POST")
	mockServer := httptest.NewServer(router)
	return mockServer
}

type mockSheetsConnection struct {
	srv *sheets.Service
}

func (c *mockSheetsConnection) Get(impl interface{}) {
	*(impl.(**sheets.Service)) = c.srv
}

func TestGoogleSheetOutput(t *testing.T) {
	// GIVEN
	spreadsheetsHandler := mockSpreadsheetsHandlerFixture()
	batchUpdateHandler := mockBatchUpdateHandlerFixture(nil)
	mockServer := mockServerFixture(spreadsheetsHandler, batchUpdateHandler)
	httpClient := mockServer.Client()
	srv, err := sheets.New(httpClient)
	srv.BasePath = mockServer.URL
	connection := &mockSheetsConnection{srv}

	sheetNumber := int64(1)
	spreadsheetId := spreadsheetIdFixture()
	sheetTitle := sheetTitleFixture(sheetNumber)
	gridCoordinate := gridCoordinateFixture(sheetNumber)
	startCellRef := fmt.Sprintf(
		"%s!%v%v",
		sheetTitle,
		string(rune(int64('A')+gridCoordinate.ColumnIndex)),
		gridCoordinate.RowIndex+1,
	)
	it := iteratorFixture()
	output := NewGoogleSheetOutput(spreadsheetId, startCellRef)

	expected := map[string]interface{}{
		"gridCoordinate": *gridCoordinate,
		"data":           dataFixture(),
		"delimiter":      "|",
		"type":           "PASTE_NORMAL",
	}

	// WHEN
	err = output.Write(connection, it)

	if err != nil {
		t.Fatal(err)
	}

	// THEN
	pasteDataRequest := batchUpdateHandler.request.Requests[0].PasteData
	actual := map[string]interface{}{
		"gridCoordinate": *pasteDataRequest.Coordinate,
		"data":           pasteDataRequest.Data,
		"delimiter":      pasteDataRequest.Delimiter,
		"type":           pasteDataRequest.Type,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"TestGoogleSheetOutput: expected %v, actual %v",
			expected,
			actual,
		)
	}
}

type mockIterator struct {
	data string
}

func (it *mockIterator) Next(data interface{}) bool {
	*(data.(*string)) = it.data
	return false
}

func TestGoogleSheetOutputFailures(t *testing.T) {
	// GIVEN
	sheetNumber := int64(1)
	testcases := []struct {
		testcase      string
		spreadsheetId string
		sheetTitle    string
		err           error
	}{
		{
			testcase:      "badSpreadsheetId",
			spreadsheetId: badSpreadsheetIdFixture(),
			sheetTitle:    sheetTitleFixture(sheetNumber),
			err:           nil,
		},
		{
			testcase:      "badSheetTitle",
			spreadsheetId: spreadsheetIdFixture(),
			sheetTitle:    badSheetTitleFixture(),
			err:           nil,
		},
		{
			testcase:      "forcedError",
			spreadsheetId: spreadsheetIdFixture(),
			sheetTitle:    sheetTitleFixture(sheetNumber),
			err:           errors.New("forced error"),
		},
	}

	for _, tt := range testcases {
		t.Run(
			tt.testcase,
			func(t *testing.T) {
				// GIVEN
				spreadsheetId := tt.spreadsheetId
				sheetTitle := tt.sheetTitle
				gridCoordinate := gridCoordinateFixture(sheetNumber)
				spreadsheetsHandler := mockSpreadsheetsHandlerFixture()
				batchUpdateHandler := mockBatchUpdateHandlerFixture(tt.err)
				mockServer := mockServerFixture(spreadsheetsHandler, batchUpdateHandler)
				httpClient := mockServer.Client()
				srv, err := sheets.New(httpClient)
				srv.BasePath = mockServer.URL
				connection := &mockSheetsConnection{srv}
				startCellRef := fmt.Sprintf(
					"%s!%v%v",
					sheetTitle,
					string(rune(int64('A')+gridCoordinate.ColumnIndex)),
					gridCoordinate.RowIndex+1,
				)
				it := iteratorFixture()
				output := NewGoogleSheetOutput(spreadsheetId, startCellRef)

				// WHEN
				err = output.Write(connection, it)

				// THEN
				if err == nil {
					t.Errorf("%v: Expected error. None returned.", tt.testcase)
				}
			},
		)
	}
}

func TestGoogleSheetsAPIMockServer(t *testing.T) {
	// GIVEN
	spreadsheetsHandler := mockSpreadsheetsHandlerFixture()
	batchUpdateHandler := mockBatchUpdateHandlerFixture(nil)
	mockServer := mockServerFixture(spreadsheetsHandler, batchUpdateHandler)
	client := mockServer.Client()
	sheetNumber := int64(1)
	spreadsheetId := spreadsheetIdFixture()

	// WHEN
	srv, err := sheets.New(client)
	// Update the basepath to use the mock url.
	srv.BasePath = mockServer.URL
	if err != nil {
		t.Fatal(err)
	}
	spreadsheetsGetCall := srv.Spreadsheets.Get(spreadsheetId)
	spreadsheet, err := spreadsheetsGetCall.Context(context.Background()).Do()
	if err != nil {
		t.Fatalf("Unable to get spreadsheets: %v", err)
	}
	var sheetId int64
	sheetId = -1
	for _, s := range spreadsheet.Sheets {
		if s.Properties.Title == sheetTitleFixture(sheetNumber) {
			sheetId = s.Properties.SheetId
			break
		}
	}
	if sheetId == -1 {
		t.Fatalf(`Sheet ${sheetName} not found`)
	}
	data := dataFixture()
	gridCoordinate := gridCoordinateFixture(sheetNumber)
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
		spreadsheetId,
		&batchUpdateSpreadsheetRequest,
	)
	_, err = call.Context(context.Background()).Do()
	if err != nil {
		log.Fatalf("Unable to store data in sheet: %v", err)
	}
}
