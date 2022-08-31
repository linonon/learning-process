# Testing & go test 入門
- 單元測試 & 性能測試
- 可單獨寫測試用例
- 確保高性能（高併發中保持穩定），可用性（結果正確）。

## Testing 的好處
- 傳統方法（在main函數中執行）很麻煩，如果項目正在運行，則要停止項目再測試
- 傳統方法不利於管理，因為你多個函數模塊都放在main中的時候，測試結果通常不夠直觀
- 印出單元測試 ，go 中 -> testing 測試框架，可以很好解決上面的問題。
- 分離測試數據
- 明確錯誤信息
- 可以部分失敗

## Testing 框架原理
1. 將 xxx_test.go的文件引入import
2. main{調用 func TextXxx()}
3. 只要 Xxx 的首字母不是[a-z]，則所有TestXxx函數都會被調用 

### 入門總結
`func TestXxx (t *testing.T)`
1. 文件名必須以 ”_test.go“結尾
2. 測試例子函數必須用Test開頭，一般來說是Test + 測試函數名
3. 一個測試文件可以有多個測試函數
4. go test -v： 無論結果是否正確都會輸出日誌
5. t.Fatalf 格式化輸出錯誤信息，並退出函數
6. 不需將函數放在main()中
7. PASS 通過， FAIL 失敗。
8. 測試單文件，必須帶上原文件
9. 測試單獨一個method go test -v -test.run TestAddUpper

## coverprofile
`go test -coverprofile=c.out`
`go tool cover -html=c.out` 代碼覆蓋率

## benchmark
- `func BenchmarkXxx(b *testing.B)`
- `go test -bench .` 
- `go test -bench . -cpuprofile cpu.out` 生成一個cpuprofile性能數據文檔, 名字叫 cpu.out
- `go tool pprof cpu.out` 使用pprof 查看cpu.out文檔
- `(pprof) web` 使用Graphviz 查看cpu.out 文檔, 需要先用 `brew install svn` 後再 `brew install graphviz`
- 如果是使用vscode打開的.svg格式的文件,記得在mac中更改對.svg文件的打開方式(chrome)
- 
- **分析性能, 優化代碼.**