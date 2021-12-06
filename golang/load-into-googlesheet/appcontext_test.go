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

	spreadsheetID := fixture.SpreadsheetId()
	os.Setenv("SPREADSHEET_ID", spreadsheetID)

	delimiter := fixture.Delimiter()
	os.Setenv("DELIMITER", delimiter)

	sheetTitle := fixture.SheetTitle()
	rowIndex := fixture.RowIndex()
	colIndex := fixture.ColIndex()
	cellRef := fixture.CellRef(rowIndex, colIndex)
	scopedCellRef := fixture.ScopedCellRef(sheetTitle, cellRef)
	args := []string{sheetTitle, cellRef}

	scopes := fixture.Scopes()

	expected := map[string]interface{}{
		"CredentialsFilePath":           credentialsFilePath,
		"SpreadsheetId":                 spreadsheetID,
		"CellRef":                       scopedCellRef,
		"Delimiter":                     delimiter,
		"Scopes":                        scopes,
		"GoogleContextFactory":          testutils.GetFnPtr(config.NewGoogleContext),
		"InputContextFactory":           testutils.GetFnPtr(config.NewInputContext),
		"GoogleSheetsConnectionFactory": testutils.GetFnPtr(operator.NewGoogleSheetsConnection),
		"GoogleSheetsOutputFactory":     testutils.GetFnPtr(operator.NewGoogleSheetsOutput),
		"ReaderConnectionFactory":       testutils.GetFnPtr(operator.NewReaderConnection),
		"DelimitedTextInputFactory":     testutils.GetFnPtr(operator.NewDelimitedTextInput),
		"error":                         nil,
	}

	// WHEN
	appContext, err := newAppContext(args)
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
		"error":                         err,
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

func TestNewAppContextFailures(t *testing.T) {
	// GIVEN
	allEnvVars := []string{
		"CREDENTIALS_FILEPATH",
		"SPREADSHEET_ID",
		"DELIMITER",
	}
	// Clear the environment.
	os.Clearenv()
	// The program expects to caller to set all environment variables.
	// For each test case, set the environment variable to skip.
	// If there is no error, then the code did not check for that env variable.
	mustSetEnvironmentVariables := "must set environment variables CREDENTIALS_FILEPATH, SPREADSHEET_ID, and DELIMITER"
	testcases := []struct {
		testcase            string
		skip                string
		expectedErrorString string
	}{
		{
			testcase:            "CREDENTIALS_FILEPATH",
			skip:                "CREDENTIALS_FILEPATH",
			expectedErrorString: mustSetEnvironmentVariables,
		},
		{
			testcase:            "SPREADSHEET_ID",
			skip:                "SPREADSHEET_ID",
			expectedErrorString: mustSetEnvironmentVariables,
		},
		{
			testcase:            "SHEET_TITLE",
			skip:                "SHEET_TITLE",
			expectedErrorString: "must provide the sheet title and destination cell reference as command-line parameters",
		},
		{
			testcase:            "CELL_REF",
			skip:                "CELL_REF",
			expectedErrorString: "must provide the sheet title and destination cell reference as command-line parameters",
		},
		{
			testcase:            "DELIMITER",
			skip:                "DELIMITER",
			expectedErrorString: mustSetEnvironmentVariables,
		},
	}

	for _, tt := range testcases {
		t.Run(
			tt.testcase,
			func(t *testing.T) {
				// GIVEN
				for _, envVar := range allEnvVars {
					if envVar != tt.skip {
						os.Setenv(envVar, envVar)
						defer os.Clearenv()
					}
				}
				var args []string
				if tt.skip != "SHEET_TITLE" {
					args = append(args, "SHEET_TITLE")
				}
				if tt.skip != "CELL_REF" {
					args = append(args, "CELL_REF")
				}

				// WHEN
				_, err := newAppContext(args)

				// THEN
				if err == nil {
					t.Errorf("expected an error and got none")
				} else if err.Error() != tt.expectedErrorString {
					t.Errorf(
						"expected an error with text %v, got text %v",
						tt.expectedErrorString,
						err.Error(),
					)
				}
			},
		)
	}
}
