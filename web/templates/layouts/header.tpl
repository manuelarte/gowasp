{{ define "layouts/header.tpl"}}
    <!DOCTYPE html>
    <html lang="en">
        <head>
            <meta charset="utf-8">
            <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
            <meta name="description" content="">
            <meta name="author" content="">

            <title>Gowasp App</title>

            <style type="text/css" media="screen">
                    @import url("/css/index.css");
              </style>

            <!-- Custom fonts for this template -->
            <link href="https://fonts.googleapis.com/css?family=Lato:300,400,700,300italic,400italic,700italic" rel="stylesheet" type="text/css">

            <script>
                function logout() {
                    fetch('/users/logout', {
                        method: 'DELETE',
                    })
                    .then(response => {
                        if (!response.ok) {
                            throw new Error('Network response was not ok');
                        }
                        window.location.href = "/users/login"
                    })
                    .then(data => {
                        console.log(data); // Process your data here
                    })
                    .catch(error => {
                        console.error('There has been a problem with your fetch operation:', error);
                    });
                }
            </script>
        </head>
    <body>
        <div>
            {{ if .user }}
                <button onclick="logout()">Logout</button>
            {{ else }}
                <a href="/users/login" role="button">Login</a>
                <a href="/users/signup" role="button">Signup</a>
            {{ end }}
        </div>
{{end}}