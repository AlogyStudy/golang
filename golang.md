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