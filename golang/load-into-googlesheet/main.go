package main

import (
	"log"
	"os"
)

// In order to test the main function,
// make a run() function that main invokes a variable
// that the test can update.
type appContextFactory func(args []string) (*appContext, error)
type runFn func(factory appContextFactory, args []string) error

var run runFn = func(factory appContextFactory, args []string) error {
	appContext, err := factory(args)
	if err != nil {
		return err
	}

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
		return err
	}
	output := appContext.GoogleSheetsOutputFactory(
		appContext.SpreadsheetId,
		appContext.CellRef,
		appContext.Delimiter,
	)

	it, err := input.Read(readerConnection)
	if err != nil {
		return err
	}
	err = output.Write(googleSheetsConnection, it)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := run(newAppContext, os.Args[1:])
	if err != nil {
		// The test will not cover these two lines of code.
		// Change it your own risk.
		log.Fatal(err)
		os.Exit(1)
	}
}
