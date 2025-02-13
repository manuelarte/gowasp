{{ define "users/signup.tpl"}}
    {{ template "layouts/header.tpl" .}}

        <script type="text/javascript">
              function signup(event) {
                event.preventDefault(); // Prevent form from refreshing the page
                document.getElementById("error-message").style.visibility="hidden";
                const username = document.getElementById("username").value
                const password = document.getElementById("password").value
                const errorMessage = document.getElementById("error-message");
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
                    console.log("Response", response)
                    if (response.ok) {
                        errorMessage.style.visibility = "hidden";
                        window.location.href = "/users/welcome"
                        return
                    } else {
                        errorMessage.style.visibility = "visible";
                        response.json().then(errorResponse => {
                            console.log(errorResponse)
                            errorMessage.textContent = errorResponse.data.message
                        })
                    }


                })
                .catch(error => {
                    errorMessage.style.visibility="visible";
                    errorMessage.textContent = "An unexpected error occurred.";
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
        <p id="error-message" style="color:red; visibility: hidden;">
              An error occurred
        </p>
{{end}}


