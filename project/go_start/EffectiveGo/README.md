# Effective Go

## Commentary

### godoc

### package
- package的命名應該都為小寫
- 遇到 `ring.Ring` 這種情況, 可以用`ring.New`
- `once.Do(setup)` 比 `once.DoOrWaitUnitilDone(setup)` 好, 長命名還不如寫一份有用的說明文檔.

## Name

### Getter
- owner 字段
    - get: Owner()
    - set: SetOwner(user)

### Interface
- 只有一個方法的Interface命名應該用 `-er` 結尾, 如 `Reader` `Writer` `Formatter` `ClosNotifier` 等.
- 假如實現了一個`func A()`, 且A與知名方法B的含義相同, 則應該將 `func A()` --> `fun B()`

### MixedCaps
用 `MixedCaps` or `mixedCaps` 方式命名,不要用 `_`.

## Semicolon 分號
- 在GO中, 通常在 `for` 裡面才會用到 `;`
- 在control struct 中 `{` 不應該放在下一行

## Control Structures

### If
- `if` 強制 `{}`, 同一編碼風格
    - 可接受初始化語句
    - `if err := file.Chomd(0664); err != nil {...}` 初始化err

### Redeclaration and reassignment
- `f, err := os.Open(name)` & `d, err := f.Stat()` 
- 上面的情況中, `err` 是合法的, 只是re-assignment了
  
滿足下列條件時， 已聲明過的 `v` 可以出現在 `:=` 中
- `v` 在外層聲明過，則內部聲明會創建新的`v`
- 多值聲明時，至少另一個變量是新的聲明。

### For
- `for _, value := range array{...}` Blank identifier，第一個不讀
- `for` 中 多參數賦值時，需要平行賦值

### Switch
- `case ' ', '?': return true` case 可以用 `,` 列舉相同的處理條件
- 未完成

## Function

### Multiple return values
- `func name(b []byte, i int) (int int){.. return x, y}`
-  `func name(b []byte, pos int) (n int, err error) {...  return}` 直接 `return`則直接返回 `n` & `err`

### Defer
- 函數返回之前立即執行`defer`
- 典型用在 Mutex lock 或 file.Close() 上
- LIFO順序執行

## Data

### Allocation with new
- `new` 
    - 分配為零的存儲空間，並返回地址，指針類型

### Constructor and composite literals
```go
func NewFile(fd int, name string) *File{
    // ...
    f := File{fd, name,nil,0}
    return &f
}
```
也等於
```go
return &File{fd,name,nil,0}
```
定向賦值，沒賦值的則默認0或nil
```go
return  &File{fd:fd,name:name}
```

- `new(File)` 等價於 `&File{}`

### Allocation with make
```go
var p *[]int = nwe([]int)       // *p == nil, 少用
var v []int = make([]int, 100)  //  v slice 有100個 int 在裡面
// 常規定義方法
v := make([]int, 100) 
```
顯然，`make` 適用於 `maps` `slices` & `channels`類型，不返回 *p。如果需要返回 *p 請使用 `new`

### Arrays
- Pass with Value (copy for the array).
- size is part of **arrays** type

### Slices
- 相當於 *arrays
```go
func (f *File) Read(buf []byte) (n int, error)
```
返回讀取byte的數量和error的函數
```go
n, err := f.Read(buf[0:32])
```
取得 `buf` 中的前32byte

- `cap` 為 slices 定義了可變的最大長度，最長不會超過底層Array的長度。
- `append` 追加元素，但是超長時會創建新的函數。

### Two-dimensional slices
- Slices 會增長 或者 收縮時， slices
- 不會時，可單次分配一個大的，讓各個 slices 指向他，更高效。

- 按行分配

```go
// 底層 slices
picture := make([][]uint8, YSize)
for i := range picture {
    // 遍歷每行
    picture[i] = make([]uint8, XSize)
}
```

- 一次性分配

```go
// 底層 slices
picture := make([][]uint8, YSize)
pixels := make([]uint8, XSize* YSize)

// 遍歷圖片所有行，從
for i := range picture {
    picture[i], pixels = pixels[:XSize], pixels[XSize:]
}
```

### Maps
區分 某項不存在還是本身為0，可以用 `.. , ok = ..`
```go
seconds, ok := timeZone[tz]
```

### printing
- `Printf` 格式化string，通用格式 `%v`，`%+v` 輸出字段名，`%#v` GO語法輸出
- `Println` 參數之間會有 `" "`
- `Fprint` 實現了io.Writer 的第一個實參 

### append
- 內建函數，需要編譯器支持
- 必須輸出結果 `x = append(x, ...int)`

## Initialization

### Constants
- 只能是 numbers, characters(runes), strings or booleans.
- `enum` 用 `iota` 創建 [具體](https://learnku.com/docs/effective-go/2020/initialization/6244)

### Variables
可像cons創建
```go
var(
  home = ..
  uer = ..
  gopath = ..
)
```

### The init Function
- 除了不能表示成聲明的初始化外， `init` 還長在程序真正開始前執行，校驗或驗證程序的狀態。

## Methods 

### Pointers vs. Values
- pointer & interface 不能定義新方法

```go
type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte{
    l := len(slice)
    if l + len(data) > cap(slice){
        newslice := make([]byte, (l+len(data))*2)
        copy(newSlice,slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    copy(slice[l:],data)
    return slice
}
```
升級版(無需接收者)
```go
func (p *Byteslice) Append(data []byte){
    slice := *p
    // pass
    *p = slice
}
```

## Interfaces and other types

### Interfaces

### 可深入 ["Conversions" & "Interface conversions and type assertions"](https://learnku.com/docs/effective-go/2020/interface-and-other-types/6246)

### Interface & Method
- duck typing
    - 描述事物的外部行為而非內部結構
    - go -> 結構化類型系統

## Concurrency
- `main` 執行完畢後， goroutine會被kill
- 非搶佔式多任務處理，攜程主動交出控制權
- 輕量級“執行緒“
- 多個routine可以在多個執行緒上run。
- 子程序是Coroutine的特例
- goroutine 由調度器調整不同goroutine在什麼執行緒上操作。
- goroutine可能到切換點
    - I/O, select
    - channel
    - 等待鎖
    - 函數調用（maybe）
    - runtime.Gosched()
    - 以上都為參考
### Channel
`Don't communicate by sharing memory; share memory by communicating.`

#### `channel/done/done.go`
- 更改結構體
    - 創建了channel結構體，用來保存channel數據，worker.done 代表動作完成。
#### `channel/waitgroup/waitgroup.go`
- 使用 waitgroup
    - 記得用 * 類型，因為只有一個wg
- 類似傳統的傳送機制

#### `channel/select/select.go`
- select 選擇調度
- 通過通信來共享內存
- 定時器的使用
- select 中使用了Nil Channel，數據沒準備好的時候就不會部署Channel

#### `channel/mutex/mutex.go`
- mutex鎖使用
- `go run -race .` 檢查是否有衝突
- 系統也有提供原子化操作