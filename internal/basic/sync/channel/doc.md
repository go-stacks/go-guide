# 关键点
## 不要通过共享内存来通信，通过通信来共享内存


### 一、ping-pong 模式
ping-pong 模式即乒乓球模式，它比较形象地呈现了数据之间一来一回的关系。收到数据的协程可以在不加锁的情况下对数据进行处理，而不必担心有并发冲突。
具体代码实现请看[pingpong_test.go](pingpong_test.go)

![](https://blob.hixforever.com/20230205182004.png)

