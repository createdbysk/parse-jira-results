package operator

import "io"

type MockRenderer struct{}

func (r *MockRenderer) Render(w io.Writer, data interface{}) {}
