[toc]

## Distributed System

### Introduction

Fault tolerence

Performance

Concurrency

##### Infra

Storage, Communication, Computation

##### Impl

RPC, thread, Concurrency

##### Performance

Scalability

##### Fault Tolerance

Availability, Recoverability --- NonVault Storage, Replication

##### Consistency

####  Map Redunce

Input1 File -> Map output pairs of kv

Collect all the keys from all the maps

Reduce function

Ex: word count

Map(k,v): split text into words, for each words, emit(w,"1")

Reduce(k,v): emit(len(v))

GFS: google file system

bootle net is net 50M/s/machine, so map same machine as the GFS server, reduce must net tranfer, expensive part



### RPC & Thread

thread importance:

I/O concurrency, parallism, convenience(异步线程, ex 定时器)

另一种风格: event-driven programming 一个线程 监听事件

 