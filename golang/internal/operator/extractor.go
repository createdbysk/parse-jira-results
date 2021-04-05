package operator

import "errors"

type Extractor interface {
	Extract(q Query) Iterator
}

type extractor struct {
	c Connection
}

func NewExtractor(c Connection) (Extractor, error) {
	return nil, errors.New("Not Implemented")
}

func (e *extractor) Extract(q Query) Iterator {
	it := e.c.Execute(q)
	return it
}
