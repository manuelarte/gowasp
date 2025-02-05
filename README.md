# Gowasp

Example of vulnerable application in Go.
To run the app, in the root directort type
> go run ./cmd/gowasp/.

## Vulnerabilities

Let's explore different vulnerabilities by checking some of the functionality that the app provides:

### 1. Create User - POST /users

We are going to explote the vulnerabilities related to the endpoint to create an user in [/users](http://localhost:8080/users).
The vulnerabilities that we are going to check are:
+ [Weak Password Requirements](https://cwe.mitre.org/data/definitions/521.html)
+ [Weak Hash Algorithm](https://cwe.mitre.org/data/definitions/328.html)

An HTTP client is provided in [create-users.http](./tools/create-users.http) to follow along.

Let's first add 

#### Weak Password Requirements

As you can see in [users_service.go](./internal/services/user_service.go), the only requirement for a password is to have *more than 4 characters* (#Scenario 1).
Let's try to improve that by adding **stronger requirements**:
+ minimum 8 characters, (let's set also a maximum password length of 256)
+ include non-alphanumerical characters

Once you have implemented this restrictions, test them using the http client.

#### Weak Hash Algorithm

Explanation https://knowledge-base.secureflag.com/vulnerabilities/broken_cryptography/weak_hashing_algorithm_vulnerability.html
Run the http requests described in #Scenario 2. 
+ Get your md5 hashed password, and check how long does it take for a computer to decrypt it (https://10015.io/tools/md5-encrypt-decrypt#google_vignette) 
+ Check that password with same value generate the same hash.. 
To solve it, the best solution is to use up to date hashing algorithms, like bcrypt, scrypt or PBKDF2.
+ Change the MD5 hashing algorithm for [pbkdf2](https://pkg.go.dev/golang.org/x/crypto/pbkdf2).
+ Check that the same password generates different hashes.

Check that if you try to generate two users with the same password, you don't get the same hash

## TODO

### NEXT VULNERABILITIES, SQL Injection for login, for example, add posts for other users

### CSRF, SSRF, wrong authorization to get user's posts

### Excessive logging
...