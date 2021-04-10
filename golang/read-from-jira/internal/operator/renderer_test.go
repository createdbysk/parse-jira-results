package operator

import (
	"bytes"
	"testing"
)

func patternFixture() string {
	pattern := `{{- /* Skip this line*/ -}}
|Name|Type|
{{- range .}}
|{{.Key}}|{{.Fields.Type.Name}}|
{{- end}}
`
	return pattern
}

func TestRender(t *testing.T) {
	// GIVEN
	pattern := patternFixture()
	renderer := NewTemplateRenderer(pattern)
	issues := jiraSearchResultsIssuesFixture(3)

	expected := `|Name|Type|
|TEST-26|Task|
|TEST-25|Story|
|TEST-24|Task|
`

	// WHEN
	var s string
	w := bytes.NewBufferString(s)
	renderer.Render(w, issues)
	actual := w.String()

	// THEN
	if actual != expected {
		t.Errorf(
			"TestRender: expected\n%v\nactual\n%v",
			expected,
			actual,
		)
	}
}
