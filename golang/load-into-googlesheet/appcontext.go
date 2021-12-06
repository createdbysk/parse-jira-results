package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"google.golang.org/api/sheets/v4"
	"local.dev/sheetsLoader/internal/config"
	"local.dev/sheetsLoader/internal/operator"
)

type appContext struct {
	CredentialsFilePath           string
	SpreadsheetId                 string
	CellRef                       string
	Delimiter                     string
	Scopes                        []string
	GoogleContextFactory          func() *config.GoogleContext
	InputContextFactory           func() *config.InputContext
	GoogleSheetsConnectionFactory func(googleCtx *config.GoogleContext, credentialsFilePath string, scope ...string) (operator.Connection, error)
	GoogleSheetsOutputFactory     func(spreadsheetId string, cellRef string, delimiter string) operator.Output
	ReaderConnectionFactory       func(inputContext *config.InputContext) operator.Connection
	DelimitedTextInputFactory     func() operator.Input
}

func parseCommandLineArgs(args []string) (string, string, error) {
	flagSet := flag.NewFlagSet("", flag.ContinueOnError)

	flagSet.Parse(args)
	if flagSet.NArg() < 2 {
		return "", "", errors.New("must provide the sheet title and destination cell reference as command-line parameters")
	}

	sheetTitle := flagSet.Arg(0)
	cellRef := flagSet.Arg(1)

	return sheetTitle, cellRef, nil
}

func newAppContext(args []string) (*appContext, error) {
	credentialsFilePath := os.Getenv("CREDENTIALS_FILEPATH")
	spreadsheetId := os.Getenv("SPREADSHEET_ID")
	delimiter := os.Getenv("DELIMITER")
	sheetTitle, cellRef, err := parseCommandLineArgs(args)
	if credentialsFilePath == "" || spreadsheetId == "" || delimiter == "" {
		return nil, errors.New("must set environment variables CREDENTIALS_FILEPATH, SPREADSHEET_ID, and DELIMITER")
	}
	scopedCellRef := fmt.Sprintf(
		"%s!%s",
		sheetTitle,
		cellRef,
	)
	scopes := []string{
		sheets.SpreadsheetsScope,
	}

	return &appContext{
		CredentialsFilePath:           credentialsFilePath,
		SpreadsheetId:                 spreadsheetId,
		CellRef:                       scopedCellRef,
		Delimiter:                     delimiter,
		Scopes:                        scopes,
		GoogleContextFactory:          config.NewGoogleContext,
		InputContextFactory:           config.NewInputContext,
		GoogleSheetsConnectionFactory: operator.NewGoogleSheetsConnection,
		GoogleSheetsOutputFactory:     operator.NewGoogleSheetsOutput,
		ReaderConnectionFactory:       operator.NewReaderConnection,
		DelimitedTextInputFactory:     operator.NewDelimitedTextInput,
	}, err
}
