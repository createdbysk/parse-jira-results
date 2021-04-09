{{- define "Priority" -}}
{{.ID}} - {{.Name}}
{{- end -}}
{{- define "Date" -}}
{{ slice . 0 10}}
{{- end -}}

{{- define "IsStart" -}}
{{/* Pass a []ChangeLogHistory parameter - https://pkg.go.dev/github.com/andygrunwald/go-jira#ChangelogHistory */}}
{{- $memo := false -}}
{{- range . -}}
{{- $created := .Created -}}
{{- range .Items -}}
{{- if and (not $memo) (eq .Field "status") (eq .FromString "To Do") (eq .ToString "In Progress") -}}
{{ template "Date" $created }}
{{- $memo = true -}}
{{- end -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{- define "IsEnd" -}}
{{/* Pass a []ChangeLogHistory parameter - https://pkg.go.dev/github.com/andygrunwald/go-jira#ChangelogHistory */}}
{{- $memo := false -}}
{{- $created := false -}}
{{- range . -}}
{{- $created = .Created -}}
{{- range .Items -}}
{{- if and (eq .Field "status") (eq .ToString "Done") -}}
{{- $memo = true -}}
{{- end -}}
{{- end -}}
{{- end -}}
{{- if $memo -}}
{{ template "Date" $created }}
{{- end -}}
{{- end -}}

{{/*|Name|Type|StartDate|EndDate|*/}}
{{- range . -}}
{{.Key}}|{{.Fields.Type.Name}}|{{ template "Priority" .Fields.Priority }}|{{template "IsStart" .Changelog.Histories}}|{{template "IsEnd" .Changelog.Histories}}
{{ end -}}
