# StarShip 配置

## Q: starship

教我有關 starship 的配置相關的知識, starship 是一個 zsh prompt

## A: starship

Starship 是一個高度可定製的命令提示符工具，它是用 Rust 編寫，支援多種 shell，包括 Zsh。以下是有關 Starship 配置的相關知識：

1. 安裝 Starship：首先，你需要在你的系統上安裝 Starship 工具。你可以參考 Starship 的官方文件，根據你的作業系統進行安裝。

2. 設定 Starship：在 Zsh 中配置 Starship 非常簡單。你需要在你的 Zsh 配置文件中（通常是 `~/.zshrc`）加入類似下面的行：

```shell
eval "$(starship init zsh)"
```

這行會引入並初始化 Starship。

3. 客製化主題：Starship 提供了許多預設的主題可以選擇，你可以在 `~/.config/starship.toml` 中配置你喜歡的主題。你可以在 Starship 的 GitHub 存儲庫中找到更多的主題選項。

4. 指示文檔：Starship 支援許多指示符用於定製你的提示符。你可以在 `~/.config/starship.toml` 配置文件中添加或修改這些指示符，以滿足你的需求。一些常見的指示符包括：

- `username`: 顯示用戶名。
- `hostname`: 顯示主機名。
- `directory`: 顯示當前目錄的路徑。
- `git_branch`: 顯示 Git 庫的當前分支。
- `git_status`: 顯示 Git 庫的狀態，例如是否有未提交的更改。
- `package`: 顯示當前所使用的包管理器，如 `npm` 或 `cargo`。

你可以根據自己的需要在 `starship.toml` 文件中，使用格式如下：

```toml
[character]
success_symbol = "✅"
error_symbol = "❌"

[git_branch]
symbol = " branch: "
```

在這個例子中，我們更改了成功和錯誤符號的樣式，以及 Git 分支符號。

5. 重新載入配置：當你更改了 `starship.toml` 文件後，你需要重新載入你的終端機窗口或運行 `source ~/.zshrc` 來讓更改生效。

以上就是有關 Starship 配置的一些相關知識。你可以根據需要進一步自定義 Starship 的提示符。祝你使用愉快！

## Q: directory

告訴我 starship.toml 中 [directory] 的配置

## A: directory 

在 Starship 的 `starship.toml` 配置文件中，你可以使用 `[directory]` 部分來配置目錄指示器的外觀和行為。以下是一些常見的配置選項：

1. `truncation_length`：設置目錄路徑的截斷長度。例如，`truncation_length = 2` 將 `/home/user/Documents` 縮寫為 `~/Documents`。預設值為 3。

2. `truncate_to_repo`：設置是否只在 Git 庫中顯示被截斷的目錄。如果設置為 `true`，則目錄 `~/myproject/src` 將顯示為 `src`（只在 Git 庫中），而不是 `~/myproject/src`。預設值為 `false`。

3. `prefix` 和 `suffix`：設置目錄指示器的前綴和後綴。你可以使用字符、Unicode 字符或 ASCII 碼（如 `prefix = "["` 或 `prefix = "\u{1F512}"`）。

4. `style`：設置目錄指示器的樣式，可以是對象包括 `fg`（前景色）、`bg`（背景色）、`bold`（粗體）、`underline`（底線）、`italic`（斜體）等。例如：

```toml
[directory.style]
fg = "green"
bg = "black"
bold = true
```

這將使目錄指示器的文本顯示為綠色的粗體文本，黑色背景。

以上是 `[directory]` 配置的一些常見選項。你可以根據需要配置其他選項或嘗試不同的樣式來個性化你的目錄指示器。另外，如果你想要了解更多關於 Starship 配置的詳細信息，可以參考官方文檔。

## Q: git_branch & git_status

告訴我 starship.toml 中 [git_branch] 和 [git_status] 的配置


## A: git_branch & git_status

在 Starship 的 `starship.toml` 配置文件中，你可以使用 `[git_branch]` 部分來配置 Git 分支指示器的外觀和行為，而 `[git_status]` 部分則用於配置 Git 狀態指示器的外觀和行為。以下是一些常用的配置選項：

## 配置 Git 分支指示器（[git_branch]）：

1. `symbol`：設置分支指示器的符號。例如，`symbol = " "` 將使用 " " 作為分支指示器。你可以使用字母、符號、Unicode 字符或 ASCII 碼。

2. `truncation_length`：設置分支名稱的截斷長度。例如，`truncation_length = 10` 將把 `my-feature-branch` 縮寫為 `my-feat...`。預設值為 2。

3. `symbol_style` 和 `branch_name_style`：分別設置分支指示器符號和分支名稱的樣式，可以配置顏色、粗體、斜體、底線等。例如：

```toml
[git_branch.style]
fg = "green"
bg = "black"
bold = true
```

這將使分支指示器符號顯示為綠色的粗體文本，黑色背景。

## 配置 Git 狀態指示器（[git_status]）：

1. `disabled`：設置是否禁用 Git 狀態指示器。可以將其設置為 `true` 或 `false`。預設值為 `false`（啟用）。

2. `format`：設置狀態指示器的格式。你可以自定義顯示的信息，例如 `"[{\$all}]({$symbols} ST=[$status] in {$name})"`。該格式字符串可以使用變數如 `$all`（所有文件），`$symbols`（狀態符號），`$status`（狀態描述符），`$name`（Git 庫名稱）等。

3. `style`：設置狀態指示器的樣式，可以設置顏色、粗體、斜體、底線等。例如：

```toml
[git_status.style]
fg = "red"
bg = "black"
bold = true
```

這將使狀態指示器的文本顏色為紅色的粗體，黑色背景。

這些是 `[git_branch]` 和 `[git_status]` 配置的一些常見選項。你可以根據需要進一步設置其他選項或嘗試不同的樣式來個性化你的 Git 分支和狀態指示器。如有需要，你可以參考 Starship 官方文檔以獲取更多詳細信息。