{{ define "users/login.tpl"}}
    {{ template "layouts/header.tpl" .}}

    <script type="text/javascript">
      function login(event) {
        event.preventDefault(); // Prevent form from refreshing the page
        document.getElementById("loginError").style.visibility="hidden";
        username = document.getElementById("username").value
        password = document.getElementById("password").value
        const userLogin = {
          username: username,
          password: password
        };
        console.log("Sending data:", JSON.stringify(userLogin));
        fetch('/api/users/login', {
          method: 'POST',
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify(userLogin),
        })
        .then(response => {
            if (!response.ok) {
                document.getElementById("loginError").style.visibility="visible";
                throw new Error('Network response was not ok');
            }
            document.getElementById("loginError").style.visibility="hidden";
            window.location.href = "/users/welcome"
        })
        .catch(error => {
            document.getElementById("loginError").style.visibility="visible";
            console.error('There has been a problem with your fetch operation:', error);
        });
      }
    </script>

    <h2>Login</h2>
    <form onsubmit="login(event)">
        <label for="username">Username:</label><br>
        <input type="text" id="username" name="username" required><br>
        <label for="password">Password:</label><br>
        <input type="password" id="password" name="password" required><br>
        <input type="submit" value="Login">
    </form>
    <div id="loginError" style="color:red; visibility: hidden;">
      An error occurred
    </div>
{{end}}
