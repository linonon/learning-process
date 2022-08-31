# 06-1

## Set cookie in Go
```go
http.SetCookie(w ResponseWriter, cookie *Cookie)
```
- See what cookie looks like:
```go
type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string

// MaxAge=0 means no 'Max-Age' attribute specified.
// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}
```
- Setting a cookie

```go
expiration := time.Now()
expiration = expiration.AddData(1, 0, 0)
cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
http.SetCookie(w, &cookie)
```

## Fetches cookies in Go
- Get cookies

```go
cookie, _ := r.Cookie("username")
fmt.Fprint(w, cookie)
```
- Anohter way to get

```go
for _, cookie := range r.Cookies(){
    fmt.Fprintf(w, cookie.Name)
}
```