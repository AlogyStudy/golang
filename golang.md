# Go介绍

**并发支持**,**垃圾回收**的编译型系统编程语言。

特点：

- 类型安全 和内存安全
- 以非常直观和极低代价的方案实现高并发
- 高效的垃圾回收机制
- 快速编译 （同时解决C语言中头文件太多的问题）
- 为多核计算机提供性能提升的方案
- UTF-8编码支持

> Go环境变量


![clipboard.png](/img/bVS9gF)

```
GOEXE=.exe // 形成可执行文件的后缀
GOPATH // 工作目录
```

GOPATH下约定俗成的目录：
```
bin // 存放编译后生成的可执行文件
pkg // 存放编译后生成的包文件
src // 存放项目源码
```

> 常用命令

```
go get // 获取远程包 （git或hg（如果是谷歌code上托管））
go run // 直接运行程序
go build // 编译，检查是否有编译错误
go fmt // 格式化源码
go install // 编译包文件并且编译整个程序
go test // 运行测试文件
go doc // 查看文档
godoc -http=:8080 // 查看文档
```

# GO程序结构


![clipboard.png](/img/bVS9mN)

```
mpathapp // 可执行文件存放位置
math.a // 包文件
```

# 基础知识

hello.go
```
package main
import (
    "fmt"
)
func main () {
    fmt.Println("hello world")
}
```

`.go`文件的一般结构

Go程序是通过`package`来组织的
只有`package`名称为`main`的包可以包含`main`函数
一个可执行程序**有且仅有**一个`main`包

通过`import`关键字来导入其它非`main`包
通过`const`关键字来定义常量
通过在函数体外部使用`var`关键字来进行全局变量的声明和赋值
通过`type`关键字来进行结构`struct`或接口`interface`的声明
通过`func`关键字来进行函数的声明

```
// 当前程序的包名
package main

// 导入其它的包
// import . "fmt" // 省略调用
import (
	"fmt"
	"io"
	"os"
	"time"
	"strings"
)

// 常量的定义
const PI = 3.141592653

// 全局变量的声明与赋值
var name = "zf"

// 一般类型声明
type newType int

// 结构的声明
type struc struct{} // 接口关键字和大括号不能有空格

// 接口的声明
type inter interface{} // 接口关键字和大括号不能有空格

// 由 main 函数作为程序入口点启动
func main () {
	fmt.Println("hello world!")
}
```

如果导入包之后，未调用其中的函数或者类型将会报出编译错误.


> package 别名

当使用第三方包时，包名可能会非常接近或者相同，此时就可以使用别名来进行区别和调用

```
import std "fmt" // 别名

std.Println("hello world")
```


> 注释

```
// 单行注释
/* 多行注释 */ 
```


> 可见性规则


Go语言中，使用**大小写**来决定该常量，变量，类型，接口，结构，或函数是否可以被外部包所调用：根据约定，函数名首字母小写即为`private`;函数名首字母大写即为`public`


# 类型与变量

## 基本类型

> 布尔型：bool

长度：1字节
取值范围：true，false
注意：不可以用数字代表true或false

> 整型：int/uint
根据运行平台可能为32或64位

- 8位整型：int8/uint8
    长度：1字节
    取值范围：-128~127/0~255
- 16位整型：int16/uint16
    长度：2字节
	取值范围：-32768~32767/0~65535
- 32位整型：int32（rune）/uint32
	长度：4字节
	取值范围：-2^32/2~2^32/2-1/0~2^32-1
- 64位整型：int64/uint64
	长度：8字节
	取值范围：-2^64/2~2^64/2-1/0~2^64-1
- 字节型：byte（uint8别名）

从严格意义上讲`type newint int` 这里的`newint`ing不能说是`int`的别名，而是底层数据结构相同，在这里自定义类型，在进行类型转换时扔旧需要显示转换，但`byte`和`rune`确实为`uint8`和`int32`的别名，可以相互进行转换。

> 浮点型：float32/float64

长度：4/8字节
小数位：精确到7/15小数位

> 复数：complex64/complex128
长度：8/16字节




> 足够保存指针的 32 位或 64 位整数型：uintptr


> 其它值类型

array、struct、string

> 引用类型

slice、map、chan

> 接口类型

interface

> 函数类型

func


## 类型零值

零值并不等于空值，而是当变量被声明为某种类型后的默认值。

通常情况下：
- 值类型的默认值是为0
- bool为false
- string为空字符串


## 变量

单个变量的声明与赋值

变量的声明格式：`var <变量名称> <变量类型>`
变量的赋值格式: `<变量名称> = <表达式>`
声明的同时赋值：`var <变量名称> [变量类型] = <表达式>`
写法：
```
var a int // 变量声明
a = 10 // 变量赋值

var b int = 20 // 变量声明的同时赋值
var b = 1 // 变量声明与赋值，由系统推荐是那种类型

b := 10 // 函数中的变量声明与赋值的最简写法

var 是全局的变量
:= 只能在函数中使用，局部变量
```

多个变量的声明与赋值

全局变量的声明可使用`var()`的方式进行简写
全局的变量的声明不可以省略var，但`可使用并行方式`
所有变量都可以使用类型推断
局部变量不可以使用`var()`的方式简写，只能使用并行方式

```
var (
	// 使用常规方式
	aaa = "hello"
	// 使用并行方式以及类型推断
	a, b = 1, 2
	// cc := 2 // 不可以省略 var
)


func main () {
	// 多个变量的声明
	var a, b, c, d int
	// 多个变量的赋值
	a, b, c, d = 1, 2, 3, 4

	// 多个变量声明的同时赋值
	var e, f, g, h int = 5, 6, 7, 8
	// 省略变量类型，由系统推断
	var i, j, k, l = 9, 10, 11, 12
	// 多个变量声明与赋值的最简写法
	i, m, n, o := 13, 14, 15, 16

    _, dd = 10, 20 // 空白符号，省略该表达式赋值(应用函数返回值)
}
```

## 类型转换

- Go中不存在隐式转换，都是显示声明（Go的类型安全）
- 转换只能发生在两种相互兼容的类型之间
- 类型转换的格式：
	<ValueA> [:]= <TypeOfValueA>(<ValueB>)

```
// 在相互兼容的两种类型之间转换
var a float32 = 1.1
b := int(a)


// 表达式无法通过编译
var c bool = true
d := int(c)
```
-----
```
package main

import (
	"fmt"
)

func main () {
	var a float32 = 100.01
	fmt.Println(a) // 100.01
	b := int(a)
	fmt.Println(b) // 100
}
```

整型无法和布尔型兼容
`float`类型无法和字符串类型兼容

`int`和`string`互转
```
var c int = 3
// d := string(c)
d := strconv.Itoa(c) // 字符串 3
c, _ = strconv.Atoi(d) // int 3
```

现象：
```
var a int = 65
string(a)
fmt.Println(a) // A
```
`string()`表示将数据转换成文本格式，因为计算机中存储的任何东西本质上都是数字，因此此函数自然的认为需要的是用数字65表示文本A


# 常量

- 常量的值在编译时就已经确定
- 常量的定义格式与变量基本相同
- 等号右侧必须是常量或者常量表达式
- 常量表达式中的函数必须是内置函数


> 常量的初始化规则与枚举

- 在定义常量组时，如果不提供初始值，则表示将使用上行的表达式
- 使用相同的表达式不代表具有相同的值
- `iota`是常量的计数器，从0开始，**组中每定义1个常量自动递增1**
- 通过初始化规则与`iota`可以达到枚举的效果
- 每遇到一个const关键字,`iota`就会重置为0
```
const (
	_A = "A"
	_B
	_C = iota
	_D
)
func main () {
	fmt.Println(_A, _B, _C, _D) // A A 2 3
}
```
-----
```
const (
	a = "123"
	b = len(a)
	c
)

func main () {
	fmt.Println(a, b, c) // 123, 3, 3
}
```

编译不通过：
```
var ss = "123"

const (
	a = len(ss)
	b
	c
)

func main () {
	fmt.Println(a, b, c)
}
```
错误信息：
```
# command-line-arguments
.\const.go:39: const initializer len(ss) is not a constant
```
-----
```
const (
	a, b = 1, "xixi"
	c
)

func main () {
	fmt.Println(a, b, c)
}
``` 
错误信息：
```
# command-line-arguments
.\const.go:40: extra expression in const declaration
```


# 运算符


Go中的运算符从左至右结合

优先级（从高到低）

- `^` `!`  (一元运算符)
-  `*` `/` `%` `<<` `>>` `&` `&^` (二元运算符)
- `+` `-` `|` `^` (二元运算符)
- `==` `!=` `<` `<=` 	`>=` `>` (二元运算符)
- `<-`  (专门用于channel)
- `&&`
- `||`

```
fmt.Println(1 ^ 2) // 二元运算符
fmt.Println(^2) // 一元运算符

/*

	6: 0110
 11: 1101
------------
  &	 0010  // 2
  |  1111  // 15
  ^  1101  // 13
  &^ 0100  // 4

	6 -> 110
	5 -> 101
	4 -> 100

13 / 2 = 1 // 6
1101
 */
```

# 控制语句

> 指针

Go虽然保留了指针，但与其他编程语言不同的是，在Go当中不支持指针运算以及`->`运算符，而是直接采用`.`选择符来操作指针目标对象的成员

- 操作符`&`取变量地址，使用`*`通过指针间接访问目标对象
- 默认值为`nil`而非`NULL`

```
package main

import (
	"fmt"
)

func main () {
	a := 1
	var p *int = &a
	fmt.Println(p) // 0xc0420361d0
	fmt.Println(*p) // 1
}
```

递增递减语句
在Go当中，`++`与`--`是作为语句而并不是作为表达式

> 判断if

- 条件表达式没有括号
- 支持一个初始化表达式（可以是并行方式）
- 左大括号必须和条件语句或else在同一行
- 支持单行模式
- 初始化语句中的变量为`block`级别，同时隐藏外部同名变量

```
func main () {
	a := 10
	if a := 0; a > 0 {
		fmt.Println(a)
	} else if a == 0 {
		fmt.Println(0111) // 73
	}
	fmt.Println(a) // 10
}
```

> 循环for

- Go只有for一个循环语句关键字，但支持3中形式
- 初始化和step表达式可以是多个值
- 条件语句每次循环都会重新检查，因此不建议在条件语句中使用函数，尽量提前计算好条件并以变量或常量代替
- 左大括号必须和条件语句在同一行

```
// 第一种形式
func main () {
	a := 1
	for {
		a++
		if a > 3 {
			break
		}
		fmt.Println(a) // 2, 3
	}
	fmt.Println(a) // 4
}
```
-----
```
// 第二种形式
func main () {
	a := 1
	for a <= 3 {
		a++
		fmt.Println(a) // 2, 3, 4
	}
	fmt.Println(a) // 4
}
```
-----
```
// 第三种形式
func main () {
	a := 1
	for i := 0; i < 3; i++ {
		a++
		fmt.Println(a) // 2, 3, 4
	}
	fmt.Println(a) // 4
}
```

> swtich

- 可以使用任何类型或表达式作为条件语句
- 不需要break，一旦条件符合自动终止
- 如希望继续执行下一个case，需使用fallthrough语句
- 支持下一个初始化表达式（可以是并行方式），右侧需跟分号
- 做大括号必须和条件语句在同一行

```
func main () {
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")	
	}
	fmt.Println(a)
}
```
-----
```
func main () {
	a := 1
	switch {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a>=1")
	}
	fmt.Println(a)
}
```
-----
```
func main () {
	switch a := 1; {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a>=1")
    default:
        fmt.Println("none")
	}
    fmt.Println(a) // undefined: a //for,if,switch都具有块级作用域
}
```

> 跳转语句goto，break，continue

- 三个语法都可以配合标签使用
- 标签名区分大小写，若不使用会造成编译错误
- `break`与`continue`配合标签可用于多层循环的跳出
- `goto`**调整执行位置**，与其它2个语句配合标签的结果并不相同

```
func main () {
LABEL:
	for {
		for i := 0; i < 10; i++ {
			if i > 2 {
				break LABEL
			} else {
				fmt.Println(i)
			}
		}
	}
}
```
-----
```
func main () {
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue LABEL
		}
	}	
}
```

# 数组Array

- 定义数组的格式：var <varName> [n]<type>, n>=0
- 数组的**长度也是类型的一部分**，因此具有不同长度的数组为不同类型
- 数组在Go中为**值类型**
- 数组之间可以使用`==`或`!=`进行比较，但不可以使用`<`或>``(相同类型之间，才可以使用相等或不能判断。也就是数组长度也要相同，长度也是数组类型的一部分)
- 可以使用new来创建数组，此方法返回一个指向数组的指针
- Go支持多维数组

创建数组
```
func main () {
	var a [2]string
	var b [1]int
	c := [2]int{11, 12}
	d := [20]int{19: 1}
	e := [...]int{1, 2, 3, 4, 5}
	f := [...]int{0: 11, 1: 22, 2: 33}
	b[0] = 10
	a[1] = "100"
    arr := [...]string{0: "xixi", 1: "hhh"}
    
	fmt.Println(a, b, c, d, e, f)
}
```
-----

```
p := new([10]int)
p[1] = 2
fmt.Println(&p) // 取地址
fmt.Println(*p) // 取值
```
-----
```
// 多维数组
a := [2][3]int{
	{1, 1, 1},
	{2, 2, 2},
}
fmt.Println(a)
```
冒泡排序：

```
func main () {
	a := [...]int{3, 4, 234, 2, 3, 5}
	fmt.Println(a)
	num := len(a)
	for i := 0; i< num; i++ {
		for j := i+1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}				
	}
	fmt.Println(a)
}
```

# 切片Slice

- 其本身并不是数组，它指向底层的数组
- 作为变长数组的代替方案，可以关联底层数组的局部或全部
- 为引用类型
- 可以直接创建或从底层数组获取生成
- 使用`len()`获取元素个数，`cap()`获取容量
- 一般使用`make()`创建
- 如果多个`slice`指向相同底层数组，其中一个的值改变会影响全部
- make([]T, len, cap)
	其中`cap`可以省略，则和`len`的值相同
	`len`表示存数的元素个数，`cap`表示容量

声明：
```
// 声明方法：
var s1 []int // 中括号中没有数字或`...`
fmt.Println(s1) // []
```

 reslice方法： 从数组中截取`Slice`
```
a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
fmt.Println(a)
s1 := a[5: len(a)] // 包含起始索引，不包含终止索引 // a[5 6 7 8 9]
s2 := a[5: ] // 包含起始索引，不包含终止索引 // a[5 6 7 8 9]
fmt.Println(s1, s2)
```
make方法 (一般使用make创建)
```
s1 := make([]int, 3, 10) // 10小块连续的内存,如果slice超过10，内存卡会继续申请，重新生成内存地址
s2 := make([]int, 10) // cap不给定，是slice的最大长度
fmt.Println(len(s1), cap(s1), s1) // 3 10 [0 0 0]
fmt.Println(len(s2), cap(s2), s2) // 10 10 [0 0 0 0 0 0 0 0 0 0]
```

Reslice

- `reslice`时索引以被`slice`的切片为准
- 索引不可以超过被`slice`的切片的容量`cap()`值
- 索引越界不会导致底层数组的重新分配而引发错误

Append
- 可以在`slice`尾部追加元素
- 可以将一个`slice`追加在另一个`slice`尾部
- 如果最终长度为超过追加到`slice`的容量则返回元素`slice`
- 如果超过追加到的`slice`的容量则将重新分配数组并拷贝原始数据

```
s1 := make([]int, 3, 6)
fmt.Println("%p\n", s1)
s1 = append(s1, 1, 2, 3, 4)

fmt.Println("%v %p\n", s1) //  [0 0 0 1 2 3 4] 
```
Copy

```
nt{1, 2, 3, 4, 5, 6, 7}
s2 := []int{8, 9}
copy(s1, s2)
// copy(s1, s2[1: 2])
fmt.Println(s1, s2) // [8 9 3 4 5 6 7] [8 9]
```

# Map

- 类似其它语言中的哈希表或者字典，以key-value形式存储的数据
- key必须是支持`==`或`!=`比较运算符的类型，不可以是**函数**，**map**或**slice**
- Map查找比线性搜索快很多，但比使用索引访问数据的类型慢100倍
- Map使用`make()`创建，支持`:=`简写方式
- 键值对不存在时自动添加，使用`delete()`删除某键值对
- 使用`for range`对`map`和`slice`进行迭代操作

- make([keyType]valueType, cap), cap表示容量，可省略
	超过容量时会自动扩容，但尽量提供一个合理的初始值
	使用`len()`获取元素个数

Map初始化
```
var m map[int]string
// m = map[int]string{}
m = make(map[int]string)
var m1 map[int]string = make(map[int]string)
m2 := make(map[int]string)

fmt.Println(m, m1) // map[]
```
删除和常用Map方法赋值
```
m2 := make(map[int]string)
m2[1] = "OK"
delete(m2, 1)
a := m2[1]

fmt.Println(a)
```

迭代：
```
for i,v := range  slice {
	// slice[i]
}

for k,v := range map {
	// map[k]
}
```
取Map中的key
```
m := map[int]string{1: "A", 2: "B", 3: "C", 4: "D"}
s := make([]int, len(m))
i := 0
for k,_ := range m {
	s[i] = k
	i++
}
sort.Ints(s)
fmt.Println(s)

```
Map中的 key-value互换:
```
m1 := map[int]string{1: "A", 2: "B", 3: "C"}
m2 := make(map[string]int)
// m2 := map[string]int{"A": 1, "B": 2, "C": 3}

for k,v := range m1 {
	m2[v] = k
}
fmt.Println(m1)
fmt.Println(m2)
```

# Function

函数function

Go 函数 不支持**嵌套**、**重载**和**默认参数**
支持的特性：
- 无需声明原型
- 不定长度变参
- 多返回值
- 命名返回值参数
- 匿名函数
- 闭包

定义函数使用关键字 **func**，且左大括号不能另起一行
函数也可以作为一种类型使用


闭包：
```
func closure (x int) func (int) int {
	return func (y int) int {
		return x + y
	}
}
```

> defer

- `defer`的执行方式类似其它语言中的析构函数，在函数体执行结束后按照调用顺序的**相反顺序**逐个执行
- 即使函数发生**严重错误**也会执行
- 支持匿名函数的调用
- 常用于**资源清理**、**文件关闭**、**解锁**以及**记录时间**等操作
- 通过与匿名函数配合可在return之后**修改**函数计算结果
- 如果函数体内某个变量作为`defer`时匿名函数的参数，则在定义defer时即已经获得了拷贝，否则则是引用某个变量的地址

- Go 没有异常机制，但有`panic/recover`模式来处理错误
- `panic`可以在任何地方引发，但`recover`**只有**在defer调用的函数中有效

defer使用：
```
func main () {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	// a, c, b
	
	
	for i := 0; i < 3; i++ {
		defer fmt.Println(i) // 2 1 0
	}
	
	for i := 0; i < 3; i++ {
		defer func () {
			fmt.Println(i) // 3 3 3
		}()
	}

}
```

`defer`要放在`panic()`之前：
```
func main () {
	A()
	B()
	C()
}

func A () {
	fmt.Println("func A")
}

func B () {
	defer func () {
		if err := recover(); err != nil {
			fmt.Println("Recover")
		}
	}()
	panic("Panic in B")
}

func C () {
	fmt.Println("func C")
}
```

# 结构struct

- Go 中的`struct`与C中的`struct`非常相似，并且Go没有`class`
- 使用 `type <Name> struct{}` 定义结构，名称遵循可见性规则
- 支持指向自身的指针类型成员
- 支持匿名结构，可用作成员或定义成员变量
- 匿名结构也可以用于`map`的值
- 可以使用字面值对结构进行初始化
- 允许直接通过指针来读写结构成员
- 相同类型的成员可进行直接拷贝赋值
- 支持 `==` 与 `!=`比较运算符，但不支持 `>` 或 `<`
- 支持匿名字段，本质上是定义了以某个类型名为名称的字段
- 嵌入结构作为匿名字段看起来像继承，但不是继承
- 可以使用匿名字段指针

```
type person struct {
	name string
	age int
}

func main () {
	a := person{}
	a.name = "zf"
	a.age = 23
	fmt.Println(a) // { 0}
}
```

`struct`也是值类型

对初始化结构`struct`使用地址符

```
type person struct {
	name string
	age int
}
func main () {
	a := &person{ // 调用结构使用地址符  // 字面值初始化
		name: "zf",
		age: 24,
	}
	a.name = "pink"
	// a.name = "zf"
	// a.age = 23
	fmt.Println(a) // { 0}
	// A(&a)
	A(a)
	B(a)
	fmt.Println(a)
}

func A (per *person) {
	per.age = 18
	fmt.Println("A", per)
}

func B (per *person) {
	per.age = 20
	fmt.Println("B", per)
}
```


匿名结构：

```
func main () {
	a := &struct {
		name string
		age int
	} {
		name: "tan",
		age: 19,
	}
	fmt.Println(a)
}
```

外层结构：
```
type person struct {
	name string
	age int
	contact struct {
		phone,city string
	}
}

func main () {
	b := person {
		name: "yellow",
		age: 18,
	}
	b.contact.phone = "123123"
	b.contact.city = "xiamen"

	fmt.Println(b)
}
```

匿名字段：

```
type p1 struct {
	string
	int
}

func main () {
	c := p1{"cyan", 20} // 字段的类型严格按照结构声明的字段

	fmt.Println(a, b, c)
}
```
匿名函数和匿名字段在函数中使用的次数非常少，没有必要声明，才会使用到。

嵌入（继承）结构：
```
type human struct {
	Sex int
}

type teacher struct {
	human
	name string
	age int
}

type student struct {
	human
	name string
	age int
}

func main () {
	// a := teacher{name: "cyan", age: 20, human{sex: 0}}
	a := teacher{name: "cyan", age: 20, human: human{Sex: 0}}
	b := student{name: "pink", age: 22, human: human{Sex: 1}}

	a.name = "xixi"
	a.age = 23
	// a.Sex = 100
	a.human.Sex = 200

	fmt.Println(a, b)
}
```


# 方法method

- Go 中虽没有`class`，但依旧有`method`
- 通过显示说明`receiver`来实现与某个类型的组合
- 只能为同一个包中的类型定义方法
- `Receiver` 可以是类型的值或者指针
- 不存在方法重载
- 可以使用值或指针来调用方法，编译器会自动完成转换
- 从某种意义上来说，方法是函数的语法糖，因为receiver其实就是
- 方法所接收的第1个参数（Method Value vs. Method Expression）
- 如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
- 类型别名不会拥有底层类型所附带的方法
- 方法可以调用结构中的非公开字段

```
type Test struct {
	name string
}

type Person struct {
	name string
}

func main () {
	t := Test{}
	t.Print()
	fmt.Println(t.name)

	p := Person{}
	p.Print()
	fmt.Println(p.name)
}

func (t *Test) Print() {
	t.name = "red"
	fmt.Println("Test")
}

func (p Person) Print() {
	fmt.Println("Person")
}

```
----
```
// 类型别名不会拥有底层类型所附带的方法
type TZ int

func main () {
	var a TZ
	a.Print() 
	(*TZ).Print(&a)
}

func (a *TZ)  Print() {
	fmt.Println("TZ")
}
```
方法不同调用方式
```
type A struct {
	name string
}

func main () {
	a := A{}
	a.Print() 
	// (*TZ).Print(&a)
}

func (a *A)  Print() {
	a.name = "123"
	fmt.Println(a.name)
	// fmt.Println("TZ")
}
```
方法访问权限
`struct`中的私有属性，在方法中可以访问

```
// 属性的访问范围是在`package`中的可以访问的，如果需要在外部包中访问，需要大写字母
type A struct {
	name string
}

func main () {
	a := A{}
	a.Print()
	fmt.Println(a.name)
}

func (a *A)  Print() {
	a.name = "123"
	fmt.Println(a.name)
}
```

# 接口interface

- 接口是一个或多个方法签名的集合
- 只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显示声明了哪个接口，称之为：`Structural Typing`
- 接口只有方法声明，没有实现，没有数据字段
- 接口可以匿名嵌入其它接口，或嵌入到结构中
- 将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，即无法修改复制品的状态，也无法获取指针
- 只有当接口存储的类型和对象都为nil时，接口才等于nil
- 接口调用不会做`receiver`
- 接口同样支持匿名字段方法
- 接口有可以实现OOP中的多态
- 空接口可以作为任何类型数据的容器

```
type USB interface {
	Name() string
	Connect()
}

type Phone struct {
	name string
}

func (pc Phone) Name() string {
	return pc.name
}

func (pc Phone) Connect() {
	fmt.Println("Connect: ", pc.name)
}

func main () {
	var a USB
	a = Phone{name: "phone"}
	a.Connect()
	Disconnect(a)
}

func Disconnect(usb USB) {
	fmt.Println("Disconnect")
}
```

接口嵌套：
```
type USB interface {
	Name() string
	// Connect()
	Connecter
}

type Connecter interface {
	Connect()
}

type Phone struct {
	name string
}

func (pc Phone) Name() string {
	return pc.name
}

func (pc Phone) Connect() {
	fmt.Println("Connect: ", pc.name)
}

func main () {
	var a USB
	a = Phone{name: "phone"}
	a.Connect()
	Disconnect(a)
}

func Disconnect(usb USB) {
	fmt.Println("Disconnect")
}

```

空接口的使用：
```
// Go语言中所有类型都实现空接口
// type empty interface {
// }

type USB interface {
	Name() string
	// Connect()
	Connecter
}

type Connecter interface {
	Connect()
}

type Phone struct {
	name string
}

func (pc Phone) Name() string {
	return pc.name
}

func (pc Phone) Connect() {
	fmt.Println("Connect: ", pc.name)
}

func main () {
	var a USB
	a = Phone{name: "phone"}
	a.Connect()
	Disconnect(a)

    // 空接口的判断
    var b interface{}
    fmt.Println(b == nil) // true
}

func Disconnect(usb interface{}) { // interface{} 空接口
	// if pc,ok := usb.(Phone); ok { // 类型判断
	// 	fmt.Println("Disconnect：", pc.name)	
	// 	return
	// }
	switch v := usb.(type) {
		case Phone:
			fmt.Println("Disconnect：", v.name)
		default :
			fmt.Println("Unknown")
	}
	fmt.Println("Disconnect")
}
```

只有当接口存储的类型和对象都为nil时，接口才等于nil
```
func main () {
	// 空接口的判断
	var b interface{}
	fmt.Println(b == nil) // true

	var p *int = nil
	b = p
	fmt.Println(b == nil) // false
}
```

> 类型断言

通过类型断言的`ok pattern`可以判断接口中的数据类型
使用`type switch`则可针对空接口进行比较全面的类型判断

```
type USB interface {
	Name() string
	// Connect()
	Connecter
}

type Connecter interface {
	Connect()
}

func Disconnect(usb interface{}) { // interface{} 空接口
	// if pc,ok := usb.(Phone); ok { // 类型判断
	// 	fmt.Println("Disconnect：", pc.name)	
	// 	return
	// }
	switch v := usb.(type) {
		case Phone:
			fmt.Println("Disconnect：", v.name)
		default :
			fmt.Println("Unknown")
	}
	fmt.Println("Disconnect")
}

```

> 接口转换

可以将拥有超集的接口转换为子集的接口


# 反射reflection

- 反射课大大提高程序的灵活性，使得`interface{}`有更大的发挥余地
- 反射使用`TypeOf`和`ValueOf`函数从接口中获取目标对象信息
- 反射会将匿名字段作为独立字段(匿名字段本质)
- 想要利用反射修改对象状态，前提是`interface.data`是`settabel`即`pointer-interface`
- 通过反射可以"动态"调用方法


```
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (u User) Hello() {
	fmt.Println("Hello world")
}

func main () {
	u := User{1, "alogy", 12}
	Info(u)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type: ", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields: ")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println(f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
```
如果是地址引用通过`Kind()`来获取与`reflect.Struct`匹配的对象。

```
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (u User) Hello() {
	fmt.Println("Hello world")
}

func main () {
	u := User{1, "alogy", 12}
	Info(&u)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type: ", t.Name())

	fmt.Println(t.Kind())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("XX")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields: ")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println(f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}

}
```
匿名字段反射： 
```
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

type Manager struct {
	User
	title string
}

func main () {
	// 反射会将匿名字段当作独立字段来处理
	m := Manager{User: User{1, "OK", 12}, title: "123123"}
	t := reflect.TypeOf(m)
	fmt.Println(t.Field(0))

	// 取匿名当中的字段
	fmt.Println(t.FieldByIndex([]int{0, 0}))
}
```
指针操作：
```
package main

import (
	"fmt"
	"reflect"
)

func main () {
	x := 123
	v := reflect.ValueOf(&x)
	// fmt.Println(v)
	v.Elem().SetInt(999)
	fmt.Println(x)
}
```

类型判断修改字段：
 
```
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func main () {
	u := User{1, "Ok", 18}
	Set(&u)
	fmt.Println(u)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if reflect.Ptr == v.Kind() && !v.Elem().CanSet() { // 指针是否正确
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem() // 重写赋值
	}

	f := v.FieldByName("Name") // 获取字段
	if !f.IsValid() { // 判读是否取到当前字段
		fmt.Println("BAD")
		return
	}
	if f.Kind() == reflect.String { // 类型判断
		f.SetString("MM") // 重新赋值
	}

}
```

通过反射调用方法，动态调用方法：

```
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (u User) Hello(name string) {
	fmt.Println("Hello", name, ", my name is", u.Name)
}

func main () {
	u := User{1, "OK", 123}
	v := reflect.ValueOf(u)

	mv := v.MethodByName("Hello") // 获取函数名
	args := []reflect.Value{reflect.ValueOf("alogy")} // 传递参数

	mv.Call(args)

	u.Hello("alogy")
}

```

# 并发concurrency


并发主要由切换时间片来实现“同时”运行，在并行则是直接利用多核实现多线程的运行，但Go可以设置使用核数，以发挥多核计算机的能力。

Goroutine奉行通过**通信来共享内存**，而不是共享内存来通信。

```
package main

import (
	"fmt"
	"time"
)

func main () {
	fmt.Println(2 * time.Second) // 2s
	go Go()
	time.Sleep(2 * time.Second) // 延迟2s
	// 在main函数运行Sleep的时候，Go函数也运行了,执行完之后，退出。
}

func Go () {
	fmt.Println("Go...")
}
```


> Channel

- `Channel`是`goroutine`沟通的桥梁，大都是阻塞同步的.通过关键字`go`加函数的名称，来实现`goroutine`
- 通过`make`创建，`close`关闭
- `Channel`是引用类型
- 可以使用`for range`来迭代不断操作的`Channel`
- 可以设置单向或双向通道
- 可以设置缓存大小，在未被填满前不会发生阻塞(没有设置为0，就是阻塞的)

> Select

- 可处理一个或多个`channel`的发送与接收
- 同时有多个可用的`channel`时按随机顺序处理
- 可用空的`select`来阻塞`main`函数
- 可设置超时

channel简单使用：
```
package main

import (
	"fmt"
)

func main () {
	c := make(chan bool)
	go func () {
		fmt.Println("Go...")
		c <- true // 存 // 声明的时候是bool
	}()
	<-c // 取 // 消息存取，阻塞执行
}
```
-----

```
package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go...")
		c <- true
		close(c) // 关闭chan // 没有明确关闭，会出现死锁，崩溃退出
	}()
	// <- c
	for v := range c {
		fmt.Println(v)
	}
}
```


> 有无缓存区别

有缓存：异步，无缓存，同步。

```
package main

import (
	"fmt"
)

func main() {
	// c := make(chan bool, 1) // 有缓存 异步
	c := make(chan bool) // 无缓存的时候是阻塞的
	go func () {
		fmt.Println("Go...")	
		<- c // 读取
	}()
	c <- true // 传进去
}

```

> 多核分配任务

出现任务分配出现不一定的情况，解决方法(多个`goroutine`的打印)：

1. 设置缓存
2. 利用内置包`sync`

```
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //  使用多核分配的时候，任务分配时不一定的
	c := make(chan bool, 10) // 设置缓存
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}

	for i := 0; i < 10; i++ {
		<-c
	}
}

func Go(c chan bool, idx int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(idx, a)

	c <- true
}
```
-----
```
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //  使用多核分配的时候，任务分配时不一定的
	wg := sync.WaitGroup{} // 内置包sync使用
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}

	wg.Wait()
}

func Go(wg *sync.WaitGroup, idx int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(idx, a)

	wg.Done()
}

```

> 多个chan

通过`select语句`

```
package main

import (
	"fmt"
)

func main() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		for { // 通过死循环来不断发送和接收chan
			select {
			case v, ok := <-c1 :
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <- c2 :
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()
		
	c1 <- 1
	c2 <- "hello"
	c1 <- 2
	c2 <- "zf"

	close(c1)
	close(c2)

	for i := 0; i < 2; i++ {
		<-o
	}

}
```

> select作为发送者的应用

```
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for v := range c {
			fmt.Println(v)
		}	
	}()

	for i := 0; i < 10; i++ {
		select {
			case c <- 0:
			case c <- 1:	
		}
	}

    // select{} // 阻塞main函数退出，卡死main函数，场景经常使用在事件循环
}

```

> select超时设置

```
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	select {
		case v := <-c:
			fmt.Println(v)
		case t := <-time.After(3 * time.Second):
			fmt.Println(t)
			fmt.Println("Timeout")
	}
}
```
