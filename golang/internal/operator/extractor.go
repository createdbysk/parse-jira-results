package operator

type Extractor interface {
	Extract(q Query) Iterator
}

type extractor struct {
	c Connection
}

func NewExtractor(c Connection) Extractor {
	return &extractor{c}
}

func (e *extractor) Extract(q Query) Iterator {
	it := e.c.Execute(q)
	return it
}
