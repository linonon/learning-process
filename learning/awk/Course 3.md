---
topic: AWK 處理文本和字段
home page: "[[Learning AWK]]"
---
## 第3節：處理文本和字段

在這一節課中，我們將深入學習如何在 AWK 中處理文本和字段，特別是如何利用 AWK 強大的字段分割能力，打印特定字段，以及如何使用一些關鍵的內置變量來定制我們的文本處理任務。

### 字段分割

AWK 會根據字段分隔符（`FS`）自動將每一行分割成多個字段。默認的字段分隔符是空白字符，包括空格和制表符。這意味著，如果一行文本包含由空格或制表符分隔的多個單詞，AWK 會將每個單詞視為單獨的字段。

- **設置字段分隔符：** 我們可以通過設置 `FS` 變量在 AWK 程序中改變字段分隔符。例如，要處理逗號分隔值（CSV）格式的文件，我們可以將 `FS` 設為逗號：

```awk
BEGIN { FS = "," }
```

### 打印特定字段

一旦文本按字段分割，我們就可以選擇打印某一行的一個或多個特定字段。字段由 `$` 符號和一個數字標識，其中 `$1` 表示第一個字段，`$2` 表示第二個字段，依此類推。`$0` 表示整行文本。

- **例子：** 假設我們有一個文本文件 `people.txt`，包含以下內容：

```
Alice 30 New York
Bob 25 Los Angeles
Charlie 35 Chicago
```

要打印每一行的名字和城市，我們可以使用以下命令：

```sh
awk '{ print $1, $3 }' people.txt
```

這會輸出：

```
Alice New York
Bob Los Angeles
Charlie Chicago
```

#### 使用內置變量（如 FS, OFS）

- **`FS`（Field Separator）：** 如上所述，用於定義字段的分隔符。
- **`OFS`（Output Field Separator）：** 定義輸出中字段之間的分隔符。默認情況下，`OFS` 是空格。如果我們想在輸出中使用不同的分隔符（比如逗號），我們可以設置 `OFS`：

```awk
BEGIN { OFS = "," }
```

- **應用示例：** 修改上面的例子，讓輸出的名字和城市之間用逗號分隔：

```awk
BEGIN { OFS = "," } { print $1, $3 }
```

這樣，如果我們再次執行 AWK 腳本處理 `people.txt` 文件，輸出將會是：

```
Alice,New York
Bob,Los Angeles
Charlie,Chicago
```

通過這些基本的操作和變量設置，AWK 使得文本和字段的處理變得非常靈活和強大。瞭解如何有效地分割和操作字段是使用 AWK 進行文本處理的關鍵。

當然可以。AWK 提供了一系列的內置變量，這些變量在文本處理中非常有用。除了前面提到的 `FS`（字段分隔符）和 `OFS`（輸出字段分隔符）之外，還有其他一些重要的內置變量：

### `NR`：Number of Records

`NR` 保存了到目前為止讀取的輸入記錄（行）的數量。這在處理多行輸入時非常有用，比如，打印文件的當前行號：

```awk
awk '{print NR, $0}' filename
```

這會在每行的開頭加上該行的行號和原始內容。

### `NF`：Number of Fields in the Current Record
`NF` 表示當前記錄中的字段數量。這在需要處理具有不同字段數的行時特別有用，比如，打印每行的最後一個字段：

```awk
awk '{print $NF}' filename
```

### `FNR`：File Number of Records
`FNR` 是對於當前輸入文件的紀錄數（與 `NR` 類似，但 `FNR` 會在每個新文件開始時重置為1）。當同時處理多個文件時，`FNR` 對於區分每個文件中的記錄位置很有用。

```awk
awk '{print FNR, $0}' file1 file2
```

### `RS`：Record Separator
`RS` 定義了一條記錄的結束符，默認是換行符。這在讀取多行記錄的文件時非常有用，比如，設置 `RS` 為空白行來處理段落或記錄塊。

### `ORS`：Output Record Separator
`ORS` 用於定義輸出中記錄之間的分隔符，默認值同樣是換行符。這對於改變輸出格式非常有用。

```awk
BEGIN { ORS = "\n\n" } { print }
```

### `ARGV`、`ARGC`：Argument Vector and Argument Count
這些變量分別用於存儲命令行參數和參數的數量，使得 AWK 腳本可以處理傳遞給它的參數。

### `ENVIRON`：Environment Variables
`ENVIRON` 是一個關聯數組，包含了當前環境的環境變量。可以使用 `ENVIRON["HOME"]` 這樣的語法來訪問特定的環境變量。

這些內置變量擴展了 AWK 的功能，使其能夠更加靈活和強大地處理各種文本處理任務。通过結合使用這些變量，可以實現從簡單的文本處理到複雜的數據分析的各種操作。