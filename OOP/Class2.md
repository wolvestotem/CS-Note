Class post 2
======

Four things of every C++ variable:
+ Name
+ Type
+ Value
+ location of memory

stack memory

自顶向下的 stack，先入栈的 memory 大

return of function will lead to reuse of stack memory. 所以返回一个变量地址给指针是无效的。它指向的是一个被释放的stack地址

heap memory

only way to use heap memory is to use "new". 

```C++
int *p = new int;
```

A stack pointer points to heap memory.

只能通过delete释放pointer所指向的heap memory，并且只能释放一次

```C++
int *x = new int;
int &y = *x;
y = 4;
```

reference y: give name to a piece of memory

