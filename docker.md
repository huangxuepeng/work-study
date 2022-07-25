### 为什么拉取一个镜像的时候会有多条数据同时下载
1. 一个镜像是由多个只读层组成的
    - 这些只读层除了最下面的不指向外, 其他的都指向父级
    - 展示给用户的只有一个文件系统的原因是隐藏多层的存在, 使用的技术是使用统一文件系统技术, 将多层文件整合成一个文件系统.
#### 分层的好处
1. 共享资源: 有多个镜像都是从` base `镜像构建而来的， 那么只需要在宿主机的磁盘上和内存中各保存一份镜像, 每一层都可以被共享.
    - base镜像
### 当删除一个正在运行或者已经停止的容器的镜像报错的原因
报错信息
> Error response from daemon: conflict: unable to remove repository reference "hxp-test:v1" (must force) - container 6df4e796bf15 is using its referenced image 0964cd5e00ce

#### 原因剖析: 
##### 容器组成: 
- 不在运行中的容器是由 镜像 + 读写层组成
- 正在运行的容器是由   镜像 + 读写层 + 隔离的进程空间 + 其中运行的进程
- 对容器的修改, 创建, 删除 只是作用于读写层
#### 解决方案: 
1. 查询依赖此容器的镜像
- 容器正在运行, 应当先停止镜像的运行, 而后进入2
- 容器已经停止 进入2
2. 将停止的容器删除
3. 删除镜像

### Dockerfile在书写的时候可能会有多个run
```
1. 执行多个run (ex1)
    FROM busybox
    RUN echo This is the A > a
    RUN echo This is the B > b
    RUN echo This is the C > c
2. 执行一个run (ex2)
    FROM busybox
    RUN echo This is the A > a &&\
        echo This is the B > b &&\
        echo This is the C > c
```
1. 当run删除由前一个run(yum install nano && yum clean all) 添加的东西时, 第二种方法显然是正确的.
2. 层应该只是在前一层之上添加一个差异, 所以后一层如果对前一层没有删除的操作的时候, 两种方法之间没有太多的磁盘空间节省;
3. 镜像下载的时候, 镜像的每层是并行下载的, 所以理论上ex1打包出来的镜像下载的时候理论会快一点
4. 如果添加第四句的时候(echo this is the D > d) 并且在本地重建, 由于缓存, ex1会构建的更快, ex2会再次运行这四个命令
