{{- define "IsStart" -}}
{{/* Pass a []ChangeLogHistory parameter - https://pkg.go.dev/github.com/andygrunwald/go-jira#ChangelogHistory */}}
{{- range . -}}
{{- $created := .Created -}}
{{- range .Items -}}
{{- if and (eq .Field "status") (eq .FromString "To Do") (eq .ToString "In Progress") -}}
{{ $created }}
{{- end -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "IsEnd" -}}
{{/* Pass a []ChangeLogHistory parameter - https://pkg.go.dev/github.com/andygrunwald/go-jira#ChangelogHistory */}}
{{- range . -}}
{{- $created := .Created -}}
{{- range .Items -}}
{{- if and (eq .Field "status") (eq .FromString "In Progress") (eq .ToString "Done") -}}
{{ $created }}
{{- end -}}
{{- end -}}
{{- end -}}
{{- end -}}

|Name|Type|
{{- range .}}
|{{.Key}}|{{.Fields.Type.Name}}|{{template "IsStart" .Changelog.Histories}}|{{template "IsEnd" .Changelog.Histories}}|
{{- end }}
