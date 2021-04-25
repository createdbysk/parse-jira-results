package main

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	"local.dev/sheetsLoader/internal/operator"
	"local.dev/sheetsLoader/internal/testutils/fixture"
)

func inputReaderFixture() io.Reader {
	return strings.NewReader("test")
}

func errorFixture() error {
	return errors.New("test error")
}

type mockConnection struct{}

func newMockConnection() operator.Connection {
	return &mockConnection{}
}

func (c *mockConnection) Get(impl interface{}) {}

type mockIterator struct{}

func newIterator() operator.Iterator {
	return &mockIterator{}
}

func (it *mockIterator) Next(data interface{}) bool {
	return false
}

type mockInput struct {
	it  operator.Iterator
	err error
}

func newMockInput(it operator.Iterator, err error) operator.Input {
	return &mockInput{it, err}
}

func (i *mockInput) Read(c operator.Connection) (operator.Iterator, error) {
	return i.it, i.err
}

type mockOutput struct {
	err error
}

func newMockOutput(err error) operator.Output {
	return &mockOutput{err}
}

func (o *mockOutput) Write(c operator.Connection, it operator.Iterator) error {
	return o.err
}

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
	os.Setenv("CELLREF", cellRef)

	inputError := error(nil)
	inputReader := inputReaderFixture()

	iterator := newIterator()
	inputConnection := newMockConnection()
	input := newMockInput(iterator, inputError)

	outputError := error(nil)
	outputConnection := newMockConnection()
	output := newMockOutput(outputError)

	mockNewReaderConnection := operator.MockNewReaderConnection(inputConnection)
	defer mockNewReaderConnection.Unpatch()

	mockNewReaderConnection := operator.MockNewReaderConnection(inputConnection)
	defer mockNewReaderConnection.Unpatch()
}
