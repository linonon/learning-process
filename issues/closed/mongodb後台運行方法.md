# mongodb设置后台运行的方法 

## brew

```shell
# alias start-mongo="brew services restart mongodb-community@5.0"
brew services restart mongodb-community@5.0
```

## shell cmd(deprecated)

默认的情况下，关闭shell，mongodb就停止运行了。

如果想在后台运行，启动时只需添加 --fork函数即可。

可以在日志路径后面添加--logappend，防止日志被删除。

代码如下:

```bash
bin/mongodb  --fork --dbpath=//  --logpath=//  --logappend
```

### shutdownServer()

在后台运行，如果想要关闭它的话，需要给他发送shutdownServer()

#### 普通命令：

要注意的是，这个命令只允许在本地，或是一个经过认证的客户端。

代码如下:

```bash
$ ./mongod
> use admin
> db.shutdownServer()
```

#### 主從式的複製集群

如果这是一个主从式的复制集群，在1.9.1版本后将按下面的步骤来关闭

1. 检查从Mongodb的数据更新时间
2. 如果所有的从Mongodb和主的时间差都超过10，这个时候不会关闭mongodb（在这种情况下面，我们可以通过配置timeoutSecs的方式来让从Mongodb完成数据的更新）
3. 如果其中有一个从Mongodb与主服务时间差在10秒内，那么主服务器将会关闭，并且等待从Mongodb更新完成并关闭。

#### 強制關閉服務

如果没有up-to-date 从Mongodb且你想强制关闭服务，可以通过添加force:true;命令如下：
代码如下:

```bash
$ ./mongod
> db.adminCommand({shutdown : 1, force : true})
> //or
> db.shutdownServer({force : true})
```

#### 超時關閉

指定特定超时时间的关闭服务器，命令同上，另外加上一个timeoutsec:参数
代码如下:

```bash
$ ./mongod
> db.adminCommand(shutdown : 1, force : true, timeoutsec : 5)
> //or
> db.shutdownServer({force : true, timeoutsec : 5})
```

原文地址：<https://m.xp.cn/b.php/61959.html>
