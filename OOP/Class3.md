Class post 3
===========

Constructor:

copy constructor:
----------------

- A class constructor
- Has exactly one argument--const reference of the same type as the class

```C++
Cube::Cube(const Cude& c)
```

用处

- function parameter(by value)
- function return(by value)
- initialize

```C++
Cube foo(){
    Cube c;
    return c;
}
int main{
    Cube c2 = foo();
    return 0;
}
```

OUPUT:

default constructor

copy constr

copy constr

```C++
int main(){
    Cube c;//default
    Cube c2=c;//copy--Initializeing
}

int main(){
    Cube c;//default
    Cube c2;//default

    c2=c;//no Constructor---not initializing
    //assignment operator
}
```

Copy assignment operator
---------------------

assignment: =

automatic assignment: already constructed obj copy all member values to the another.

custom assignment operator:

- public function of the class
- has funciton name *operator=*
- return reference of the class type
- has exactly one argument -- must be const reference of the class type.

```C++
    Cube& Cube::operator=(const Cube& obj){
        length_ = obj.length_;
        return *this//instance of the class
    }
```

Vaiable storage
----------------

- directly  Cube c
- pointer   Cube *p = &c
- refernce  Cube &r = c

reference is an alias to *existing* memory, **does not store memory itself**

```C++
Cube c;
Cube c2=c;//copy constr
Cube &c3=c;//alias
Cube *c4=&c;//extra pointer memory
```

function transform parameters by value, by reference, by pointer

best not by value(copy constr)

return by value, by reference, by pointer

**Never return a reference to a stack variable crated on the stack of current function**

return by value best

```C++
bool sendcube(Cube c){}
int main(){
    Cube c;
    sendcube(c);
}


bool sendcube(Cube &c){}
int main(){
    Cube c;
    sendcube(c);
}

bool sendcube(Cube *c){}
int main(){
    Cube c;
    sendcube(&c);
}
```

Destructor
--------------

