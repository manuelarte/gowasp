{{ define "users/signup.tpl"}}
    {{ template "layouts/header.tpl" .}}

        <script type="text/javascript">
              function signup(event) {
                event.preventDefault(); // Prevent form from refreshing the page
                document.getElementById("signupError").style.visibility="hidden";
                username = document.getElementById("username").value
                password =document.getElementById("password").value
                const userSignup = {
                  username: username,
                  password: password
                };
                console.log("Sending data:", JSON.stringify(userSignup));
                fetch('/users/signup', {
                  method: 'POST',
                  headers: {
                    "Content-Type": "application/json"
                  },
                  body: JSON.stringify(userSignup),
                })
                .then(response => {
                    if (!response.ok) {
                        document.getElementById("signupError").style.visibility="visible";
                        throw new Error('Network response was not ok');
                    }
                    document.getElementById("signupError").style.visibility="hidden";
                    window.location.href = "/users/welcome"
                })
                .then(data => {
                    console.log(data); // Process your data here
                })
                .catch(error => {
                    document.getElementById("signupError").style.visibility="visible";
                    console.error('There has been a problem with your fetch operation:', error);
                });
              }
            </script>


        <h2>Sign up</h2>
        <form onsubmit="signup(event)">
            <label for="username">Username:</label><br>
            <input type="text" id="username" name="username" required><br>
            <label for="password">Password:</label><br>
            <input type="password" id="password" name="password" required><br>
            <input type="submit" value="Signup">
        </form>
        <div id="signupError" style="color:red; visibility: hidden;">
              An error occurred
        </div>
{{end}}


