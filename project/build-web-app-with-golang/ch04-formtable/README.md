# 4 Form Table

## 4.1 Process form table input
- 通過 `GET` & `POST` 獲取 `Data`

## 4.2 Verification of input
- Required Fiedls
- Numbers
    - (opt) using regular expressions.
- Chinese / English
    - (Only) using regular expressions.
- E-mail, Drop down list, Radio button(單選)
- Check-boxes, Data & time, Identify

## 4.3 Cross site scripting
- Validation of all data form users
- Carefully handle data that will be sent to clients in order to prevent any injected scripts from running on browsers.

## 4.4 Duplicate submissions
- Add a hidden field with a unique token to `Form`