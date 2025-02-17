{{ define "blogs/one.tpl"}}
    {{ template "layouts/header.tpl" .}}

        <h2>{{ .blog.Title }}</h2>
        <p>{{ .blog.Content }}</p>

        <div>
            <p>Add comment</p>
        </div>

        <div id="comments">
        This post has {{ len .comments.Data }} comment(s)
        <br>
        {{range $val := .comments.Data}}
            {{ template "blogs/comment.tpl" $val }}
        {{end}}
        </div>
{{end}}
