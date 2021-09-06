# Go实践经验积累

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
