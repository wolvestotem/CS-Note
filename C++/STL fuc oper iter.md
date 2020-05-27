# STL跨越container泛型用法

### 常用algorithm

```C++
    reverse(v.begin(), v.end());
    sort(v.begin(), v.end(),myfunc);
    for (int i = 0; i < v.size(); i++)
        cout << v[i] << '\t';
    swap(v[1],v[2]);
    vector<int>::iterator low,hi;
    low = lower_bound(v.begin(),v.end(),20);
    hi = upper_bound(v.begin(),v.end(),20);
    cout<<distance(v.begin(),low)<<distance(v.begin(),hi)<<endl;
    // 10 10 10 20 20 20 30 30--->low->3  hi->6
```

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

### lambda

函数指针、函数符、lambda比较

函数指针

```C++
bool f3(int x){return x%3 == 0;}

vector<int> vector(1000);
generate(vector.begin(),vector.end(),rand);
int count3 = count_if(vector.begin(),vector.end(),f3);
```

lambda函数：匿名函数，使用`[]`代替了函数名；没有返回类型，靠自动推断

```C++
int count3 = count_if(vector.begin(),vector.end(),[](int x){return x%3 == 0;});
```

添加按照访问和按引用访问变量

```C++
int count13 = 0;
for_each(vector.begin(),vector.end(),[&count13](int x){count13 += x%13==0});//void 
```
