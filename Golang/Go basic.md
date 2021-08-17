## Golang Basic

### 语法规则

在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。

Go 的基本类型有
```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
    // 表示一个 Unicode 码点

float32 float64

complex64 complex128
```

for
```go
for initialization; condition; post{
    // zero or more statements
}
```

If  an entity is declared within  a  function,  it is local to  that function.  If declared outside  of a function,  however,  it is visible  in  all files  of the  package to which it belongs.  The case of  the first letter of a name determines  its  visibilit yacrosspackagebound aries.  Ifthe  namebeg inswith anupp er-case letter,  itisexpor ted,whichmeans thatitisvisible andaccessibleoutside ofitsown packageand  may berefer red tobyother parts of  the  program, as with Printfin the fmt package.  Packagenames themselves are always in lowercase.

Four  major kinds  of declarations: `var`, `const`, `type`, and `func`.

