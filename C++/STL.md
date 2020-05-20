# STL库基本用法

## Vector

## String

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

