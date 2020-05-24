# STL库基本用法

## Vector

## String

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

## Queue

## Stack

## unordered_set

```C++
    unordered_set<int> lookup;
    lookup.insert(4);
    lookup.insert(1);
    lookup.insert(1);
    lookup.insert(2);
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
