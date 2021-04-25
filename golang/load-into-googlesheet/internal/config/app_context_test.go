package config

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	"local.dev/sheetsLoader/internal/testutils/fixture"
)

func inputReaderFixture() io.Reader {
	return strings.NewReader("test")
}

func errorFixture() error {
	return errors.New("test error")
}

type mockConnection struct{}

func newMockConnection() Connection {
	return &mockConnection{}
}

func (c *mockConnection) Get(impl interface{}) {}

type mockIterator struct{}

func newIterator() Iterator {
	return &mockIterator{}
}

func (it *mockIterator) Next(data interface{}) bool {
	return false
}

type mockInput struct {
	it  Iterator
	err error
}

func newMockInput(it Iterator, err error) Input {
	return &mockInput{it, err}
}

func (i *mockInput) Read(c Connection) (Iterator, error) {
	return i.it, i.err
}

type mockOutput struct {
	err error
}

func newMockOutput(err error) Output {
	return &mockOutput{err}
}

func (o *mockOutput) Write(c Connection, it Iterator) error {
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

	mockNewReaderConnection := NewMockNewReaderConnection(inputConnection)
	defer mockNewReaderConnection.Unpatch()

	expected := &AppContext{
		InputConnection:  inputConnection,
		OutputConnection: outputConnection,
		Input:            input,
		Output:           output,
	}
}
