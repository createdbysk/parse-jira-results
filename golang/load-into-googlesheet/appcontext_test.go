package main

import (
	"os"
	"reflect"
	"testing"

	"local.dev/sheetsLoader/internal/config"
	"local.dev/sheetsLoader/internal/operator"
	"local.dev/sheetsLoader/internal/testutils"
	"local.dev/sheetsLoader/internal/testutils/fixture"
)

func TestNewAppContext(t *testing.T) {
	// GIVEN
	credentialsFilePath := fixture.CredentialsFilePath()
	os.Setenv("CREDENTIALS_FILEPATH", credentialsFilePath)

	spreadsheetId := fixture.SpreadsheetId()
	os.Setenv("SPREADSHEET_ID", spreadsheetId)

	sheetTitle := fixture.SheetTitle()
	rowIndex := fixture.RowIndex()
	colIndex := fixture.ColIndex()
	cellRef := fixture.CellRef(sheetTitle, rowIndex, colIndex)
	os.Setenv("CELL_REF", cellRef)

	delimiter := fixture.Delimiter()
	os.Setenv("DELIMITER", delimiter)

	scopes := fixture.Scopes()

	expected := map[string]interface{}{
		"CredentialsFilePath":           credentialsFilePath,
		"SpreadsheetId":                 spreadsheetId,
		"CellRef":                       cellRef,
		"Delimiter":                     delimiter,
		"Scopes":                        scopes,
		"GoogleContextFactory":          testutils.GetFnPtr(config.NewGoogleContext),
		"InputContextFactory":           testutils.GetFnPtr(config.NewInputContext),
		"GoogleSheetsConnectionFactory": testutils.GetFnPtr(operator.NewGoogleSheetsConnection),
		"GoogleSheetsOutputFactory":     testutils.GetFnPtr(operator.NewGoogleSheetsOutput),
		"ReaderConnectionFactory":       testutils.GetFnPtr(operator.NewReaderConnection),
		"DelimitedTextInputFactory":     testutils.GetFnPtr(operator.NewDelimitedTextInput),
	}

	// WHEN
	appContext := newAppContext()
	actual := map[string]interface{}{
		"CredentialsFilePath":           appContext.CredentialsFilePath,
		"SpreadsheetId":                 appContext.SpreadsheetId,
		"CellRef":                       appContext.CellRef,
		"Delimiter":                     appContext.Delimiter,
		"Scopes":                        appContext.Scopes,
		"GoogleContextFactory":          testutils.GetFnPtr(appContext.GoogleContextFactory),
		"InputContextFactory":           testutils.GetFnPtr(appContext.InputContextFactory),
		"GoogleSheetsConnectionFactory": testutils.GetFnPtr(appContext.GoogleSheetsConnectionFactory),
		"GoogleSheetsOutputFactory":     testutils.GetFnPtr(appContext.GoogleSheetsOutputFactory),
		"ReaderConnectionFactory":       testutils.GetFnPtr(appContext.ReaderConnectionFactory),
		"DelimitedTextInputFactory":     testutils.GetFnPtr(appContext.DelimitedTextInputFactory),
	}

	// THEN
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"TestAppContext: expected: %v, actual %v",
			expected,
			actual,
		)
	}
}
