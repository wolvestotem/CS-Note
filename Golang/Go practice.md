# Go实践经验积累

[toc]

### 文档资料整理

Go中为什么有代理，有的情况下问什么没有代理：
[GOPRIVATE & GOPROXY](https://jfrog.com/blog/why-goproxy-matters-and-which-to-pick/)

一个不错的go module思想和使用方法说明,包含go.mod，go.sum作用，版本追踪方法..
https://cloud.tencent.com/developer/article/1523542

go结构由vendor到module变化历史和原因，还有形象的呈现
https://www.jianshu.com/p/07ffc5827b26

本地包调用方案
https://zhuanlan.zhihu.com/p/109828249


### range

range slice和map时，index, key, value都是复制出来的值，直接改变不影响slice和map的原始值。如果想改，应该slice[i]，map[key]=value，这种形式修改

```go
for i:=0;i<len(slice);i++{
    //change & update i
}
for key,value := range map{
    map[key] = value2
}
```

### for循环的捕获迭代变量问题

所有for循环，延迟循环元素执行的函数，和迭代变量耦合，出问题：延迟执行的是最后一次迭代变量的值，捕获的是迭代变量的地址

```go
var rmdirs []func()
for _, dir := range tempDirs() {
    os.MkdirAll(dir, 0755)
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir) // NOTE: incorrect!
    })
}
```

核心问题在于储存带迭代变量的函数值，需要将迭代变量赋值为外部变量储存
