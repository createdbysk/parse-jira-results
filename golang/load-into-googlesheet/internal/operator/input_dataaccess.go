package operator

import (
	"bufio"
	"fmt"
	"io"

	"local.dev/sheetsLoader/internal/config"
)

type readerConnection struct {
	r io.Reader
}

func NewReaderConnection(inputContext *config.InputContext) Connection {
	return &readerConnection{r: inputContext.Reader}
}

func (c *readerConnection) Get(impl interface{}) {
	*(impl.(*io.Reader)) = c.r
}

type readerIterator struct {
	s *bufio.Scanner
}

func (it *readerIterator) Next(data interface{}) bool {
	var result string
	for it.s.Scan() {
		result += fmt.Sprintf("%s\n", it.s.Text())
	}
	*(data.(*string)) = result
	return false
}

type delimitedTextInput struct{}

func NewDelimitedTextInput() Input {
	return &delimitedTextInput{}
}

func (d *delimitedTextInput) Read(c Connection) (Iterator, error) {
	var r io.Reader
	c.Get(&r)

	s := bufio.NewScanner(r)
	it := readerIterator{s}
	return &it, nil
}
