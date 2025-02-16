{{ define "blogs/one.tpl"}}
    {{ template "layouts/header.tpl" .}}

        <h2>{{ .blog.Title }}</h2>
        <p>{{ .blog.Content }}</p>

        <div>
            <p>Add comment</p>
        </div>
{{end}}
