# 11 Error Handling, Debugging, ant Testing


```go
func Sqrt() {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // implementation
}
```

```go
type SyntaxError struct{
    msg string
    Offset int64
}

func (e *SyntaxError) Error() string { return e.msg }

funt main(){
    //... implementation
    if err := dec.Decode(&var); err != nil {
        if serr, ok := err.(*json.SyntaxError); ok{
            line, col := findLine(f, serr.Offset)
            return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
        }
        return err
    }
    // ... implementation
}
```
- [Why is my nil error value not equal to nil?](https://go.dev/doc/faq#nil_error)

```go
func Decode() *SyntaxError {
    var err *SyntaxError
    if IsBad() {
        err = &SyntaxError{}
    }

    return err // FIXME: err{T = type,  V = nil} != nil
}
```

Moer complicated error handling

```go
package net

type Error interface {
    error
    IsTimeout() bool
    IsTemporary() bool
}
```

Using type assertion to check whether or not our error is of type `net.Error`

```go
if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
    time.Sleep(time.Second)
    continue
}
if err != nil {
    log.Fatal(err)
}
```

### Error handling

- Normal version: such verbose
```golang
func init() {
    http.HandleFunc("/view", viewRecord)
}

func viewRecord(w http.ResponseWriter, r *http.Request) {
     c := appengine.NewContext(r)
     key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
     record := new(Record)
     if err := datastore.Get(c, key, record); err != nil {
         http.Error(w, err.Error(), 500)
         return
     }
     if err := viewTemplate.Execute(w,record); err != nil {
         http.Error(w, err.Error(), 500)
     }
}

```

- Error router version: reduce code
```go
type appHandler func(w http.ResponseWriter, r *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if err := fn(w, r); err != nil {
        http.Error(w, err.Error(), 500)
    }
}

func init() {
    http.HandleFunc("/view", appHandler(viewRecord))
}

func viewRecord(w http.ResponseWriter, r *http.Request) error {
    c := appengine.NewContext(t)
    key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
    record := new(Record)
    if err := datastore.Get(c, key, record); err != nil {
        return err
    }

    return viewTemplate.Execute(w, record)
}
```

More friendly version: custom error code & messages
```go
type appError struct {
    Error error
    Message string
    Code int
}

type appHandler func(w http.ResponseWriter, r *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if e := fn(w, r); e != nil {
        c := appengine.NewContext(r)
        c.Errorf("%v", e.Error)
        http.Error(w, e.Message, e.Code)
    }
}

func viewRecord(w http.ResponseWriter, r *http.Request) *appError {
    c := appengine.NewContext(r)
    key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
    record := new(Record)
    if err := datastore.Get(c, key, record); err != nil {
        return &appError{err, "Record note found", 404}
    }
    if err := viewTemplate.Execute(w, record); err != nil {
        return &appError{err, "Can't display record", 500}
    }

    return nil
}
```