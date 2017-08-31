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