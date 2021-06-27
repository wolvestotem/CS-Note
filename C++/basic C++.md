## 位运算

```C++
//取数
for(int i=31;i>=0;i--){
    int u = (num>>1)&1;
}
//设置数
int res=0;
for(int i=31;i>=0;i--){
    res = (res<<1)|1;//1
    res = (res<<1);//0
}
#include<bitset>
bitset<31> bs(num);
```