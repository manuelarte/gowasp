{{ define "blogs/comment.tpl"}}
<p>User: {{ .UserID }}</p>
<p>{{ .Comment }}</p>
{{ end }}