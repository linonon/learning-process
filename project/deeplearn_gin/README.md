# DeepLearn Gin

## Second
- GET
- POST
- PATCH
- PUT
- DELETE

### Third
- 路由參數
    - `:` 通過 `c.Param()` 來獲得定義路由的值
        - `/user/:id` 與 `/user/list` 衝突，因為路由必須要唯一。
    - `*` 不常用
        - `/user/*id/` 用 `c.Param("id")` 獲得的是 `/id`
        - 訪問`/users`會被重定向到 `/users/`

### Fourth

### Fifith

### Sixth

### Seventh
- Group 分組路由
```go 
r.Group("/v1", ...HandlerFunc) // 執行這個路由下到子路由也會調用到這個middelware
{
    r.GET("/user", func(c *gin.Context){
        c.String(200,"/v1/users")
    })
}
```
- Group.Group 套用
