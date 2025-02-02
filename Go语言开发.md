---
title: Go语言开发
data: 2025-01-22 23:56
categories: 
- 知识
- 编程语言
description: go语言开发相关知识
---

# Go语言开发

## 基本语法

### 包

每个 Go 程序都由包构成。程序从 `main` 包开始运行。

#### 导入包

```go
import(
	"fmt"
	"math/rand"
)
```

#### 导出名

在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。

在导入一个包时，你只能引用其中已导出的名字。 任何「未导出」的名字在该包外均无法访问。

### 函数

```go
func add(x int, y int)
{
    return x + y
}
```

#### 参数

注意类型在变量名的**后面**。

当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。如：`x int, y int`可简写为`x ,y int`

#### 多返回值

函数可以返回任意数量的返回值。

```go
func swap(x int, y int)
{
    return y, x;
}
```

没有参数的 `return` 语句会直接返回已命名的返回值，也就是「裸」返回值。

### 变量

`var` 语句用于声明一系列变量。和函数的参数列表一样，类型在最后。

```go
var c, python, java bool
```

和导入语句一样，变量声明也可以「分组」成一个代码块。

```go
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

#### 变量的初始化

变量声明可以包含初始值，每个变量对应一个。如果提供了初始值，则类型可以省略；变量会从初始值中推断出类型。

```go
var c, python, java = true, false, "no!"
```

#### 短变量声明

**在函数中**，短赋值语句 `:=` 可在隐式确定类型的 `var` 声明中使用。函数外的每个语句都 **必须** 以关键字开始（`var`、`func` 等），因此 `:=` 结构不能在函数外使用。

```go
c, python, java := true, false, "no!"
```

#### 基本类型

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
     // 表示一个 Unicode 码位

float32 float64

complex64 complex128
```

#### 零值

没有明确初始化的变量声明会被赋予对应类型的 **零值**。

零值是：

- 数值类型为 `0`，
- 布尔类型为 `false`，
- 字符串为 `""`（空字符串）。

#### 类型转换

表达式 `T(v)` 将值 `v` 转换为类型 `T`。

#### 常量

常量的声明与变量类似，只不过使用 `const` 关键字。

常量可以是字符、字符串、布尔值或数值。

常量不能用 `:=` 语法声明。

### 循环

go语言只有`for`循环一种

```go
for i := 0; i < 10; i++ {
	sum += i
}
```

初始化语句和后置语句是可选的。

此时你可以去掉分号，因为 C 的 `while` 在 Go 中叫做 `for`。

### 判断

```go
if x < 0 {
	return sqrt(-x) + "i"
}

if v := math.Pow(x, n); v < lim {
	return v
}
else {
	fmt.Printf("%g >= %g\n", v, lim)
}
```

和 `for` 一样，`if` 语句可以在条件表达式前执行一个简短语句。

**该语句声明的变量作用域仅在 `if` 之内。**

`switch` 语句是编写一连串 `if - else` 语句的简便方法。它运行第一个 `case` 值 值等于条件表达式的子句。

Go 的 `switch` 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只会运行选定的 `case`，而非之后所有的 `case`。在效果上，Go 的做法相当于这些语言中为每个 `case` 后面自动添加了所需的 `break` 语句。

```go
switch os := runtime.GOOS; os {
case "darwin":
	fmt.Println("macOS.")
case "linux":
	fmt.Println("Linux.")
default:
	// freebsd, openbsd,
	// plan9, windows...
	fmt.Printf("%s.\n", os)
```

### defer推迟

defer 语句会将函数推迟到外层函数返回之后执行。

推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。

```go
defer fmt.Println("world")
```

推迟调用的函数**调用会被压入一个栈中**。 当外层函数返回时，被推迟的调用会按照**后进先出**的顺序调用。

### 指针

Go 拥有指针。指针保存了**值的内存地址**。

类型 `*T` 是指向 `T` 类型值的指针，其零值为 `nil`。

`&` 操作符会生成一个指向其操作数的指针。也就是这个值的地址。

`*` 操作符表示指针指向的底层值。

与 C 不同，**Go 没有指针运算。**

### 结构体

```go
type Vertex struct {
	X int
	Y int
}
```

结构体字段可通过点号 `.` 来访问。

结构体字段可通过结构体指针来访问。如果我们有一个指向结构体的指针 `p` 那么可以通过 `(*p).X` 来访问其字段 `X`。 不过这么写太啰嗦了，**所以语言也允许我们使用隐式解引用，直接写 `p.X` 就可以。**

使用 `Name:` 语法可以仅列出部分字段（字段名的顺序无关）。

特殊的前缀 `&` 返回一个指向结构体的指针。

### 数组

类型 `[n]T` 表示一个数组，它拥有 `n` 个类型为 `T` 的值。

```go
var a [10]int
```

数组不能改变大小

### 切片

类型 `[]T` 表示一个元素类型为 `T` 的切片。

切片通过两个下标来界定，一个下界和一个上界，二者以冒号分隔：

```go
a[low : high]
```

它会选出一个半闭半开区间，包括第一个元素，**但排除最后一个元素。**

#### 长度和容量

切片拥有 **长度** 和 **容量**。

切片的长度就是它所包含的元素个数。

切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

切片 `s` 的长度和容量可通过表达式 `len(s)` 和 `cap(s)` 来获取。

切片的零值是 `nil`。

nil 切片的长度和容量为 0 且没有底层数组。

#### 切片的创建

切片可以用内置函数 `make` 来创建，这也是你创建动态数组的方式。

`make` 函数会分配一个元素为零值的数组并返回一个引用了它的切片：

```go
a := make([]int, 5)  // len(a)=5
```

要指定它的容量，需向 `make` 传入第三个参数：

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
```

#### 追加元素

为切片追加新的元素是种常见的操作，为此 Go 提供了内置的 `append` 函数。

#### range遍历

`for` 循环的 `range` 形式可遍历切片或映射。

```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

可以将下标或值赋予 `_` 来忽略它。

```go
for i, _ := range pow
for _, value := range pow
```

### map映射

`map` 映射将键映射到值。

映射的零值为 `nil` 。`nil` 映射既没有键，也不能添加键。

`make` 函数会返回给定类型的映射，并将其初始化备用。

```go
m = make(map[string]Vertex)
m["Bell Labs"] = Vertex{
	40.68433, -74.39967,
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
```

若顶层类型只是一个类型名，那么你可以在字面量的元素中省略它。

```go
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google": {37.42202, -122.08408},
}
```

#### 修改映射

在映射 `m` 中插入或修改元素：

```go
m[key] = elem
```

获取元素：

```go
elem = m[key]
```

删除元素：

```go
delete(m, key)
```

通过双赋值检测某个键是否存在：

```go
elem, ok = m[key]
```

若 `key` 在 `m` 中，`ok` 为 `true` ；否则，`ok` 为 `false`。

若 `key` 不在映射中，则 `elem` 是该映射元素类型的零值。

### 函数值

函数值可以用作函数的参数或返回值。

### 函数闭包

Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。 该函数可以访问并赋予其引用的变量值，换句话说，该函数被“绑定”到了这些变量。

### JSON处理

 ```go
 type userInfo struct	// 每个变量的命名都要大写
 {
 	Name string
     Age int 'json:"age"'//这里可以将输出的Age变小写
     Hobby []string
 }
 func main(){
     a:= userInfo{Name:"Wang", Age:18, Hobby:[]string{"Golang", "TypeScript"}}
    buf, err := json.Marshal()	// json序列化
     err = json.Unmarshak(buf, &b)	//json反序列化
 }
 
 ```

### 时间处理

```go
now := time.Now() // 获取当前时间
t := time.Date(2022, 3, 27, 1, 25, 36, 0, time.UTC)
fmt.Println(t.Year())
fmt.Println(t.Format("2006-01-02 15:04:05")) //格式化时间，解析时间
t3, err := time.Parse("2006-01-02 15:04:05", "2022-03-27 01:25:36") 	// 解析时间
fmt.Println(now.Unix) // 获取时间戳
```

### 数字解析

```go
strconv.ParseFloat("1.234", 64) // 字符串转float
strconv.ParseInt("111", 10, 64)	// 字符串转int
strconv.Atoi("111")	 // 字符串转int
```



## 方法和接口

### 方法

Go 没有类。不过你可以为类型定义方法。

方法就是一类带特殊的 **接收者** 参数的函数。

```go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { // float64为返回值类型
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

记住：方法只是个带接收者参数的函数。

**接收者的类型定义和方法声明必须在同一包内。**

#### 指针类型的接收者

指针接收者的方法可以修改接收者指向的值（如这里的 `Scale` 所示）。 由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。

```go
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
```

若使用值接收者，那么 `Scale` 方法会对原始 `Vertex` 值的副本进行操作。（对于函数的其它参数也是如此。）

#### 方法与指针重定向

接收者为指针的的方法被调用时，接收者既能是值又能是指针。

```go
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

v.Scale(2) // 这样调用也是成立的
```

以值为接收者的方法被调用时，接收者既能为值又能为指针。

#### 使用指针接收者优点

使用指针接收者的原因有二：

首先，方法能够修改其接收者指向的值。

其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样会更加高效。

### 接口

```go
// 接口
type Shape interface {
    Area() float64
    Perimeter() float64
}


type Rectangle struct {
    Width, Height float64
}

// 实现 Area 方法
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// 实现 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

```

某个类型实现了某个 接口的所有方法，则该类型 隐式地实现了相应的接口

同时注意指针类型定义和值类型定义会导致接口类型的区别

```go
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	v := Vertex{3, 4}
	a = &v // a *Vertex 实现了 Abser
	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	a = v

```

#### 隐式实现

在 Go 中，不需要显式声明类型实现了哪个接口。

只要类型的方法集满足接口定义，它就实现了该接口。

#### 接口值

接口也是值。它们可以像其它值一样传递。

接口值可以用作函数的参数或返回值。

在内部，接口值可以看做包含值和具体类型的元组：

```go
(value, type)
```

接口值保存了一个具体底层类型的具体值。

接口值调用方法时会执行其底层类型的同名方法。

#### 底层值为nil的接口值

即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

**注意:** 保存了 nil 具体值的接口其自身并不为 nil。

#### 空接口

指定了零个方法的接口值被称为 *空接口：*

```go
interface{}
```

空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）

空接口被用来处理未知类型的值。

#### 类型断言

**类型断言** 提供了访问接口值底层具体值的方式。

```go
t := i.(T)
```

该语句断言接口值 `i` 保存了具体类型 `T`，并将其底层类型为 `T` 的值赋予变量 `t`。

若 `i` 并未保存 `T` 类型的值，**该语句就会触发一个 panic。**

为了 **判断** 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。

```
t, ok := i.(T)
```

若 `i` 保存了一个 `T`，那么 `t` 将会是其底层值，而 `ok` 为 `true`。

否则，`ok` 将为 `false` 而 `t` 将为 `T` 类型的零值，**程序并不会产生 panic。**

#### 类型选择

**类型选择** 是一种按顺序从几个类型断言中选择分支的结构。

类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）， 它们针对给定接口值所存储的值的类型进行比较。

```go
switch v := i.(type) {
case T:
    // v 的类型为 T
case S:
    // v 的类型为 S
default:
    // 没有匹配，v 与 i 的类型相同
}
```

类型选择中的声明与类型断言 `i.(T)` 的语法相同，只是具体类型 `T` 被替换成了关键字 `type`。

#### Stringer

[`fmt`](https://go-zh.org/pkg/fmt/) 包中定义的 [`Stringer`](https://go-zh.org/pkg/fmt/#Stringer) 是最普遍的接口之一。

```go
type Stringer interface {
    String() string
}
```

`Stringer` 是一个可以用字符串描述自己的类型。`fmt` 包（还有很多包）都通过此接口来打印值。

#### 错误

Go 程序使用 `error` 值来表示错误状态，可以使用`error`作为错误提示信息，返回错误值。

与 `fmt.Stringer` 类似，`error` 类型是一个内建接口：

```go
type error interface {
    Error() string
}
```

#### Readers

`io` 包指定了 `io.Reader` 接口，它表示数据流的读取端。

Go 标准库包含了该接口的[许多实现](https://cs.opensource.google/search?q=Read\(\w%2B\s\[\]byte\)&ss=go%2Fgo)，包括文件、网络连接、压缩和加密等等。

`io.Reader` 接口有一个 `Read` 方法：

```go
func (T) Read(b []byte) (n int, err error)
```

#### 图像

[`image`](https://go-zh.org/pkg/image/#Image) 包定义了 `Image` 接口：

```go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

**注意:** `Bounds` 方法的返回值 `Rectangle` 实际上是一个 [`image.Rectangle`](https://go-zh.org/pkg/image/#Rectangle)，它在 `image` 包中声明。

## 泛型

### 泛型参数

可以使用类型参数编写 Go 函数来处理多种类型。 函数的类型参数出现在函数参数之前的方括号之间。

```go
func Index[T comparable](s []T, x T) int
```

此声明意味着 `s` 是满足内置约束 `comparable` 的任何类型 `T` 的切片。 `x` 也是相同类型的值。

`comparable` 是一个有用的约束，它能让我们对任意满足该类型的值使用 `==` 和 `!=` 运算符。

### 泛型类型

除了泛型函数之外，Go 还支持泛型类型。 类型可以使用类型参数进行参数化，这对于实现通用数据结构非常有用。

## 多线程

### 概念

Go 程（goroutine）是由 Go 运行时管理的轻量级线程。

```go
go f(x, y, z)
```

会启动一个新的 Go 协程并执行

```go
f(x, y, z)
```

`f`, `x`, `y` 和 `z` 的求值发生在当前的 Go 协程中，而 `f` 的执行发生在新的 Go 协程中。

Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。

### 信道

信道是带有类型的管道，你可以通过它用信道操作符 `<-` 来发送或者接收值。

```go
ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
```

（“箭头”就是数据流的方向。）

和映射与切片一样，信道在使用前必须创建：

```go
ch := make(chan int)
```

默认情况下，发送和接收操作在另一端准备好之前都会阻塞。**这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。**

#### 带缓冲的信道

信道可以是 **带缓冲的**。将缓冲长度作为第二个参数提供给 `make` 来初始化一个带缓冲的信道：

```go
ch := make(chan int, 100)
```

仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。

#### 关闭信道

发送者可通过 `close` 关闭一个信道来表示没有需要发送的值了。接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完

```go
v, ok := <-ch
```

此时 `ok` 会被设置为 `false`。

### select语句

`select` 语句使一个 Go 程可以等待多个通信操作。

`select` 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。

#### 默认选择

当 `select` 中的其它分支都没有准备好时，`default` 分支就会执行。

为了在尝试发送或者接收时不发生阻塞，可使用 `default` 分支：

### 同步

Go 标准库中提供了 [`sync.Mutex`](https://go-zh.org/pkg/sync/#Mutex) 互斥锁类型及其两个方法：

- `Lock`
- `Unlock`

我们可以通过在代码前调用 `Lock` 方法，在代码后调用 `Unlock` 方法来保证一段代码的互斥执行。参见 `Inc` 方法。

我们也可以用 `defer` 语句来保证互斥锁一定会被解锁。参见 `Value` 方法。











