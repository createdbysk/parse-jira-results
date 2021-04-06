package operator

type Transformer interface {
	Transform(from interface{}, to interface{})
}

type FieldExtractor interface {
	Extract(from interface{}, to map[string]interface{})
}

type extractFieldsTransformer struct {
	fieldExtractors []FieldExtractor
}

func (t *extractFieldsTransformer) Transform(from interface{}, to interface{}) {
	result := to.(map[string]interface{})
	for _, fieldExtractor := range t.fieldExtractors {
		fieldExtractor.Extract(from, result)
	}
}

func NewExtractFieldsTransformer(fieldExtractors ...FieldExtractor) Transformer {
	return &extractFieldsTransformer{fieldExtractors}
}
