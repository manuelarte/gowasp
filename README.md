# GOwasp

![version](https://img.shields.io/github/v/release/manuelarte/gowasp)

GOwasp simulates a vulnerable web application built with Go. 
It showcases some of the most common security flaws found in modern web applications, based on the [OWASP Top 10](https://owasp.org/www-project-top-ten/) list.

The project encourages hands-on learning: 

1. Exploit each vulnerability
2. Understand the risk
3. Apply the fix.

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

Once the application is up and running, you can begin by exploring its core features.

### 🔑 Login/Signup Page

Start by navigating to the [signup page][signup]. Try creating a user account with the following credentials:

```yaml
username: test
password: test
```

### 🏠 Welcome Page

After logging in or creating an account, you’ll be redirected to the [welcome page][welcome].
This page displays an introductory blog post along with links to the latest entries.
Click on one of the recent posts to continue exploring.

### 📝 Post Page

Clicking a post takes you to its detail page, where you can read the content, view existing comments, and submit your own.

> [!NOTE]  
> Try to submit a comment like:
> 
> *Very nice post!*

Now that you've explored the basic functionality, it is time to dive into the fun part: *hacking the application*.

## ☣️ Vulnerabilities

Let's explore different vulnerabilities by exploiting some of the functionalities that this app provides:

### 1. Create User (POST /api/users/signup)

We'll start by exploring vulnerabilities in the [/api/users/signup][signup] endpoint.
Specifically, you'll investigate the following issues:

+ [Weak Password Requirements](#-weak-password-requirements)
+ [Weak Hash Algorithm](#-weak-hash-algorithm)
+ [Mass Assignment](#-mass-assignment)

> [!TIP]
> Use the provided HTTP client file [users-signup.http](./tools/users-signup.http), to follow along and test each case.

#### 🔐 [Weak Password Requirements](https://cwe.mitre.org/data/definitions/521.html)

In [users/service.go](./internal/users/service.go), the password policy currently allows any value with *more than four characters* (`#1. Scenario`).

To strengthen this, update the logic to enforce the following rules:

+ Require a minimum of 8 characters (and a maximum of 256).
+ Include at least one non-alphanumeric character.

After applying these changes, verify the new password validation behavior.

#### 🤖 [Weak Hash Algorithm](https://cwe.mitre.org/data/definitions/328.html)

For a detailed explanation of this vulnerability, see the [Weak Hashing Algorithm Vulnerability](https://knowledge-base.secureflag.com/vulnerabilities/broken_cryptography/weak_hashing_algorithm_vulnerability.html).

To explore this issue, follow the steps in [#2. Scenario](./tools/users-signup.http):

+ Submit a signup request and extract the generated MD5 hash.
+ Use a tool like [md5-encrypt-decrypt](https://10015.io/tools/md5-encrypt-decrypt#google_vignette) to estimate how quickly an attacker can reverse the hash.
+ Confirm that hashing the same password always produces the same result.

> [!IMPORTANT]  
> Avoid outdated hashing algorithms like MD5. They offer weak protection against brute-force attacks.

To improve password security, the best solution is to use up-to date hashing algorithms, like `bcrypt`, `scrypt` or `PBKDF2`.

+ Replace MD5 with [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt).
+ Generate a unique salt for each user.
+ Re-run `#2 Scenario` and verify that the same password produces different hashes.

#### 📝 [Mass Assignment](https://www.veracode.com/security/dotnet/cwe-915/)

For a detailed explanation of this vulnerability, see the [OWASP Mass Assignment Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Mass_Assignment_Cheat_Sheet.html).

We are going to exploit the vulnerability related to the API endpoint [/api/users/signup][signup].

When inspecting the login response, you’ll notice a field named `isAdmin`. 
The HTML signup form doesn’t expose this field, but what happens if you *include it directly in the API request?*
Try sending a crafted request that sets `isAdmin` to true and observe the outcome.

### 2. Login User (POST /users/login)

We are going to exploit the vulnerabilities related to the endpoint to log in a user in [/api/users/login][login].
The vulnerability that we are going to check is:

+ [SQL injection](#-sql-injection)

An HTTP client is provided in [users-login.http](./tools/users-login.http) to follow along.
As you can see in [users repository.go](./internal/users/repository.go), in the `Login` method, the query is created by string concatenation.

#### 💉🛢 [SQL injection](https://owasp.org/www-community/attacks/SQL_Injection)

For a detailed explanation of this vulnerability, see the [OWASP SQL Injection Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/SQL_Injection_Prevention_Cheat_Sheet.html).

Try to exploit this query concatenation by concatenating an `always true` SQL statement (something like `-OR '1'='1'-`).
The goal is to avoid the execution of the password clause (maybe by injecting a comment (`--`) to comment out the rest of the query)

> [!IMPORTANT]  
> Never concatenate strings in a query.

### 3. View Posts

Once you're logged in, you are redirected to the [Welcome][welcome] page. There you can see an **Intro Post**.

The vulnerabilities that we are going to check in this scenario:

+ [SSRF](#-ssrf---server-side-request-forgery)

To follow along, check [posts.http](./tools/posts.http)

#### 📥 [SSRF - Server Side Request Forgery](https://owasp.org/Top10/A10_2021-Server-Side_Request_Forgery_%28SSRF%29/)

If you open the network tab `developer console` in your web browser (by default F12), and **refresh the welcome page**, the program makes a call to [/posts?name=intro.txt](http://localhost:8080/posts?name=intro.txt).
Let's check how the `GetStaticPostFileByName` method is implemented in [posts controller](./internal/api/rest/posts.go).
We can see that we are using [`os.Open`](https://pkg.go.dev/os#Open):

```go
file, err := os.Open(fmt.Sprintf("./resources/posts/%s", name))
```

+ What would happen in we change the name query parameter to point to a different file in a different location?, maybe we could try with `../internal/private.txt`
+ Try to also display `/etc/passwd` file content.

To solve this issue for this scenario, we could validate the user input, and avoid path traversal with functions like [`os.Root`](https://pkg.go.dev/os#Root)

> [!IMPORTANT]  
> Always validate user input.

### 4. Add Comments

The vulnerabilities we are going to check here:

+ [Broken Access Control](#-broken-access-control)
+ [CSRF](#-csrf---cross-site-request-forgery)
+ [HTML Template injection](#-html-template-injection)

#### 🩹 [Broken Access Control](https://owasp.org/Top10/A01_2021-Broken_Access_Control/)

For a detailed explanation of this vulnerability, see the [OWASP Authorization Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Authorization_Cheat_Sheet.html).

In Scenario 1 of the http tool [post_comments.http](/tools/post_comments.http), you can create a comment for a post.
However, the payload includes `postId` and `userId` fields that you can manipulate.
By modifying these values, you can create comments as any user on any post.

To fix this vulnerability, consider these approaches:

+ Override the `userId` and `postId` fields in the payload with trusted values (the user id coming from the session cookie and the `postId` coming from the url).
+ (**Preferred**) Define a dedicated struct that accepts only valid fields, similar to the `UserSignup` struct, to prevent unauthorized data injection.

#### 🔄 [CSRF - Cross Site Request Forgery](https://owasp.org/www-community/attacks/csrf)

For a detailed explanation of this vulnerability, see the [Cross Site Request Forgery Prevention Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html).

The post-comments endpoint lacks protection against CSRF attacks. Verify this by following these steps:

+ Log in to the application using your browser.
+ Open [price-win.html](/tools/price-win.html) in the same browser.
+ Click the rewards button.
+ Visit [/post/2/comments](http://localhost:8080/posts/2/comments) and observe the result.

This exploit combines vulnerabilities from both CSRF and HTML template injection.

Prevent CSRF attacks by adding a CSRF token cookie to requests and validating that the token in the JSON payload matches the cookie value.
In the template [add_edit_comment.tpl](/web/templates/posts/add_edit_comment.tpl) the token appears as:

```html
<input type='hidden' id='csrf' name="csrf" value='{{ .csrf }}'>
```

Implement validation to ensure the `csrf` value from the JSON payload matches the `csrf` cookie sent in the request.
After applying the fix, restart the application and confirm that creating comments via the `win price` button no longer works.

#### 💉🌐 [HTML Template Injection](https://owasp.org/www-project-web-security-testing-guide/latest/4-Web_Application_Security_Testing/11-Client-side_Testing/03-Testing_for_HTML_Injection)

Earlier, you encountered a template injection vulnerability.

Run the `#Scenario 2` HTTP requests that attempts to inject `<script>` tag in your comment.

To solve this, remember to always escape and validate user input to prevent injection attacks.
The [Gin Framework](https://gin-gonic.com/) provides built-in protection against this vulnerability.
In this project, a custom `unsafe` function bypassed Gin’s escaping to render raw HTML.
Check [`gowasp.main`](cmd/gowasp/gowasp.go) to see how `unsafe` function renders HTML content without escaping.

[signup]: http://localhost:8080/api/users/signup
[login]: http://localhost:8080/users/login
[welcome]: http://localhost:8080/users/welcome
