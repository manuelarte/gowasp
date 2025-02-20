{{ define "users/welcome.tpl"}}
    {{ template "layouts/header.tpl" .}}

        <h1>Welcome {{ .user.Username }} to Gowasp website</h1>
        <p>Warning: This is just for information purposes: Password hash: {{ .user.Password }} </p>

        <br>
        <h2>Posts</h2>
        <div class="intro-post">
            <h3>Intro Post</h1>
            <p id="first-post"></p>
        </div>

        <div id="latest-posts">
            <h2>Latest Posts</h2>
            <ul>
                {{range $val := .latestPosts}}
                    <li><a href="/posts/{{ $val.ID }}/view">{{ $val.Title }}</a></li>
                {{end}}
            </ul>
        <div>


        <script type="text/javascript">

            const firstPost = document.getElementById("first-post")
            fetch('/static/posts?name=intro.txt', {
                method: 'GET',
                headers: {
                    "Content-Type": "text/plain"
                }
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                response.text().then(data => firstPost.textContent = data)
            })
            .catch(error => {
                console.error('There has been a problem with your fetch operation:', error);
            });
        </script>
{{end}}