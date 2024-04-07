{{ define "telegram.alert" -}}
{{ range .Alerts }}
Alert: {{ .Labels.alertname }}
Severity: {{ .Labels.severity }}
Start at:{{ .StartsAt}}
Silence: <a href="{{ .SilenceURL }}">{{ .SilenceURL }}</a>
{{ if gt (len .DashboardURL) 0 }}Dashboard: {{ .DashboardURL }}
{{ end }}
{{ if len .Annotations }}
Annotations:
{{ range .Annotations.SortedPairs -}}
- {{ .Name }}: {{ .Value }}
{{ end -}}
{{ end }}
{{ if len .Labels -}}
Labels:
{{ range .Labels.SortedPairs -}}
- {{ .Name }}: {{ .Value }}
{{ end -}}
{{ end }}
{{ end }}
{{- end }}

{{ define "telegram.message" }}
{{ if gt (len .Alerts.Firing) 0 }}
🔥Firing
{{ template "telegram.alert" . }}
{{end}}
{{ if gt (len .Alerts.Resolved) 0 }}
🟢Resolved
{{ template "telegram.alert" . }}
{{end}}
{{- end}}