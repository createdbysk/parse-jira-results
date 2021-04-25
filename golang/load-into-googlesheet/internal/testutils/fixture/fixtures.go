package fixture

import (
	"fmt"
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

func CellRef(sheetTitle string, rowIndex int64, colIndex int64) string {
	cellRef := fmt.Sprintf(
		"%s!%v%v",
		sheetTitle,
		string(rune(int64('A')+colIndex)),
		rowIndex+1,
	)
	return cellRef
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
