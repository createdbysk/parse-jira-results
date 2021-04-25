package main

import (
	"os"

	"local.dev/sheetsLoader/internal/config"
	"local.dev/sheetsLoader/internal/operator"
)

type appContext struct {
	CredentialsFilePath           string
	SpreadsheetId                 string
	CellRef                       string
	Delimiter                     string
	GoogleSheetsConnectionFactory func(googleCtx *config.GoogleContext, credentialsFilePath string, scope ...string) (operator.Connection, error)
	GoogleSheetsOutputFactory     func(spreadsheetId string, cellRef string, delimiter string) operator.Output
	ReaderConnectionFactory       func(inputContext *config.InputContext) operator.Connection
	DelimitedTextInputFactory     func() operator.Input
}

func newAppContext() *appContext {
	credentialsFilePath := os.Getenv("CREDENTIALS_FILEPATH")
	spreadsheetId := os.Getenv("SPREADSHEET_ID")
	cellRef := os.Getenv("CELL_REF")
	delimiter := os.Getenv("DELIMITER")

	return &appContext{
		CredentialsFilePath:           credentialsFilePath,
		SpreadsheetId:                 spreadsheetId,
		CellRef:                       cellRef,
		Delimiter:                     delimiter,
		GoogleSheetsConnectionFactory: operator.NewGoogleSheetsConnection,
		GoogleSheetsOutputFactory:     operator.NewGoogleSheetsOutput,
		ReaderConnectionFactory:       operator.NewReaderConnection,
		DelimitedTextInputFactory:     operator.NewDelimitedTextInput,
	}
}
