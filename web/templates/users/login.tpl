{{ define "users/login.tpl"}}
    {{ template "layouts/header.tpl" .}}
    <h2>Login</h2>
    <form action="/users/login" method="POST">
        <label for="username">Username:</label><br>
        <input type="text" id="username" name="username" required><br>
        <label for="password">Password:</label><br>
        <input type="password" id="password" name="password" required><br>
        <input type="submit" value="Login">
    </form>
{{end}}
