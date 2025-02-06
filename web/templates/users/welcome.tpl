{{ define "users/welcome.tpl"}}
    {{ template "layouts/header.tpl" .}}
        <h1>Welcome {{ .user.Email }} to Gowasp website</h1>
        <p>Warning: This is just for information purposes: Password hash: {{ .user.Password }} </p>

        <p>Todo links to blogs</p>
{{end}}