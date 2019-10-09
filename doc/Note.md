
## Go笔记

### Go简介

**Go**（又称 **Golang**）是 Google 的 Robert Griesemer，Rob Pike 及 Ken Thompson 开发的一种*静态强类型*、*编译型语言*。Go 语言语法与 C 相近，但功能上有：**内存安全**，**GC（垃圾回收）**，**结构形态**及 **CSP-style 并发计算**。

与C++相比，Go并不包括如枚举、异常处理、继承、泛型、断言、虚函数等功能，但增加了 切片(Slice) 型、并发、管道、垃圾回收、接口（Interface）等特性的语言级支持。Go 2.0版本将支持泛型，对于断言的存在，则持负面态度，同时也为自己不提供类型继承来辩护。

### Go语言基础

Go语言的基础组成由以下几个部分

* **包声明**
* **引入包**
* **函数**
* **变量**
* **语句 & 表达式**
* **注释**

### 第一个Go程序

```go
package main

import "fmt"

func main() {
    /* This is a go demo code. */
    fmt.Println("Hello World!")
}

```

* 第一行代码*package main*定义了包名，必须在源文件中非注释的第一行指明这个文件属于哪一个包
* 下一行*import "fmt"*告诉Go编译器需要使用fmt包，fmt包实现了格式化IO的函数
* 
* 
* 
* 
* 
