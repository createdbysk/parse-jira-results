package main

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"local.dev/sheetsLoader/internal/config"
	"local.dev/sheetsLoader/internal/operator"
	"local.dev/sheetsLoader/internal/testutils"
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
	spreadsheetID          string
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
func (f *mockFactory) newGoogleSheetsOutput(spreadsheetID string, cellRef string, delimiter string) operator.Output {
	f.spreadsheetID = spreadsheetID
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
	spreadsheetID := fixture.SpreadsheetId()
	sheetTitle := fixture.SheetTitle()
	rowIndex := fixture.RowIndex()
	colIndex := fixture.ColIndex()
	cellRef := fixture.CellRef(rowIndex, colIndex)
	delimiter := fixture.Delimiter()
	scopes := fixture.Scopes()

	it := &mockIterator{}
	googleSheetsConnection := &mockGoogleSheetsConnection{}
	output := &mockOutput{err: nil}
	readerConnection := &mockReaderConnection{}
	input := &mockInput{err: nil, it: it}
	googleContext := &config.GoogleContext{}
	inputContext := &config.InputContext{}
	scopedCellRef := fixture.ScopedCellRef(sheetTitle, cellRef)
	commandlineArgs := []string{sheetTitle, cellRef}

	factory := &mockFactory{
		googleSheetsConnection: googleSheetsConnection,
		googleSheetsOutput:     output,
		readerConnection:       readerConnection,
		delimitedTextInput:     input,
	}

	appContextFactory := func(args []string) (*appContext, error) {
		return &appContext{
			CredentialsFilePath:           credentialsFilePath,
			SpreadsheetId:                 spreadsheetID,
			CellRef:                       scopedCellRef,
			Delimiter:                     delimiter,
			Scopes:                        scopes,
			GoogleContextFactory:          func() *config.GoogleContext { return googleContext },
			InputContextFactory:           func() *config.InputContext { return inputContext },
			GoogleSheetsConnectionFactory: factory.newGoogleSheetsConnection,
			GoogleSheetsOutputFactory:     factory.newGoogleSheetsOutput,
			ReaderConnectionFactory:       factory.newReaderConnection,
			DelimitedTextInputFactory:     factory.newDelimitedTextInput,
		}, nil
	}

	expected := map[string]interface{}{
		"credentialsFilePath":    credentialsFilePath,
		"spreadsheetId":          spreadsheetID,
		"cellRef":                scopedCellRef,
		"delimiter":              delimiter,
		"scopes":                 scopes,
		"googleContext":          googleContext,
		"inputContext":           inputContext,
		"googleSheetsConnection": googleSheetsConnection,
		"readerConnection":       readerConnection,
		"iterator":               it,
		"error":                  nil,
	}

	// WHEN
	err := run(appContextFactory, commandlineArgs)
	actual := map[string]interface{}{
		"credentialsFilePath":    factory.credentialsFilePath,
		"spreadsheetId":          factory.spreadsheetID,
		"cellRef":                factory.cellRef,
		"delimiter":              factory.delimiter,
		"scopes":                 factory.scopes,
		"googleContext":          factory.googleContext,
		"inputContext":           factory.inputContext,
		"googleSheetsConnection": output.c.(*mockGoogleSheetsConnection),
		"readerConnection":       input.c.(*mockReaderConnection),
		"iterator":               output.it.(*mockIterator),
		"error":                  err,
	}

	// THEN
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"TestRun: expected: %v, actual %v",
			expected,
			actual,
		)
	}
}

func TestRunFailures(t *testing.T) {
	// GIVEN
	expected := errors.New("An error")
	testcases := []struct {
		testcase        string
		appContextError error
		connectionError error
		inputError      error
		outputError     error
		expectedError   error
	}{
		{
			testcase:        "NewAppContextError",
			appContextError: expected,
		},
		{
			testcase:        "ConnectionError",
			connectionError: expected,
		},
		{
			testcase:   "InputError",
			inputError: expected,
		},
		{
			testcase:    "OutputError",
			outputError: expected,
		},
	}

	for _, tt := range testcases {
		t.Run(
			tt.testcase,
			func(t *testing.T) {
				credentialsFilePath := fixture.CredentialsFilePath()
				spreadsheetID := fixture.SpreadsheetId()
				sheetTitle := fixture.SheetTitle()
				rowIndex := fixture.RowIndex()
				colIndex := fixture.ColIndex()
				cellRef := fixture.CellRef(rowIndex, colIndex)
				delimiter := fixture.Delimiter()
				scopes := fixture.Scopes()

				it := &mockIterator{}
				googleSheetsConnection := &mockGoogleSheetsConnection{}
				output := &mockOutput{err: tt.outputError}
				readerConnection := &mockReaderConnection{}
				input := &mockInput{err: tt.inputError, it: it}
				googleContext := &config.GoogleContext{}
				inputContext := &config.InputContext{}
				scopedCellRef := fixture.ScopedCellRef(sheetTitle, cellRef)
				commandlineArgs := []string{sheetTitle, cellRef}

				factory := &mockFactory{
					googleSheetsConnection: googleSheetsConnection,
					googleSheetsOutput:     output,
					readerConnection:       readerConnection,
					delimitedTextInput:     input,
					err:                    tt.connectionError,
				}

				appContextFactory := func(args []string) (*appContext, error) {
					return &appContext{
						CredentialsFilePath:           credentialsFilePath,
						SpreadsheetId:                 spreadsheetID,
						CellRef:                       scopedCellRef,
						Delimiter:                     delimiter,
						Scopes:                        scopes,
						GoogleContextFactory:          func() *config.GoogleContext { return googleContext },
						InputContextFactory:           func() *config.InputContext { return inputContext },
						GoogleSheetsConnectionFactory: factory.newGoogleSheetsConnection,
						GoogleSheetsOutputFactory:     factory.newGoogleSheetsOutput,
						ReaderConnectionFactory:       factory.newReaderConnection,
						DelimitedTextInputFactory:     factory.newDelimitedTextInput,
					}, tt.appContextError
				}

				// WHEN
				actual := run(appContextFactory, commandlineArgs)

				// THEN
				if actual != expected {
					t.Errorf("TestRun: expected run() to report an error and it didn't.")
				}
			},
		)
	}
}

func TestMain(t *testing.T) {
	// GIVEN
	oldRun := run
	defer func() { run = oldRun }()

	os.Args = []string{"fakeProgramName", "fakeCellRef"}

	var storedAppContextFactory appContextFactory
	var storedCommandLineArgs []string
	run = func(factory appContextFactory, args []string) error {
		storedAppContextFactory = factory
		storedCommandLineArgs = args
		return nil
	}

	expected := map[string]interface{}{
		"appContextFactory": testutils.GetFnPtr(newAppContext),
		"commandLineArgs":   os.Args[1:],
	}

	// WHEN
	main()
	actual := map[string]interface{}{
		"appContextFactory": testutils.GetFnPtr(storedAppContextFactory),
		"commandLineArgs":   storedCommandLineArgs,
	}

	// THEN
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("TestMain: expected %v actual %v", expected, actual)
	}
}
