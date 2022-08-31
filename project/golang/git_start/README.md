# Git 使用測試

## Global setting
- `git config --global user.name "Oliver"`
- `git config --global user.email "oliver@labsgl.com"`

## Create a new repository
- `git clone git@35.234.5.205:oliver/<filename>.git`
- `cd learngo`
- `touch README.md // 創建 readme`
- `git add README.md//添加本地文件readme.md到git當前目錄`

## [Git 30 Days](https://github.com/doggy8088/Learn-Git-in-30-days)

### Git init
- 建立Git存儲庫
- ` --bare` 共享存儲庫,多人操作.

### 04
- `git commit`                  提交
- `git add <url>`               新增
- `git log -n`                  查詢紀錄 -n: 紀錄條數
- `git rm 'filename'`           刪除
- `git mv filename1 filename2`
- `git checkout <branchname> filename` 還原 <<branchname>>分支中的filename文件

### 05 
- 物件&索引的關係

### 06
Git 中 物件結構好處
1.  有效處理大專案
2.  歷史版本保護
3.  維護封裝物件

### 07
索引四種狀態
1.  untracked
2.  unmodified
3.  modified
4.  staged

### 08
分支基本觀念
- 讓團隊有版本管控觀念的習慣

分支不能自己delete自己

`git branch <branchname>` 創建分支
`git checkout -b <branchname>` 創建分支並切換工作區域

### 09
`git diff <id1> <id2>`

### 10
物件 --'SHA1'-> 物件的絕對名稱

### 11
物件的參照名稱

### 12
參照名稱的實現原理

### 13
`git stash` 暫存,簡單的`分支`與`合併`

### 14 
`git config`

### 15
`git tag [tagname]` 版本命名

### 16
`git relog` 查看歷史log

### 17
- 合併、衝突合併
- 解決衝突: git環境自帶?

### 18
`git reset` : 
  1.  `--hard "HEAD^"` : HEAD = HEAD^ HEAD指向了前一個版本
  2.  `--soft` : 刪除版本紀錄但保持異動內容
  3.  `--hard ORIG_HEAD` : 

`git commit --amend` 修改

### 19
`.gitignore` 忽略清單

### 20
`git revert` -> `git revert --continue` or `git revert --abort`

### 21
`git cherry-pick` 手動挑出需要的變更.
  - `-e` 建立版本前, 先編輯信息
  - `-n` 暫時先不建立版本, 留下Author & Data資訊

### 22
`git rebase` 將現有分支進行 重新指定基礎版本 的操作

### 23
如果分支使用遠端數據庫Download回來的, 不應該使用Rebase修改歷史版本紀錄, 因為這樣會導致無法commit

`git rebase -i [commit_id]`

### 24
`git clone`

### 25
遠端數據庫概念

### 26
(`git fetch` -> `git merge xx/xx`) or `git pull` -> `git push`

### 27
`git push --all --tags`

### 28
github管理:
- pull only
- push & pull
- push, pull & administrative (以及專案管理權限)

設定群組

fork: 把別人的專案 變成自己管理的專案

### 29
Subversion(SVN)

### 30
小技巧


## Git branch 資料
[Git Branch 學習](https://backlog.com/git-tutorial/tw/stepup/stepup1_1.html)

## Git fetch
`git pull` = `git fetch` + `git merge`

## 問題解決: 避免重複輸入ssh-key/解決vscode中無法輸入ssh-key導致認證出錯的問題

1. 使用 `ssh-agent` 命令啟動 ssh-agent
2. 使用 `ssh-add` 命令自動添加 ssh-key
3. 完成，以後進行 `git push/pull` 等操作則不需要再次輸入passphrase


