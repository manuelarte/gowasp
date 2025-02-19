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

            <script type="text/javascript">
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
                function changeToolbarAClass(event) {
                    console.log(event)
                }
            </script>
        </head>
    <body>

        <div class="topnav">
            {{ if .user }}
                <a class="toolbar active" href="/users/welcome">Home</a>
                <button class="login-container button" onclick="logout()">Logout</button>
            {{ else }}
                <a class="toolbar" href="/users/login">Login</a>
                <a class="toolbar" href="/users/signup">Signup</a>
            {{ end }}
        </div>

        <script type="text/javascript">
            document.querySelectorAll('a.toolbar').forEach(function(ele, idx) {
                ele.classList.remove('active');
            });
            if (window.location.pathname === '/users/login' || window.location.pathname === '/users/welcome') {
                document.querySelectorAll('a.toolbar')[0].classList.add('active')
            }
            if (window.location.pathname === '/users/signup') {
                document.querySelectorAll('a.toolbar')[1].classList.add('active')
            }
        </script>
{{end}}