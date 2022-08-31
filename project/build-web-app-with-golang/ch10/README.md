# I18N and L10N

## 10.1 Setting the default region

### From the domain name

```go
if r.Host == "www.linonon.com"{
    i18n.SetLocale("en")
} else if r.Host == "www.linonon.cn" {
    i18n.SetLocale("zh-CN")
} else if r.Host == "www.linonon.tw" {
    i18n.SetLocale("zh-TW")
}
```
or use sub-domain such as "en.linonon.com"

```go
prefix := strings.Split(r.Host, ".")

if prefix[0] == "en"{
    i18n.SetLocale("en")
} else if prefix[0] == "cn"{
    i18n.SetLocale("zh-CN")
} else if prefix[0] == "tw"{
   i18n.SetLocale("zh-TW") 
}
```

### From URL parameters

`www.linonon.com/hello?locale=zh` or `www.linonon.com/zh/hello` ==> Set the region in Go: `i18n.SetLocale(params["locale"])`.

```go
router.GET("/:locale/books", listbook)

func listbook(c *gin.Context) {
    locale := c.Param("locale")
    // ...
}
```
### From the client settings area

HTTP.Headers: `Accept-Language`

```go
AL := r.Header.Get("Accept-Language")
if AL == "en" {
    i18n.SetLocale("en")
} else if AL == "zh-CN" {
    i18n.SetLocale("zh-CN")
} else if AL == "zh-TW" {
    i18n.SetLocale("zh-TW")
}
```

## 10.2 Localized Resources

### Localized textual content

Method: `map[key]value`
```go 
var locales map[string]map[string]string

func msg(locale, key string) string{
    if v, ok := locales[locale]; ok{
        if v2, ok := v[key]; ok{
            return v2
        }
    }
    return ""
}
```

```go
en["how old"] = "I am %d years old"
cn["how old"] = "我今年%d歲了"

fmt.Printf(msg(language, "how old"),30) 
// msg(language, "how old") == "I am %d years old"
```

### Localized date and time

- Time zone
```go
en["time_zone"] = "America/Chicago"
ch["time_zone"] = "Asia/Shanghai"

loc, _ := time.LoadLocation(msg(lag,"time_zone"))
t := time.Now()
t = t.In(loc)
fmt.Println(t.Format(time.RFC3339))
```
- Date
```go
en["date_format"] = "%Y-%m-%dT%H:%M:%S"
cn["date_format"] = "%Y年%m月%d日 %H時%M分%S秒"

fmt.Println(date(msg(lang,"date_format"),t))

func date(format string, time.Time) string{
    year, month, day = t.Date()
    hour, min, sec = t.Clock()
}
```

### Localized currency value
- Money
```go
en["money"] = "USD %d"
cn["money"] = "￥%d元"

fmt.Println(money_format(msg(lang,"money"),100))

func money_format(format string, money int64) string{
    return fmt.Sprintf(format, money)
}
```

### Localization of views and resources

```go
views
|--en  //英文模板
	|--images     //存储图片信息
	|--js         //存储JS文件
	|--css        //存储css文件
	index.tpl     //用户首页
	login.tpl     //登陆首页
|--zh-CN //中文模板
	|--images
	|--js
	|--css
	index.tpl
	login.tpl
```

With this directory structure, we can render locale-specific views like so.
```go 
s1, _ := template.ParseFiles("views/" + lang + "/index.tpl")
VV.Lang = lang
s1.Execute(os.Stdout, VV)
```

## 10.3 International sites

### Managing multiple local packages