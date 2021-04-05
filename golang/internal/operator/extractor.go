package operator

import "errors"

type Extractor interface {
	Extract(q Query) Iterator
}

func NewExtractor(c Connection) (Extractor, error) {
	return nil, errors.New("Not Implemented")
}
