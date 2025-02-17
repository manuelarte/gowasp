{{ define "blogs/comment.tpl"}}
<p>User: {{ .User.Username }}</p>
<p>{{ .Comment | unsafe }}</p>
{{ end }}