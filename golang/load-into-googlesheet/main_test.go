package main

import (
	"reflect"
	"testing"

	"local.dev/sheetsLoader/internal/config"
	"local.dev/sheetsLoader/internal/operator"
	"local.dev/sheetsLoader/internal/testutils/fixture"
)

type mockIterator struct{}

func (it *mockIterator) Next(data interface{}) bool {
	return false
}

type mockGoogleSheetsConnection struct{}

func (c *mockGoogleSheetsConnection) Get(impl interface{}) {
	// Just a stub. Will do nothing.
}

type mockOutput struct {
	err error
	c   operator.Connection
	it  operator.Iterator
}

func (o *mockOutput) Write(c operator.Connection, it operator.Iterator) error {
	o.c = c
	o.it = it
	return o.err
}

type mockReaderConnection struct{}

func (c *mockReaderConnection) Get(impl interface{}) {
	// Just a stub.
}

type mockInput struct {
	err error
	it  operator.Iterator
	c   operator.Connection
}

func (i *mockInput) Read(c operator.Connection) (operator.Iterator, error) {
	i.c = c
	return i.it, i.err
}

type mockFactory struct {
	credentialsFilePath    string
	scopes                 []string
	spreadsheetId          string
	cellRef                string
	delimiter              string
	err                    error
	googleContext          *config.GoogleContext
	inputContext           *config.InputContext
	googleSheetsConnection *mockGoogleSheetsConnection
	googleSheetsOutput     *mockOutput
	readerConnection       *mockReaderConnection
	delimitedTextInput     *mockInput
}

func (f *mockFactory) newGoogleSheetsConnection(googleCtx *config.GoogleContext, credentialsFilePath string, scope ...string) (operator.Connection, error) {
	f.googleContext = googleCtx
	f.credentialsFilePath = credentialsFilePath
	f.scopes = scope
	return f.googleSheetsConnection, f.err
}
func (f *mockFactory) newGoogleSheetsOutput(spreadsheetId string, cellRef string, delimiter string) operator.Output {
	f.spreadsheetId = spreadsheetId
	f.cellRef = cellRef
	f.delimiter = delimiter
	return f.googleSheetsOutput
}

func (f *mockFactory) newReaderConnection(inputContext *config.InputContext) operator.Connection {
	f.inputContext = inputContext
	return f.readerConnection
}

func (f *mockFactory) newDelimitedTextInput() operator.Input {
	return f.delimitedTextInput
}

func TestRun(t *testing.T) {
	// GIVEN
	credentialsFilePath := fixture.CredentialsFilePath()
	spreadsheetId := fixture.SpreadsheetId()
	sheetTitle := fixture.SheetTitle()
	rowIndex := fixture.RowIndex()
	colIndex := fixture.ColIndex()
	cellRef := fixture.CellRef(sheetTitle, rowIndex, colIndex)
	delimiter := fixture.Delimiter()
	scopes := fixture.Scopes()

	it := &mockIterator{}
	googleSheetsConnection := &mockGoogleSheetsConnection{}
	output := &mockOutput{err: nil}
	readerConnection := &mockReaderConnection{}
	input := &mockInput{err: nil, it: it}
	googleContext := &config.GoogleContext{}
	inputContext := &config.InputContext{}

	factory := &mockFactory{
		googleSheetsConnection: googleSheetsConnection,
		googleSheetsOutput:     output,
		readerConnection:       readerConnection,
		delimitedTextInput:     input,
	}

	appContextFactory := func() *appContext {
		return &appContext{
			CredentialsFilePath:           credentialsFilePath,
			SpreadsheetId:                 spreadsheetId,
			CellRef:                       cellRef,
			Delimiter:                     delimiter,
			Scopes:                        scopes,
			GoogleContextFactory:          func() *config.GoogleContext { return googleContext },
			InputContextFactory:           func() *config.InputContext { return inputContext },
			GoogleSheetsConnectionFactory: factory.newGoogleSheetsConnection,
			GoogleSheetsOutputFactory:     factory.newGoogleSheetsOutput,
			ReaderConnectionFactory:       factory.newReaderConnection,
			DelimitedTextInputFactory:     factory.newDelimitedTextInput,
		}
	}

	expected := map[string]interface{}{
		"credentialsFilePath":    credentialsFilePath,
		"spreadsheetId":          spreadsheetId,
		"cellRef":                cellRef,
		"delimiter":              delimiter,
		"scopes":                 scopes,
		"googleContext":          googleContext,
		"inputContext":           inputContext,
		"googleSheetsConnection": googleSheetsConnection,
		"readerConnection":       readerConnection,
		"iterator":               it,
	}

	// WHEN
	run(appContextFactory)
	actual := map[string]interface{}{
		"credentialsFilePath":    factory.credentialsFilePath,
		"spreadsheetId":          factory.spreadsheetId,
		"cellRef":                factory.cellRef,
		"delimiter":              factory.delimiter,
		"scopes":                 factory.scopes,
		"googleContext":          factory.googleContext,
		"inputContext":           factory.inputContext,
		"googleSheetsConnection": output.c.(*mockGoogleSheetsConnection),
		"readerConnection":       input.c.(*mockReaderConnection),
		"iterator":               output.it.(*mockIterator),
	}
	// THEN

	// THEN
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"TestRun: expected: %v, actual %v",
			expected,
			actual,
		)
	}
}
