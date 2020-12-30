# STL库基本用法

## Vector

vector方法

```C++
    v1.resize(distance(v1.begin(),unique(v1.begin(), v1.end())));
    // 10 20 20 20 30 30 20 20 10---->10 20 30 20 10
    //sorted -> unique vector
    if(find(v.begin(),v.end(),20) != v.end())
        cout<<"found";
```
 
## String

初始化

```C++
    string one("hello");
    string two(2,'$');
    string three(one);
    two = "hello world";
    one +="hhh";
    three[0] = '!';
    string four(one.begin(),one.end()-1);
```

string 方法

```C++
    string str("Please, replace the vowels in this sentence by asterisks.");
    int found =str.find("in");
    if(found!=string::npos)
        str[found]='*';
    int found = str.find_first_of("aeiou");
    while (found != string::npos)
    {
        str[found] = '*';
        found = str.find_first_of("aeiou", found+1);//find(string str, int start_rank)
    }
    std::cout << str << '\n';
    //OUTPUT: Pl**s*, r*pl*c* th* v*w*ls *n th*s s*nt*nc* by *st*r*sks.

    cout << "Splitting: " << str << '\n';
    int found = str.find_last_of("/\\");
    cout << " path: " << str.substr(0,found) << '\n';//str.substr(int start_rank, int length)!!
    cout << " file: " << str.substr(found+1) << '\n';

    //OUTPUT: Splitting: /usr/bin/man
                    //path: /usr/bin
                    //file: man

    str="sample phrase!!";
    str.replace(str.begin(),str.begin()+6,"replace");//左开右闭
    str.replace(9,5,"!!!");//posrank, len, str.短的只替换3个
```

### string与其他类型转换

```C++
string to_string (int val);
string to_string (long val);
string to_string (long long val);
string to_string (unsigned val);
string to_string (unsigned long val);
string to_string (unsigned long long val);
string to_string (float val);
string to_string (double val);
string to_string (long double val);

//
int i_dec = std::stoi (str_dec,&sz);
int i_hex = std::stoi (str_hex,nullptr,16);
int i_bin = std::stoi (str_bin,nullptr,2);
int i_auto = std::stoi (str_auto,nullptr,0);
```

### stringstream

```C++
#include<sstream>
string text = "To be or not to be";
stringstream ss(text);//空格分隔
string word;
while(ss>>word)
    cout<<word;
//Tobeornottobe
string s = ss.str();

stringstream sss;
sss<<100<<' '<<200;//空格分隔
int a,b;
sss >> a >> b;
//a=100, b=200
```

## Unordered_map

```C++
    unordered_map<int,string> m;
    //initailize
    m.insert(make_pair(1,"see"));
    m.insert(make_pair(2,"like"));
    m.emplace(3,"who");
    m.emplace(4,"are");
    m[5]="you";

    //traversal
    for (auto i : m)
        cout << i.first << "    " << i.second //element
             << endl;
    for(auto i=m.begin();i!=m.end();i++)
        cout << i->first << "      " << i->second //iterator
             << endl;
```

### map

```C++
    map<int,string> m;
    //initialize
    m[5]="hhh";
    m.insert(make_pair(3,"jij"));
    m.emplace(2,"who");
    //m是有序的
    //m.begin()->fir最小
```

### multimap

```C++
    multimap<int,string> m;
    //没有assignment
    m.insert(make_pair(3,"ggg"));
    m.emplace(2,"who");
    //multimap也是有序的
```

## Queue

### Priority_queue
发现STL默认都是使用()比较的，默认比较使用less（即'<'运算符）,priority_queue<>默认是大根堆的，这是因为优先队列队首指向最后，队尾指向最前面的缘故！每次入队元素进去经排序调整后，优先级最大的元素排在最前面，也就是队尾指向的位置，这时候队首指向优先级最小的元素！

优先队列的这个类型，其实有三个参数：priority_queue<class Type,class Container,class Compare>，即类型，容器和比较器，后两个参数可以缺省，这样默认的容器就是vector，比较方法是less，也就是默认大根堆，可以自定义写比较方法，但此时若有比较方法参数，则容器参数不可省略！priority_queue<>的可支持的容器必须是用数组实现的容器，如vector，deque，但不能是list（推荐vector），比较方法可以写结构体重载()运算符，也可以用less，greater这些语言实现了的，但是灵活性不够，建议手写重载结构体

这个类里重载operator()，**和自定义sort排序不同**，sort只需bool cmp（……）（当然sort也可以弄一个比较类，再重载operator()）,若是priority_queue中写为sort的cmp形式则报错，如：bool cmp1(const node &a, const node&b)//报错！

```C++
struct compare{
    bool operator()(ListNode* a, ListNode* b){
        return a->val>b->val;
    }
};

void func(){
    priority_queue<ListNode*,vector<ListNode*>,compare> q;
    priority_queue<int,vector<int>,greater<int> > p;//小顶堆--堆顶在尾
    //方法
    while(!p.empty()){
        p.push(1);
        p.pop();
        int x = p.top();
        p.emplace(2);
        int size = p.size();
    }
}
```

### heap

```C++
    //默认是maxheap
    auto compare = [](int a, int b){return a>b;};//minheap

    int myints[] = {10,20,30,5,15};
    std::vector<int> v(myints,myints+5);

    make_heap (v.begin(),v.end());
    cout << "initial max heap   : " << v.front() << '\n';
    //30

    pop_heap (v.begin(),v.end()); v.pop_back();
    cout << "max heap after pop : " << v.front() << '\n';
    //20

    v.push_back(99); push_heap (v.begin(),v.end());
    cout << "max heap after push: " << v.front() << '\n';
    //99

    sort_heap (v.begin(),v.end());
    //破坏堆结构，不再是堆

    cout << "final sorted range :";
    for (unsigned i=0; i<v.size(); i++)
    cout << ' ' << v[i];
    //5 10 15 20 99
```

## Stack

## unordered_set

```C++
    unordered_set<int> lookup;
    lookup.insert(4);
    lookup.insert(1);
    lookup.insert(1);
    lookup.insert(2);
    lookup.emplace(5);
    for (auto i:lookup)
        cout<<i<<' ';
    //OUTPUT: 4 1 2
```

## list

```C++
    list<int>::iterator it;
    list<int> mylist(5,1);
    vector<int> myvector(4,10);

    it = mylist.begin();
    it++;
    mylist.insert(it, 10);//1 10 1 1 1 1
    mylist.insert(it, 2,3);//1 10 3 3 1 1 1 1
    it--;
    mylist.insert(it, myvector.begin(),myvector.end());
    //1 10 3 10 10 10 10 3 1 1 1 1
    mylist.push_front(100);
    mylist.push_back(200);
    mylist.pop_back();
    mylist.pop_front();

    //查看
    mylist.front();
    mylist.back();
    //size;
    mylist.size();
    mylist.empty();
    //删除
    mylist.erase(mylist.begin());
    mylist.erase(mylist.begin(),mylist.end());
```

## pair

```C++
    pair<int,int> p(0,0);
    pair<int,int> p1;
    p1 = make_pair(0,0);
```
