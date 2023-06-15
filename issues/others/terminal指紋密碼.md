# Terminal 指紋密碼

[link](https://zhuanlan.zhihu.com/p/31529925)

## 安裝

```shell
sudo sed -i ".bak" '2s/^/auth       sufficient     pam_tid.so\'$'\n/g' /etc/pam.d/sudo
```

## 還原

```shell
sudo mv /etc/pam.d/sudo.bak /etc/pam.d/sudo
```

## iTerm 不生效的話

iTerm 需要进入设置，Advanced ，搜索 sudo ，把选项值改为 no
