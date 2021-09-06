## Golang Basic

## 程序结构

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

如果一个名字是在函数内部定义，那么它就只在函数内部有效。如果是在函数外部定义，那么将在当前包的所有文件中都可以访问。名字的开头字母的大小写决定了名字在包外的可见性。如果一个名字是大写字母开头的（译注：必须是在函数外部定义的包级名字；包级函数名本身也是包级名字），那么它将是导出的，也就是说可以被外部的包访问，例如fmt包的Printf函数就是导出的，可以在fmt包外部访问。包本身的名字一般总是用小写字母。

Four  major kinds  of declarations: `var`, `const`, `type`, and `func`.

请记住“`:=`”是一个变量声明语句，而“`=`”是一个变量赋值操作。局部变量，类型明确可推导

**指针** `*int`
在Go语言中，返回函数中局部变量的地址也是安全的。例如下面的代码，调用f函数时创建局部变量v，在局部变量地址被返回之后依然有效，因为指针p依然引用这个变量。
```go
var p = f()

func f() *int {
    v := 1
    return &v
}
```
`fmt.Println(f() == f()) // "false"`

**new函数**
另一个创建变量的方法是调用内建的new函数。表达式`new(T)`将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为`*T`。
用new创建变量和普通变量声明语句方式创建变量没有什么区别，除了不需要声明一个临时变量的名字外，我们还可以在表达式中使用new(T)。换言之，**new函数类似是一种语法糖**，而不是一个新的基础概念。

**完全不同于C++的生命周期机制**
因为一个变量的有效周期只取决于是否可达，因此一个循环迭代内部的局部变量的生命周期可能超出其局部作用域。同时，局部变量可能在函数返回之后依然存在。

编译器会`自动选择`在栈上还是在堆上分配局部变量的存储空间，但可能令人惊讶的是，**这个选择并不是由用var还是new声明变量的方式决定的**。
例子
```go
var global *int

func f() {
    var x int
    x = 1
    global = &x
}

func g() {
    y := new(int)
    *y = 1
}
```
f函数里的x变量必须在堆上分配，因为它在函数退出后依然可以通过包一级的global变量找到，虽然它是在函数内部定义的；`*y`并没有从函数g中逃逸，编译器可以选择在栈上分配`*y`的存储空间（译注：也可以选择在堆上分配，然后由Go语言的GC回收这个变量的内存空间）,虽然这里用的是new方式

**类型**
`type 类型名字 底层类型`
```go
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
    AbsoluteZeroC Celsius = -273.15 // 绝对零度
    FreezingC     Celsius = 0       // 结冰点温度
    BoilingC      Celsius = 100     // 沸水温度
)
```
我们在这个包声明了两种类型：Celsius和Fahrenheit分别对应不同的温度单位。它们虽然有着相同的底层类型float64，但是它们是不同的数据类型，因此它们不可以被相互比较或混在一个表达式运算。刻意区分类型，可以避免一些像无意中使用不同单位的温度混合计算导致的错误；因此需要一个类似Celsius(t)或Fahrenheit(t)形式的显式转型操作才能将float64转为对应的类型。
```go
var c Celsius
var f Fahrenheit
fmt.Println(c == f)          // compile error: type mismatch
fmt.Println(c == Celsius(f)) // "true"!
```

不要将作用域和生命周期混为一谈。声明语句的作用域对应的是一个源代码的文本区域；它是一个**编译**时的属性。一个变量的生命周期是指程序运行时变量存在的有效时间段，在此时间区域内它可以被程序的其他部分引用；是一个**运行**时的概念。

## 基础数据类型

**运算符**
`&^     位清空 (AND NOT)`
位操作运算符&^用于按位置零（AND NOT）：如果对应y中bit位为1的话, 表达式z = x &^ y结果z的对应的bit位为0，否则z对应的bit位等于x相应的bit位的值。

一般来说，需要一个显式的转换将一个值从一种类型转化为另一种类型，并且算术和逻辑运算的二元操作中必须是相同的类型。虽然这偶尔会导致需要很长的表达式，但是它消除了所有和类型相关的问题，而且也使得程序容易理解。

```go
o := 0666
fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
x := int64(0xdeadbeef)
fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
// Output:
// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

ascii := 'a'
fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
```
通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操作数，但是%之后的[1]副词告诉Printf函数再次使用第一个操作数。第二，%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。

**字符串**
文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列
```go
s := "hello, world"
fmt.Println(s[:5]) // "hello"
fmt.Println(s[7:]) // "world"
fmt.Println(s[:])  // "hello, world"
```
**一个字符串包含的字节序列永远不会被改变**，当然我们也可以给一个字符串变量分配一个新字符串值。不变性意味着如果两个字符串共享相同的底层数据的话也是安全的，这使得复制任何长度的字符串代价是低廉的。同样，一个字符串s和对应的子字符串切片s[7:]的操作也可以安全地共享相同的内存，因此字符串切片操作代价也是低廉的。在这两种情况下都没有必要分配新的内存
![字符串底层数据](./pictures/字符串底层数据.png)

**常量**
因为常量的值是在编译期就确定的，因此常量可以是构成类型的一部分，例如用于指定数组类`var p [IPv4Len]byte`

## 复合数据类型
数组和结构体是聚合类型；它们的值由许多元素或成员字段的值组成。数组是由同构的元素组成——每个数组元素都是完全相同的类型——结构体则是由异构的元素组成的。数组和结构体都是有固定内存大小的数据结构。相比之下，slice和map则是动态的数据结构，它们将根据需要动态增长。
### 数组
数组是一个由固定长度的特定类型元素组成的序列.
```go
var r [3]int = [3]int{1, 2}
q := [...]int{1, 2, 3}
fmt.Printf("%T\n", q) // "[3]int"
```
数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型。数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定。
**比较**
也可以指定一个索引和对应值列表的方式初始化`r := [...]int{99: -1}`
如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的，这时候我们可以直接通过==比较运算符来比较两个数组，只有当两个数组的所有元素都是相等的时候数组才是相等的。不相等比较运算符!=遵循同样的规则。

## Slice
Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。一个slice类型一般写作[]T，其中T代表slice中元素的类型；slice的语法和数组很像，只是没有固定长度而已。

slice和数组的字面值语法很类似，它们都是用花括弧包含一系列的初始化元素，但是对于slice并没有指明序列的长度。这会隐式地创建一个合适大小的数组，然后slice的指针指向底层的数组。
内置的make函数创建一个指定元素类型、长度和容量的slice。容量部分可以省略，在这种情况下，容量将等于长度。
```go
make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]
```
在底层，make创建了一个匿名的数组变量，然后返回一个slice；只有通过返回的slice才能引用底层匿名的数组变量。

数组和slice之间有着紧密的联系。一个slice由三个部分构成：指针、长度和容量。指针指向第一个slice元素对应的底层数组元素的地址，要注意的是slice的第一个元素并不一定就是数组的第一个元素。长度对应slice中元素的数目；长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置。内置的len和cap函数分别返回slice的长度和容量。
![](./pictures/slice和数组.png)
因为slice值包含指向第一个slice元素的指针，因此**向函数传递slice将允许在函数内部修改底层数组的元素**。换句话说，复制一个slice只是对底层的数组创建了一个新的slice别名。同时，应该认识到**slice赋值O(1)时间复杂度，非常方便赋值，创建**

**append**
一样涉及capacity和length之间的扩张问题，考虑max(oldlength*2, newcapacity)作为新的capacity，如果无需扩容直接`z = x[:zlen]`，slice扩张非常方便；如果需要扩容，先扩容复制。

slice底层删除可以用后边覆盖前边，然后截取前边的思路做

### map
哈希表是一种巧妙并且实用的数据结构。它是一个无序的key/value对的集合，其中所有的key都是不同的，然后通过给定的key可以在常数时间复杂度内检索、更新或删除对应的value。在Go语言中，一个map就是**一个哈希表的引用**，map类型可以写为map[K]V，其中K和V分别对应key和value。**其中K对应的key必须是支持==比较运算符的数据类型**
```go
var ages map[string]int //nil map
ages := make(map[string]int) // mapping from strings to ints
ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}
map[string]int{}//空类型
delete(ages, "alice") // remove element ages["alice"]
```
如果元素类型是一个数字，你可能需要区分一个已经存在的0，和不存在而返回零值的0，可以像下面这样测试
`if age, ok := ages["bob"]; !ok { /* ... */ }`

### 结构体
结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。
点操作符也可以和指向结构体的指针一起工作：
**初始化**
```go
p := Point{1, 2}
anim := gif.GIF{LoopCount: nframes}
isafunction(Point{1,2})//匿名
```
```go
var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"
//相当于下面语句
(*employeeOfTheMonth).Position += " (proactive team player)"
```
结构体值传递并不方便，**所以请多注意指针的使用**，go中结构体的指针使用还挺方便的
如果结构体成员名字是以大写字母开头的，那么该成员就是导出的；这是Go语言导出规则决定的。一个结构体可能同时包含导出和未导出的成员。（weak数据保护）
如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，那样的话两个结构体将可以使用==或!=运算符进行比较。由此可知道，struct可以作为map的key

**构体嵌入和匿名成员**
```go
type Point struct{
    X,Y int
}

type Circle struct {
    Point//匿名成员
    Radius int
}

type Wheel struct {
    Circle//匿名成员 仍然是has a的关系
    Spokes int
}
```
可以简化嵌入结构体中叶子节点访问
```go
var w Wheel
w.X = 8            // equivalent to w.Circle.Point.X = 8
w.Y = 8            // equivalent to w.Circle.Point.Y = 8
w.Radius = 5       // equivalent to w.Circle.Radius = 5
w.Spokes = 20
```
在右边的注释中给出的显式形式访问这些叶子成员的语法依然有效，因此匿名成员并不是真的无法访问了。其中匿名成员Circle和Point都有自己的名字——就是命名的类型名字——但是这些名字在点操作符中是可选的。我们在访问子成员的时候可以忽略任何匿名成员部分。
但是结构体字面值必须遵循形状类型声明时的结构
```go
w = Wheel{Circle{Point{8, 8}, 5}, 20}

w = Wheel{
    Circle: Circle{
        Point:  Point{X: 8, Y: 8},
        Radius: 5,
    },
    Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
}
```
**因为成员的名字是由其类型隐式地决定的，所以匿名成员也有可见性的规则约束。**

### JSON
[json](https://docs.hacknode.org/gopl-zh/ch4/ch4-05.html)

## 函数
函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。
```go
func name(parameter-list) (result-list) {
    body
}
```
backtracking使用slice非常方便
> 当outline调用自身时，被调用者接收的是stack的拷贝。被调用者对stack的元素追加操作，修改的是stack的拷贝，其可能会修改slice底层的数组甚至是申请一块新的内存空间进行扩容；但这个过程并不会修改调用方的stack。因此当函数返回时，调用方的stack与其调用自身之前完全一致。
```go
func outline(stack []string, n *html.Node) {
    if n.Type == html.ElementNode {
        stack = append(stack, n.Data) // push tag
        fmt.Println(stack)
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        outline(stack, c)
    }
}
```
大部分编程语言使用固定大小的函数调用栈，常见的大小从64KB到2MB不等。固定大小栈会限制递归的深度，当你用递归处理大量数据时，需要避免栈溢出；除此之外，还会导致安全性问题。与此相反，Go语言使用可变栈，栈的大小按需增加(初始时很小)。这使得我们使用递归时不必考虑溢出和安全问题。

### Error
对于大部分函数而言，永远无法确保能否成功运行。这是因为错误的原因超出了程序员的控制。举个例子，任何进行I/O操作的函数都会面临出现错误的可能，只有没有经验的程序员才会相信读写操作不会失败，即使是简单的读写。因此，当本该可信的操作出乎意料的失败后，我们必须弄清楚导致失败的原因。
在Go的错误处理中，错误是软件包API和应用程序用户界面的一个重要组成部分，**程序运行失败仅被认为是几个预期的结果之一,而非异常**
```go
fmt.Println(err)
fmt.Printf("%v", err)
```
使用`fmt.Errorf("parsing %s as HTML: %v", url,err)`补充错误信息，fmt.Errorf函数使用`fmt.Sprintf`格式化错误信息并返回
**处理错误的策略**

- 传递错误信息`fmt.Errorf("parsing %s as HTML: %v", url,err)`
- 有限时间/次数的重试
- 在main函数中输出错误信息并终止程序
```go
fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
os.Exit(1)
//等价于
log.Fatalf("Site is down: %v\n", err)
```
- 只输出错误信息，继续
- 忽略错误 - 不要这么做，或者写明意图
### 函数变量
```go
func square(n int) int { return n * n }
f:=square
```
类似于C++函数指针，函数变量可以作为函数参数传递行为
```go
func forEachNode(n *html.Node, pre, post func(n *html.Node)){
    // pre, post 传函数实现前序和后序遍历--泛型
}
```

### 匿名函数
类似于lambda函数plus，匿名函数核心在于**函数中定义函数**，优势在于闭包，使用函数之外的变量（自由变量），相比lambda函数不限制长度。匿名函数可以赋值名称，从而递归调用。
**注意！警告：捕获迭代变量**
```go
var rmdirs []func()
for _, dir := range tempDirs() {
    os.MkdirAll(dir, 0755)
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir) // NOTE: incorrect!
    })
}
```
在上面的程序中，for循环语句引入了新的词法块，循环变量dir在这个词法块中被声明。在该循环中生成的所有函数值都共享相同的循环变量。**需要注意，函数值中记录的是循环变量的内存地址，而不是循环变量某一时刻的值**。以dir为例，后续的迭代会不断更新dir的值，当删除操作执行时，for循环已完成，dir中存储的值等于最后一次迭代的值。这意味着，每次对os.RemoveAll的调用删除的都是相同的目录。
需要做
```go
for _, dir := range tempDirs() {
    dir := dir // declares inner dir, initialized to outer dir
    // ...
}
```
这个问题不仅存在基于range的循环，在下面的例子中，**对循环变量i的使用也存在同样的问题**：(和C++完全不同)
```go
for i := 0; i < len(dirs); i++ {
    os.MkdirAll(dirs[i], 0755) // OK
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dirs[i]) // NOTE: incorrect!
    })
}
```
循环中只能记录地址，所以使用即刻状态时候可以，不能储存为下一刻的状态

### 变长函数
```go
func sum(vals...int) int {}
fmt.Println(sum(1, 2, 3, 4)) // "10"
values := []int{1, 2, 3, 4}
fmt.Println(sum(values...)) // "10"
```

## 常用系统包

### time
[time.Duration](https://studygolang.com/articles/12617)
```go
func Test() {
    var waitFiveHundredMillisections time.Duration = 500 * time.Millisecond

    startingTime := time.Now().UTC()
    time.Sleep(600 * time.Millisecond)
    endingTime := time.Now().UTC()

    var duration time.Duration = endingTime.Sub(startingTime)

    if duration >= waitFiveHundredMillisections {
        fmt.Printf("Wait %v\nNative [%v]\nMilliseconds [%d]\nSeconds [%.3f]\n", waitFiveHundredMillisections, duration, duration.Nanoseconds()/1e6, duration.Seconds())
    }
}
```
