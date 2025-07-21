# GOwasp

![version](https://img.shields.io/github/v/release/manuelarte/gowasp)

GOwasp is a deliberately vulnerable web application written in [Go](https://go.dev/). This project demonstrates some of the most common security vulnerabilities affecting web applications today, based on the [OWASP](https://owasp.org/) top 10. The goal is to learn security concepts by exploiting and then fixing these vulnerabilities.

## 🚀Getting Started

To run the app, in the root directory type:

```bash
make r
```

If you want to run it with Docker 🐳 use:

```bash
make dr
```

## 🛠️ Application Overview

Once you have the application running, let's start first exploring the functionality provided by it.

### 🔑 Login/Signup Page

Navigate to the signup form [/users/signup][signup]. Let's try to create an user, for example, with:

```yaml
username: test
password: test
```

### 🏠 Welcome Page

After you login or create an account, you are redirected to the [welcome][welcome] page. In this page you can see an intro blog, and also links to the latest posts. Let's click in one of the latest posts.

### 📝 Post Page

Clicking on a post takes you to its details page, where you can read the post, view comments, and submit your own comment. Try to submit a comment like:

*Very nice post!*

After this small guide, let's now try to hack this application.

## ☣️ Vulnerabilities

Let's explore different vulnerabilities by exploiting some of the functionalities that this app provides:

### 1. Create User - POST /users/signup

We are going to explote the vulnerabilities related to the endpoint to create a user in [/users/signup][signup].
The vulnerabilities that we are going to check are:

+ [Weak Password Requirements](https://cwe.mitre.org/data/definitions/521.html)
+ [Weak Hash Algorithm](https://cwe.mitre.org/data/definitions/328.html)

An HTTP client is provided in [users-signup.http](./tools/users-signup.http) to follow along.

#### 🔐 Weak Password Requirements

As you can see in [users_service.go](./internal/services/user_service.go), the only requirement for a password is to have *more than 4 characters* (`#1. Scenario`).
Let's try to improve that by adding **stronger requirements**:

+ minimum 8 characters, (let's set also a maximum password length of 256)
+ include non-alphanumerical characters

Once you have implemented these restrictions, test them using the http client.

#### 🤖 Weak Hash Algorithm

A detailed explanation of this vulnerability can be found [here](https://knowledge-base.secureflag.com/vulnerabilities/broken_cryptography/weak_hashing_algorithm_vulnerability.html).
Run the http requests described in [#2. Scenario](./tools/users-signup.http) and:

+ Get the generated MD5 hashed password, and check how long does it take for a computer to decrypt it (e.g. https://10015.io/tools/md5-encrypt-decrypt#google_vignette)
+ Check that same password generates the same hash.

> [!IMPORTANT]  
> Never use outdated hashing algorithms like MD5.

To solve it, the best solution is to use up-to date hashing algorithms, like `bcrypt`, `scrypt` or `PBKDF2`.

+ Change the MD5 hashing algorithm for [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt).
+ Use a different salt for every user.
+ Repeat `#2 Scenario` and check that the same password generates different hashes.

#### 📝 Mass Assignment

A detailed explanation of this vulnerability can be found [here](https://cheatsheetseries.owasp.org/cheatsheets/Mass_Assignment_Cheat_Sheet.html).
We are going to explote the vulnerability related to the API endpoint [/users/signup][signup].

If you check the login endpoint, we see that we are returning a field called `isAdmin`. That field is not available in the html form.
But what about if we use the API endpoint.

### 2. Login User - POST /users/login

We are going to explote the vulnerabilities related to the endpoint to log in a user in [/users/login][login].
The vulnerability that we are going to check is:

+ [SQL injection](https://owasp.org/www-community/attacks/SQL_Injection)

An HTTP client is provided in [users-login.http](./tools/users-login.http) to follow along.
As you can see in [user_repository.go](./internal/repositories/user_repository.go), in the `Login` method, the query is created by string concatenation.

#### 💉🛢 SQL Injection

Try to explote this query concatenation by concatenating an `always true` sql statement (something like `-OR '1'='1'-`).
The goal is to avoid the execution of the password clause (maybe by injecting a comment (`--`) to comment out the rest of the query)

> [!IMPORTANT]  
> Never concatenate strings in a query.

### 3. View Posts

Once you're logged in, you are redirected to the [Welcome][welcome] page. There you can see an **Intro Post**.

The vulnerabilities that we are going to check in this scenario:

+ [SSRF](https://owasp.org/Top10/A10_2021-Server-Side_Request_Forgery_%28SSRF%29/)

To follow along, check [posts.http](./tools/posts.http)

#### 📥 SSRF - Server Side Request Forgery

If you open the network tab of the developer console of your web browser (normally F12), and **refresh the welcome page**, the program makes a call to [/posts?name=intro.txt](http://localhost:8080/posts?name=intro.txt).
Let's check how the `GetStaticPostFileByName` method is implemented in [posts_handler](./internal/handlers/posts_handlers.go).
We can see that we are using [`os.Open`](https://pkg.go.dev/os#Open):
> file, err := os.Open(fmt.Sprintf("./resources/posts/%s", name))

+ What would happen in we change the name query parameter to point to a different file in a different location?, maybe we could try with `../internal/private.txt`
+ Try to also display `/etc/passwd` file content.

To solve this issue for this scenario we could validate the user input, and avoid path traversal with functions like [`os.Root`](https://pkg.go.dev/os#Root)

> [!IMPORTANT]  
> Always validate user input.

### 4. Add Comments

The vulnerabilities we are going to check here:

+ Broken Access Control
+ [CSRF](https://owasp.org/www-community/attacks/csrf)
+ HTML Template injection

#### 🩹 Broken Access Control

If we look at the Scenario 1 in the http tool [post_comments.http](/tools/post_comments.http), we can see that we can create a comment for a post.
But if we take a look at the payload, we can see that the `postID` and the `userID` are sent as part of the payload.
We can manipulate these values and check that we can create comments for any user to any post.

There are several ways to implement a solution for this vulnerability in this case:

+ Override the values given in `userID` and/or `postID` by the proper values (the user id coming from the session cookie and the postID coming from the url)
+ (**Preferred**) Implement a new struct that contains only the valid fields as we have in `UserSignup` struct.

#### 🔄 CSRF - Cross Site Request Forgery

The add comments endpoint is not protected against CSRF attacks. And we can check it by following this steps:

+ login in the application with your browser.
+ open [price-win.html](/tools/price-win.html) with that same browser.
+ Click on the rewards button.
+ Go to `http://localhost:8080/posts/2/comments` and check what happened.

The app has been exploit by two vulnerabilities, CSRF and HTML Template injection.

To avoid CSRF attacks we can validate add a CSRF cookie in our requests, and validate in the payload that the cookie and the json field match.
In the template [add_edit_comment.tpl](/web/templates/posts/add_edit_comment.tpl) you can check that we are sending a csrf value in:

```<input type='hidden' id='csrf' name="csrf" value='{{ .csrf }}'>```

Validate that the value that we receive from that json field matches the value that we have in the `csrf` cookie in the Request.
Restart the application and check that you can't create comments anymore using the `win price` button.

#### 💉🌐 HTML Template Injection

Before we could see that we also suffered from template injection.

Run the `#Scenario 2` http requests that tries to inject a `<script>` content in your comment.

To solve this remember to always escape/validate user input.
In this case, [Gin](https://gin-gonic.com/) provides already a mechanism against this attack, and we needed to avoid it by creating a custom function to avoid escaping the html characters.
You can check [`gowasp.main`](cmd/gowasp/gowasp.go) how I created an `unsafe` function to render html content.

## 📝 TODO

+ Improve the css.
+ Avoiding brute force attacks.
+ Insufficient logging when adding the comment.

[signup]: http://localhost:8080/users/signup
[login]: http://localhost:8080/users/login
[welcome]: http://localhost:8080/users/welcome
