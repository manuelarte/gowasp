# Gowasp

Example of vulnerable application written in Go.
To run the app, in the root directory type

> go run ./cmd/gowasp/.

## Vulnerabilities

Let's explore different vulnerabilities by exploiting some of the functionalities that this app provides:

### 1. Create User - POST /users/signup

We are going to explote the vulnerabilities related to the endpoint to create a user in [/users/signup](http://localhost:8080/users/signup).
The vulnerabilities that we are going to check are:
+ [Weak Password Requirements](https://cwe.mitre.org/data/definitions/521.html)
+ [Weak Hash Algorithm](https://cwe.mitre.org/data/definitions/328.html)

An HTTP client is provided in [users-signup.http](./tools/users-signup.http) to follow along.

#### Weak Password Requirements

As you can see in [users_service.go](./internal/services/user_service.go), the only requirement for a password is to have *more than 4 characters* (`#1. Scenario`).
Let's try to improve that by adding **stronger requirements**:
+ minimum 8 characters, (let's set also a maximum password length of 256)
+ include non-alphanumerical characters

Once you have implemented these restrictions, test them using the http client.

#### Weak Hash Algorithm

A detailed explanation of this vulnerability can be found [here](https://knowledge-base.secureflag.com/vulnerabilities/broken_cryptography/weak_hashing_algorithm_vulnerability.html)
Run the http requests described in [#2. Scenario](./tools/users-signup.http) and: 
+ Get the generated MD5 hashed password, and check how long does it take for a computer to decrypt it (e.g. https://10015.io/tools/md5-encrypt-decrypt#google_vignette) 
+ Check that password with same value generate the same hash.

> [!IMPORTANT]  
> Never use outdated hashing algorithms like MD5.

To solve it, the best solution is to use up-to date hashing algorithms, like `bcrypt`, `scrypt` or `PBKDF2`.
+ Change the MD5 hashing algorithm for [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt).
+ Use a different salt for every user.
+ Repeat `#2 Scenario` and check that the same password generates different hashes.

### 2. Login User - POST /users/login

We are going to explote the vulnerabilities related to the endpoint to log in a user in [/users/login](http://localhost:8080/users/login).
The vulnerability that we are going to check is:

+ [SQL injection](https://owasp.org/www-community/attacks/SQL_Injection)

An HTTP client is provided in [users-login.http](./tools/users-login.http) to follow along.
As you can see in [user_repository.go](./internal/repositories/user_repository.go), in the `Login` method, the query is created by string concatenation.

#### SQL Injection

Try to explote this query concatenation by concatenating an `always true` sql statement (something like -OR '1'='1'-), and avoid the execution of the password clause (maybe by commenting the rest of the query with --)

### 3. View Blogs

Once you're logged in, you are redirected to http://localhost:8080/users/welcome. There you can see an Intro Blog.

The vulnerabilities that we are going to check in this scenario:

+ [SSRF](https://owasp.org/Top10/A10_2021-Server-Side_Request_Forgery_%28SSRF%29/)

To follow along, check [blogs.http](./tools/blogs.http)

#### SSRF - Server Side Request Forgery

If you open the network tab of the developer console of your web browser (F12 in Chrome), and refresh the welcome page, the program makes a call to http://localhost:8080/blogs?name=intro.txt.
Let's check how the `GetBlogFileByName` method is implemented in [blogs_handler](/internal/handlers/blogs_handler.go).
We can see that we are using `os.Open`:
> file, err := os.Open(fmt.Sprintf("./resources/blogs/%s", name))

+ What would happen in we change the name query parameter to point to a different file in a different location?, maybe we could try with `../internal/private.txt`
+ Try to also display `/etc/passwd` file content.

To solve this issue for this scenario we could validate the user input, and avoid path traversal with functions like [`filepath.Clean`](https://pkg.go.dev/path/filepath#Clean)

### 4. Add Comments

Vulnerabilities we are going to check here:

- [CSRF](https://owasp.org/www-community/attacks/csrf)
- Template injection
- Improper validation (I can create comments for other users)

#### CSRF - Cross Site Request Forgery

We are going to explode the feature of adding comments

## TODO

### NEXT CSRF, html injection

### Excessive logging
...
