package fixture

import (
	"fmt"

	"google.golang.org/api/sheets/v4"
)

func SpreadsheetId() string {
	return "spreadsheetid"
}

func SheetTitle() string {
	return "TAB"
}

func RowIndex() int64 {
	return 42
}

func ColIndex() int64 {
	return 24
}

func CellRef(rowIndex int64, colIndex int64) string {
	cellRef := fmt.Sprintf(
		"%v%v",
		string(rune(int64('A')+colIndex)),
		rowIndex+1,
	)
	return cellRef
}

func ScopedCellRef(sheetTitle string, cellRef string) string {
	scopedCellRef := fmt.Sprintf(
		"%s!%s",
		sheetTitle,
		cellRef,
	)
	return scopedCellRef
}

func CredentialsFilePath() string {
	return "/path/to/credentials.json"
}

func Delimiter() string {
	return "+"
}

func Data(delimiter string) string {
	return fmt.Sprintf("Hello%sWorld", delimiter)
}

func Scopes() []string {
	return []string{
		sheets.SpreadsheetsScope,
	}
}
