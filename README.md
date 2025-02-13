# Gowasp

Example of vulnerable application written in Go.
To run the app, in the root directory type

> go run ./cmd/gowasp/.

## Vulnerabilities

Let's explore different vulnerabilities by exploiting some of the functionalities that the app provides:

### 1. Create User - POST /users/signup

We are going to explote the vulnerabilities related to the endpoint to create a user in [/users/signup](http://localhost:8080/users/signup).
The vulnerabilities that we are going to check are:
+ [Weak Password Requirements](https://cwe.mitre.org/data/definitions/521.html)
+ [Weak Hash Algorithm](https://cwe.mitre.org/data/definitions/328.html)

An HTTP client is provided in [users-signup.http](./tools/users-signup.http) to follow along.

#### Weak Password Requirements

As you can see in [users_service.go](./internal/services/user_service.go), the only requirement for a password is to have *more than 4 characters* (`#Scenario 1`).
Let's try to improve that by adding **stronger requirements**:
+ minimum 8 characters, (let's set also a maximum password length of 256)
+ include non-alphanumerical characters

Once you have implemented these restrictions, test them using the http client.

#### Weak Hash Algorithm

A detailed explanation of this vulnerability can be found [here](https://knowledge-base.secureflag.com/vulnerabilities/broken_cryptography/weak_hashing_algorithm_vulnerability.html)
Run the http requests described in [#Scenario 2](./tools/users-signup.http). 
+ Get your MD5 hashed password, and check how long does it take for a computer to decrypt it (https://10015.io/tools/md5-encrypt-decrypt#google_vignette) 
+ Check that password with same value generate the same hash.

> [!IMPORTANT]  
> Never use outdated hashing algorithms like MD5.

To solve it, the best solution is to use up-to date hashing algorithms, like `bcrypt`, `scrypt` or `PBKDF2`.
+ Change the MD5 hashing algorithm for [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt).
+ Use a different salt for every user.
+ Repeat `#2` and check that the same password generates different hashes.

### 2. Login User - POST /users/login

We are going to explote the vulnerabilities related to the endpoint to login a user in [/users/login](http://localhost:8080/users/login).
The vulnerabilities that we are going to check are:

+ [SQL injection](https://owasp.org/www-community/attacks/SQL_Injection)

An HTTP client is provided in [users-login.http](./tools/users-login.http) to follow along.

#### SQL Injection

As you can see in [user_repository.go](./internal/repositories/user_repository.go), in the `Login` method, the query is created by string concatenation.
Try to explote this vulnerability by concatenating an always true sql statement (something like -OR '1'='1'-), and avoid the execution of the password clause (maybe by commenting the rest of the query with --)

### 3. View Blogs

Once you're logged in, you are redirected to http://localhost:8080/users/welcome. There you can see an Intro Blog. 
If you open the network tab of developer console of your web browser, and refresh, the program makes a call to http://localhost:8080/blogs?name=intro.txt.

Let's try to exploit that. The vulnerabilities that we are going to check in this scenario:

+ [SSRF](https://owasp.org/Top10/A10_2021-Server-Side_Request_Forgery_%28SSRF%29/)

To follow along, check [blogs.http](./tools/blogs.http)

#### SSRF - Server Side Request Forgery

Let's check how the `GetBlogFileByName` method is implemented in [blogs_handler](/internal/handlers/blogs_handler.go).
We can see that we are using `os.Open`:
> file, err := os.Open(fmt.Sprintf("./resources/blogs/%s", name))

What would happen in we change the name query parameter to point to a different file in a different location, maybe we could try with `../internal/private.txt`
Try to also display `/etc/passwd` file content.

## TODO

### NEXT CSRF, html injection

### Excessive logging
...
