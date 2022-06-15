# 問題地址

<https://github.com/fabric/fabric/issues/2007>

## 正確例子

```python
c = Connection(
    host="hostname",
    user="admin",
    connect_kwargs={
        "key_filename": ["/home/myuser/.ssh/private.key"],
    },
)
```

`"key_filename"`後加的應該是`陣列`
