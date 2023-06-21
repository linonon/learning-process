
# 使用 shell, 批量自動更新 git


## Q: 如何寫這個 shell

macOS: 我現在在一個目錄 ~/a 下, 這個目錄下還有多個子目錄 ~/a/b, ~/a/c, ~/a/d, 以及一個隱藏的目錄 ~/a/.e. 其中 b,c,d,e 中都有一個 .git. 現在我希望你寫一個 shell, 每次執行shell時, 在 ~/a 目錄下的子文件夾裡都會執行`git pull`, `git add .`, `git commit -m {date}`, `git push`

## A: 如何寫這個 shell

```sh
#!/bin/bash

# 獲取當前日期
date=$(date '+%Y-%m-%d %H:%M:%S')

# 設置基礎目錄
base_dir=~/a

# 遍歷所有的子目錄
for dir in "$base_dir"/* "$base_dir"/.*/
do
    # 檢查子目錄中是否存在 .git
    if [ -d "$dir/.git" ]; then
        echo "處理目錄 $dir"
        cd "$dir" || exit
        git pull
        git add .
        git commit -m "$date"
        git push
    fi
done

```