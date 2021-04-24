package config

import (
	"io"
	"os"
)

// InputContext is the context data for the input.
type InputContext struct {
	Reader io.Reader
}

func NewInputContext() *InputContext {
	return &InputContext{Reader: os.Stdin}
}
