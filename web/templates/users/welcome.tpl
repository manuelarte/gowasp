{{ define "users/welcome.tpl"}}
    {{ template "layouts/header.tpl" .}}

        <script type="text/javascript">
        function createUl(blogs) {
            var list = document.createElement('ul');

            for (var i = 0; i < blogs.length; i++) {
                const blog = blogs[i]

                var a = document.createElement('a');
                var linkText = document.createTextNode(blog.title);
                a.appendChild(linkText);
                a.title = blog.title;
                a.href = "/blogs/" + blog.id + "/view";
                
                const item = document.createElement('li');
                item.appendChild(a);
                list.appendChild(item);
            }
            return list;
        }
        </script>

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
                console.log(firstBlog)
                response.text().then(data => firstBlog.textContent = data)
            })
            .catch(error => {
                console.error('There has been a problem with your fetch operation:', error);
            });

            fetch('/blogs?page=0&size=5', {
                method: 'GET',
                headers: {
                    "Content-Type": "application/json"
                }
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                response.json().then(data => {
                    const ul = createUl(data.data)
                    const latestBlogs = document.getElementById("latest-blogs")
                    latestBlogs.appendChild(ul)
                })
            })
            .catch(error => {
                console.error('There has been a problem with your fetch operation:', error);
            });
        </script>
{{end}}