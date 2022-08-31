# Interface
- 使用Interface 實現 retriefer() 功能。
- 使用者定義方法

## Interface conversions and type assertions
例子：
- `main.go`
    - Interface conversions
        - `func inspect(r Retriever)`
    - type assertions
        - `realRetriever` 變量

- Interface:
    - 實現者的 Type
    - 實現者 Value
    - 不要用interface變量的地址
    - 變量自帶pointer
    - 值傳遞，所以不需要使用到interface pointer
    - pointer 接收者只能用pointer 方式使用；value接收者都可以
    - `interface{}` 可代表任何類型
    - `[]interface.(int)` 表示強制轉換成int

- Interface 的 Composition
    - `main.go` -> `RetrieverPoster`