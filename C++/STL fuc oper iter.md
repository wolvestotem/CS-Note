# STL跨越container泛型用法

### emplace

以vector为例，设A是一个struct结构

```C++
vector<A> seq;
A a(10);
seq.push_back(a);//copy constructor
seq.push_back(A(5));//default constructor: a tempotory object in stack, copy constructor to container
seq.emplace_back(A(3));//default constructor in container

seq.emplace(seq.begin()+1, 3);
seq.insert(seq.begin(), A(11));
```

适用STL：

```C++
vector.emplace_back();
list.emplace_back();

vector.emplace(iterator, arg);
map.emplace(key,value);
unordered_map.emplace(key,value);
list.emplace();
```
