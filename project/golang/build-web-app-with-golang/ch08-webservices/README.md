# WebServices

## 8.3 REST

REpresentational State Transfer

- Resources

通過 URI 定位的資源（圖片，文檔，影片等）。 URI == resources

- Representation

如：`TXT` 為資源，那表現就是 “把 `TXT` 通過不同方式（`JSON`，`HTML`，`XML`）表現出來）

表現形式通過設置 `Header` 的 `Accept` && `Content-Type` 決定

- State Transfer

`HTTP` 本身是無狀態的 (stateless protocol), 那麼狀態就要裝在 `Server` 中

總結就是：

1. Every URI represents a  resources.
2. There is a representation layer for transferring resources between clients and servers.
3. Clients use four HTTP methods to implement "Presentation Layer State Transfer", allowing them to operate on remote resources.

![](https://github.com/astaxie/build-web-application-with-golang/blob/master/en/images/8.3.rest2.png?raw=true)

Figure 8.5 REST architecture

![](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/images/8.3.rest.png?raw=true)

Figure 8.6 REST's expansibility.
  
## 8.4 RPC

RPC: remote precedure call

### RPC working principle
![](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/images/8.4.rpc.png?raw=true)

Figure 8.8 RPC working principle

An RPC call from client to server has the following ten steps:

1. Call the client handle, execute transfer arguments.
2. Call local system kernel to send network messages.
3. Send messages to remote hosts.
4. The server receives handle and arguments.
5. Execute remote processes.
6. Return execution result to corresponding handle.
7. The server handle calls remote system kernel.
8. Messages sent back to local system kernel.
9. The client handle receives messages from system kernel.
10. The clinet get results from corresponding handle.

### Go RPC

Correct example
```go
func (t *T) MethodName(argType T1, replyType *T2) error
```
- Where T, T1 and T2 must be able to be encoded by the `package/gob` package.
