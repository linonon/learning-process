# Text files

## XML

## JSON
### Parse to interface

JSON: Use `map[string]interface{}` && `[]interface{}` to save all JSON objects and array.
- `bool` represents `JSON booleans`
- `float64` represents `JSON numbers`
- `string` represents `JSON strings`
- `nil` represents `JSON null`

package `simplejson`:
```go
js, err := NewJson([]byte(`{
    "test":{
        "array": [1,"2",3]
        "int": 10
        "float": 3.14
        "bignum": 29929293920930920398,
        "string": "simplejson"
        "bool": true
    }
}`))

arr, _ := js.Get("test").Get("array").Array()
i, _ := js.Get("test").Get("int").Int()
ms, _ := js.Get("test").Get("string").MustString()
```

### Producing JSON

JSON `struct Tag`: For lower case letter.
```go
type Server struct{
    ServerName string `json:"serverName"`
    ServerIP string `json:"serverIP"`
}
```
- Field tags containing `"-"` will not be outputted
- `"omitempty"`, will not be outputted if empty.
- `",string"`, Go converts this field to its co

## Regexp

Verify an IP address:
```go
func IsIP(ip string) (m bool){
    // 127.0.0.1
    m, _ = regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip)
}
```

Verify whether user input is valid:
```go
func main(){
    if len(os.Args) == 1{
        fmt.Println("Usage: regexp [string]")
        os.Exit(1)
    }

    if m, _ := regexp,MatchString("^[0-9]+$", os.Args[1]); m {
        fmt.Println("Is Number")
    } else {
        fmt.Println("Is Not Number")
    }
}
```

## Templates

## Files

### Directories
Code sample:
```go
func main(){
    os.Mkdir("astaxie", 0777)
    os.MkdirAll("astaxie/test1/test2", 0777)
    err := os.Remove("astaxie")
    if err != nil {
        fmt.Println("err")
    }
    os.RemoveAll("astaxie")
}
```
### File

#### Write
```go
func main(){
    userFile := "astaxie.txt"
    fout, err := os.Create(userFile)
    if err != nil{
        fmt.Println(userFile, err)
        return
    }
    defer fout.Close()
    for i:=0; i < 10 ; i++{
        fout.WriteString("Just a test!\r\n")
        fout.Write([]byte("Just a test!\r\n"))
    }
}
```
#### Read
```go
func main(){
    userFile := "astaxie.txt"
    fl, err := os.Open(userFile)
    if err != nil{
        fmt.Println(userFile, err)
        return
    }
    defer fl.Close()
    buf := make([]byte, 1024)
    for {
        n, _ := fl.Read(buf)
        if 0 == n{
            break
        }
        os.Stdout.Write(buf[0:n])
    }
}
```

## Strings

- func Contains(s, substr string) bool
```go
strings.Contain("seafood", "foo")
// true
```
- func Join(a []string, sep string) string
```go
s := []string{"foo", "bar", "baz"}
fmt.Println(strings.Join(s, ", "))
// foo, bar, baz
```
- `func Index(s, sep string) int`
- `func Repeat(s string, count int) string`
- `func Replace(s, old, new string, n int) string`
- `func Split(s, sep string) []string`

```go 
strings.Split("a,b,c", ",") // ["a", "b", "c"]
strings.Split("a man a plan a canal panama", "a ") // ["" "man " "plan " "canal panama"]
```
- `func Trim(s string, cutset string) string`
``` go
strings.Trim(" !!! Achtung !!! ", "! ")
// Achtung
```
- package `strconv`
