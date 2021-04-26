package main

import "log"

// In order to test the main function,
// make a run() function that main invokes a variable
// that the test can update.
type appContextFactory func() *appContext
type errorReporter func(v ...interface{})
type runFn func(factory appContextFactory, reportError errorReporter)

var run runFn = func(factory appContextFactory, reportError errorReporter) {
	appContext := factory()

	googleContext := appContext.GoogleContextFactory()
	inputContext := appContext.InputContextFactory()

	readerConnection := appContext.ReaderConnectionFactory(inputContext)
	input := appContext.DelimitedTextInputFactory()

	googleSheetsConnection, err := appContext.GoogleSheetsConnectionFactory(
		googleContext,
		appContext.CredentialsFilePath,
		appContext.Scopes...,
	)
	if err != nil {
		reportError(err)
	}
	output := appContext.GoogleSheetsOutputFactory(
		appContext.SpreadsheetId,
		appContext.CellRef,
		appContext.Delimiter,
	)

	it, err := input.Read(readerConnection)
	if err != nil {
		reportError(err)
	}
	err = output.Write(googleSheetsConnection, it)
	if err != nil {
		reportError(err)
	}
}

func main() {
	run(newAppContext, log.Fatal)
}
