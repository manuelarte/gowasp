{{ define "errors/400.tpl"}}
    <h1>Error</h1>

    {{ if .message }}
        <p>An error occurred: {{ .message }}</p>
    {{else}}
        <p>An error occurred</p>
    {{end}}

{{end}}