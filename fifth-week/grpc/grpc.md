# grpc
## 组成
1. 客户端(client): 服务端的调用方
2. 服务端(Server): 真正的服务端提供方
3. 客户端存根(ClientOption): socket管理, 网络收发包的序列化
4. 服务端存根(ServerOption): socket管理, 提醒server层rpc方法调用, 以及网络是收发包的序列化  

逻辑图让如下:
![](./rpc%E9%80%BB%E8%BE%91%E5%9B%BE.png)
## 什么是grpc
grpc是rpc的一种, 它使用的是protocol buffer作为序列化格式, 

```
    http/1.1和http/2的区别
    1. http/2采用的是二进制传输的而不是文本
    2. http/2 采用的是完全多路复用, 只需要一个连接就可以实现并行操作
    3. 使用报头压缩, 降低开销
    4. http/2可以让服务器将响应主动推送到客户端缓存中
    
```