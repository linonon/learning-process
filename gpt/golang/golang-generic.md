# Golang Generic

## Q

教我 Golang 的 generic , 給出有意義的 Coding 例子, 最好能用在工作上

## A

在 Go 1.18 之前，Go並無內建的泛型功能，在這個版本之前通常會透過 interface{} 來繞過這個問題，不過這樣会牺牲了一些类型安全。但在 Go 1.18 裡面，官方終於引入了泛型，在這章節裡我們會教你如何於 Go 裡使用泛型。

以下是一個實例，在這個例子中，定義了一個可接受任何數Type的Slice和一個比較函數，返回Slice中的最大值。

首先，我們定義一個泛型比較函數，用於比較兩個素是否符合泛型約束：

```go
type Comparer[T any] func(x, y T) bool
```

接著定義最大值函數，該函數接受一個數組和一個比較函數：

```go
func Max[T any](s []T, cmp Comparer[T]) T {
    max := s[0]
    for _, v := range s[1:] {
        if cmp(max, v) {
            max = v
        }
    }
    return max
}
```

然後我們就可以在函數裡使用這個泛型函數來找尋任何數类型Slice的最大值。

```go
func main() {
    ints := []int{1, 2, 3, 4, 5}
    fmt.Println(Max(ints, func(x, y int) bool { return x < y })) // prints 5

    strings := []string{"a", "b", "c", "d", "e"}
    fmt.Println(Max(strings, func(x, y string) bool { return x < y }))  // prints "e"
}
```
稿並不保證在Go 1.18釋出前該程式能被成功的運行，但這是目前該語言針對泛型功能提案的主要概念。除此之外，也可以參考Go團隊放出的[型別參數草案](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)，了解更多Go將來可能加入的泛型特性。

泛型可以在許多地方有所使用, 如用於資料結構的實現，如Linked List、Stack、Queue、Set、Map等等。泛型還可用於實現泛型算法，如Sort、Filter、Reduce等等。透過泛型，程式員可以用一種更抽象的方式來處理數據，提高代碼的復用性和健壯性。