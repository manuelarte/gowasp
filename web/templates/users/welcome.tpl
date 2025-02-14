{{ define "users/welcome.tpl"}}
    {{ template "layouts/header.tpl" .}}

        <h1>Welcome {{ .user.Username }} to Gowasp website</h1>
        <p>Warning: This is just for information purposes: Password hash: {{ .user.Password }} </p>

        <br>
        <div class="blog-post">
            <h1>Intro Blog</h1>
            <p id="first-blog"></p>
        </div>


        <script type="text/javascript">

            const firstBlog = document.getElementById("first-blog")
            fetch('/static/blogs?name=intro.txt', {
                method: 'GET',
                headers: {
                    "Content-Type": "text/plain"
                }
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                console.log(firstBlog)
                response.text().then(data => firstBlog.textContent = data)
            })
            .catch(error => {
                console.error('There has been a problem with your fetch operation:', error);
            });
        </script>
{{end}}