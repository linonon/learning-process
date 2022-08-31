# Error handling

## 官方文檔
- fmt.Errorf(str string, f)

## 錯誤處理方法

### 用 defer + recover 處理
- 錯誤處理的方法
- recover() 不會影響代碼的運行

### 自定義錯誤處理
- errors.New()  返回錯誤類型
- panic() 內置函數， 可以接受interface類型，然後輸出並終止程序。