Class post 4
=================

OOP 思想---以汉诺塔为例
-----------------

每一个方块是一个`Cube object`，拥有自己的函数和成员；三个可以放方块的场地是三个`stack object`，有自己的成员和函数；完整的过程是一个`Game object`，有自己的成员和`solve`函数，`solve()`中将每个动作抽象成函数。

`API`接口定义：

```C++
#pragma once

//Cube.h
#include "HSLAPixel.h"

namespace uiuc {
  class Cube {
    public:
      Cube(double length, HSLAPixel color);

      double getLength() const;
      void setLength(double length);

      double getVolume() const;
      double getSurfaceArea() const;

    private:
      double length_;
      HSLAPixel color_;
  };
}

//Stack.h
#include <vector>
#include "uiuc/Cube.h"
using uiuc::Cube;

class Stack {
  public:
    void push_back(const Cube & cube);
    Cube removeTop();
    Cube & peekTop();
    unsigned size() const;

    // An overloaded operator<<, allowing us to print the stack via `cout<<`:
    friend std::ostream& operator<<(std::ostream & os, const Stack & stack);

  private:
    std::vector<Cube> cubes_;
};

//Game.h
#include "Stack.h"
#include <vector>

class Game {
  public:
    Game();
    void solve();

    // An overloaded operator<<, allowing us to print the stack via `cout<<`:
    friend std::ostream& operator<<(std::ostream & os, const Game & game);

  private:
    std::vector<Stack> stacks_;

  private:
    void _move(unsigned index1, unsigned index2);
    void _legalMove(unsigned index1, unsigned index2);
};
```


继承
---------------

公有派生类
```C++
class tennisplayer:public player{
    private:
    int _score;
};
```

继承了基类所有成员和方法，但是只能使用public, protected访问private，而不能直接访问。**派生类需要自己的构造函数、额外的成员和方法**

**派生类必然创建基类**，这一点应该体现在构造函数中（初始化列表）。

```C++
tennisplayer::tennisplayer(int i, int j , int k): player(i,j), _score(k) {}

tennisplayer::tennisplayer(player& p1, int k): player(p1), _score(k) {}
```

- 创造基类
- 派生构造函数通过初始化列表传递给基类构造函数
- 派生类构造函数初始化新增数据成员

公有派生是`is-a`关系，派生类是基类的一种。

friend--友元
---------------

遥控器与电视机的关系，一个类可以访问、改变另一个类的`private/protecterd`成员。
