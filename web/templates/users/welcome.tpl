{{ define "users/welcome.tpl"}}
    {{ template "layouts/header.tpl" .}}

        <h1>Welcome {{ .user.Username }} to Gowasp website</h1>
        <p>Warning: This is just for information purposes: Password hash: {{ .user.Password }} </p>

        <br>
        <h2>Blogs</h2>
        <div class="blog-post">
            <h3>Intro Blog</h1>
            <p id="first-blog"></p>
        </div>

        <div id="latest-blogs">
            <h2>Latest Blogs</h2>
            <ul>
                {{range $val := .latestBlogs}}
                    <li><a href="/blogs/{{ $val.ID }}/view">{{ $val.Title }}</a></li>
                {{end}}
            </ul>
        <div>


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
                response.text().then(data => firstBlog.textContent = data)
            })
            .catch(error => {
                console.error('There has been a problem with your fetch operation:', error);
            });
        </script>
{{end}}