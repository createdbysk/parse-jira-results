package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *jwt.Config) *http.Client {
	return config.Client(oauth2.NoContext)
}

func main() {
	credsFilePath := os.Getenv("CREDENTIALS_FILEPATH")
	if credsFilePath == "" {
		log.Fatalf("CREDENTIALS_FILEPATH environment variable not defined")
	}
	// b, err := ioutil.ReadFile(credsFilePath)
	// if err != nil {
	// 	log.Fatalf("Unable to read client secret file: %v", err)
	// }

	// config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsScope)
	// if err != nil {
	// 	log.Fatalf("Unable to parse client secret file to config: %v", err)
	// }
	// client := getClient(config)

	srv, err := sheets.NewService(
		context.Background(),
		option.WithCredentialsFile(credsFilePath),
		option.WithScopes(sheets.SpreadsheetsScope),
	)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var data string
	for scanner.Scan() {
		data += fmt.Sprintf("%s\n", scanner.Text())
	}
	if err != nil {
		log.Fatalf("Unable to read data from stdin")
	}
	spreadsheetId := "1Wym64hatRsBYFCxktehIBycocCz2ZQZJBfrUCa6DlAQ"
	sheetName := "TestSheet"
	spreadsheetsGetCall := srv.Spreadsheets.Get(spreadsheetId)
	spreadsheet, err := spreadsheetsGetCall.Context(context.Background()).Do()
	if err != nil {
		log.Fatalf("Unable to get spreadsheets: %v", err)
	}
	var sheetId int64
	sheetId = -1
	for _, s := range spreadsheet.Sheets {
		if s.Properties.Title == sheetName {
			sheetId = s.Properties.SheetId
		}
	}
	if sheetId == -1 {
		log.Fatalf(`Sheet ${sheetName} not found`)
	}
	gridCoordinate := sheets.GridCoordinate{
		SheetId:     sheetId,
		ColumnIndex: 0,
		RowIndex:    0,
	}
	pasteDataRequest := sheets.PasteDataRequest{
		Coordinate: &gridCoordinate,
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
