package operator

import (
	"encoding/csv"
	"io"
)

type csvFileLoader struct {
	writer *csv.Writer
}

func (loader *csvFileLoader) Load(record []string) error {
	err := loader.writer.Write(record)
	return err
}

func (loader *csvFileLoader) Flush() {
	loader.writer.Flush()
}

func NewCSVFileLoader(w io.Writer, delimiter rune) (Loader, error) {
	writer := csv.NewWriter(w)
	writer.Comma = delimiter
	loader := csvFileLoader{writer}
	return &loader, nil
}
