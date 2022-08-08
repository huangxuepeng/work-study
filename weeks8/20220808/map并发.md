### 问题出现场景
![](./photo/%E5%87%BA%E7%8E%B0%E5%9C%BA%E6%99%AF.jpg)
1. 为什么不能使用并发处理数据?
    - 传入的apps, dataCells, roles等数据都是map类型, map类型本身就是并发不安全.

### 解决方案
```
type Map struct {
	mu Mutex

	read atomic.Value 

    /*
        其中包含新写入的key和read中所有未删除的key
    */
	dirty map[interface{}]*entry

	misses int
}
```