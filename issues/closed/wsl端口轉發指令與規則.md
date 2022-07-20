# Wsl 端口轉發指令與規則

網站: https://zhuanlan.zhihu.com/p/355748937

- 端口轉發到 wsl
  netsh interface portproxy add v4tov4 listenaddress=0.0.0.0 listenport=2222 connectaddress=[IP] connectport=[PORT]

- 防火墻入站規則

  netsh advfirewall firewall add rule name=WSL2 dir=in action=allow protocol=TCP localport=2222
a

