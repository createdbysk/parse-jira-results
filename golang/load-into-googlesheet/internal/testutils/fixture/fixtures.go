package fixture

import "fmt"

func SpreadsheetId() string {
	return "spreadsheetid"
}

func SheetTitle() string {
	return "TAB"
}

func RowIndex() int {
	return 42
}

func ColIndex() int {
	return 24
}

func CellRef(sheetTitle string, rowIndex int, colIndex int) string {
	cellRef := fmt.Sprintf(
		"%s!%v%v",
		sheetTitle,
		string(rune(int64('A')+int64(colIndex))),
		rowIndex+1,
	)
	return cellRef
}

func CredentialsFilePath() string {
	return "/path/to/credentials.json"
}
