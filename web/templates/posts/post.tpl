{{ define "posts/post.tpl"}}
    {{ template "layouts/header.tpl" .}}

        <h2>{{ .post.Title }}</h2>
        <p>{{ .post.Content }}</p>

        <div>
            {{ template "posts/add_edit_comment.tpl" .}}
        </div>

        <div id="comments">
        This post has {{ len .comments.Data }} comment(s)
        <br>
        {{range $val := .comments.Data}}
            {{ template "posts/comment.tpl" $val }}
        {{end}}
        </div>
{{end}}
