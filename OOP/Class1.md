OOP class post 1
===========

private:

Data function only used by class, connot used by client code

public:

Data function can be used by client code

interface(.h) seperated from the implementation(.cpp)

.h--declarition API

.cpp--Implementation

*namespace*

specify our **Cube** class

```C++
namespace uiuc {
  class Cube {
    public:
      double getVolume();
      double getSurfaceArea();
      void setLength(double length);

    private:
      double length_;
  };
}

namespace uiuc {
  double Cube::getVolume() {
    return length_ * length_ * length_;
  }

  double Cube::getSurfaceArea() {
    return 6 * length_ * length_;
  }

  void Cube::setLength(double length) {
    length_ = length;
  }
}
```
