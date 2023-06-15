# Best Practices

## Interface

### Q

假如我有2個 interface:

```go
type best interface {
	Init() 
	Start() 
	Close() 
}
```

```go
type practices interface {
	s() string
	v() interface{}
}
```

然後有一個實現了這兩個 interface 的 struct:

```go
type impl struct {}

func Init()
func Start()
func Close() 
func s() string {"s"}
func v() interface {"v"}
```

為了功能模塊分割方便

我會希望這樣再創建兩個 struct, 給 impl 繼承

```go
type impl struct {
	*bestImpl
	*practicesImpl
}

type bestImpl struct {}

func (b *bestImpl) Init()
func (b *bestImpl)  Start()
func (b *bestImpl)  Close() 

type practicesImpl struct {}
func (p *practicesImpl) s() string {"s"}
func (p *practicesImpl) v() interface {"v"}
```

但是繼承容易造成代碼的臃腫, 而且不容易組件化, 請問有什麼好的修改意見?

### A

Go 語言並不直接支持 "類" 的繼承概念，但它支持接口和結構體 (struct) 的嵌入，這也是你現在用的方法。使用這種方法是一種有用的策略來組織和分離關注點，但是如果不當使用可能會導致結構變得複雜並且難以管理。

以下是一些可能的修改建議：

1. **減少 interface 的方法數量**：一個有效的 interface 是簡單且小型的。適當的接口應該僅包含一或兩個方法。如果你的接口包含很多方法，考慮是否可以拆分成更小的接口。

2. **單一職責原則 (Single Responsibility Principle)**：如果你的 `impl` struct需要實現太多不同的功能，那麼可能是設計上的問題。每個 struct 應該只有一個職責。如果一個 struct 需要執行太多工作，考慮是否可以拆分成多個 struct，每個 struct 只負責一種工作。

3. **嵌入與組合**：嵌入可以用來實現某種 "繼承"，但是組合更符合 Go 的設計哲學。在 Go 語言中，我們更傾向於創建小而簡單的部件，並使用組合將它們組裝成更複雜的結構。例如：

    ```go
    type bestImpl struct {}

    func (b *bestImpl) Init() {...}
    func (b *bestImpl) Start() {...}
    func (b *bestImpl) Close() {...}

    type practicesImpl struct {}

    func (p *practicesImpl) s() string {...}
    func (p *practicesImpl) v() interface {...}

    type impl struct {
        best *bestImpl
        practices *practicesImpl
    }
    ```

以上是一些通用的建議，實際的修改會根據你的實際需求而定。希望這些建議對你有所幫助！

### Q

感謝提醒, 那我這樣做是不是會更好:

```go
type bestImpl struct {}

func (b *bestImpl) Init() {...}
func (b *bestImpl) Start() {...}
func (b *bestImpl) Close() {...}

type practicesImpl struct {}

func (p *practicesImpl) s() string {...}
func (p *practicesImpl) v() interface {...}

type impl struct {
	best *bestImpl
	practices *practicesImpl
}

type (i *impl) BestImpl() *bestImpl {
	return i.best
}

type (i *impl) PracticesImpl() *practicesImpl {
	return i.practices
}
```

這樣就可以在我需要調用 impl 相應的模塊時, 調用相應的實現了?

### A

對的，這種方式可以在需要時讓你可以訪問 `impl` 中 `best` 和 `practices` 的實例並調用它們的方法。不過，需要注意的是在你的代碼中有一點小錯誤，你定義方法的語法有些問題，它應該如下：

```go
type impl struct {
	best *bestImpl
	practices *practicesImpl
}

func (i *impl) BestImpl() *bestImpl {
	return i.best
}

func (i *impl) PracticesImpl() *practicesImpl {
	return i.practices
}
```

此外，如果你希望 `impl` 仍然實現 `best` 和 `practices` 接口，你還需要為 `impl` 添加轉發函數（就像我在前一個回答中提到的那樣）。但是如果不需要 `impl` 去實現這些接口，則可以省略這一步。

這種方式確實提供了一種組織你的代碼的方式，讓你可以根據需要來訪問並使用 `bestImpl` 和 `practicesImpl` 的實例。這樣可以使你的組件更為模塊化，並更容易管理和維護。

希望這個解答能夠對你有所幫助！