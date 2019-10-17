
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
* *import "fmt"*告诉Go编译器需要使用fmt包，fmt包实现了格式化IO的函数
* *func main()*是程序开始执行的函数。main函数是每一个可执行程序所必须包含的，注意：*如果有init()函数则会先执行该函数*
* */\* \*/*是注释 或者以//开头
* *fmt.Println(...)*可以将字符串输出到控制台，并在最后添加换行符\n
* 当标识符(包括常量、变量、类型、函数名、结构字段等等)以一个大写字母开头，如Group,那么使用这种形式的标识符的对象就可以被外部包的代码所使用(类似public);标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的(像面向对象语言中的 protected)。

注意：*左括号{不能单独放在一行*


### Go语言语法

**Go标记**-可以由多个标记组成，可以是关键字，标识符，常量，字符串，符号。

**行分隔符**

在Go程序中，一行代表一个语句结束，每个语句不需要以分号结尾

**Go注释**

注释不会被编译，每一个包应该有相关注释。
单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾。如：

**标识符**

标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。

以下是有效的标识符

```go

mahesh   kumar   abc   move_name   a_123
myname50   _temp   j   a23b9   retVal

```

无效的标识符：以数字开头；Go语言的关键字；含有运算符的标识符

**字符串链接**

Go 语言的字符串可以通过 **+** 实现：

```go

package main
import "fmt"
func main() {
    fmt.Println("Google" + "Runoob")
}

```

**Go关键字**

Go|Go|关|键|字
-|-|-|-|-
beaak|default|func|interface|select
case |	defer|	go|	map|	struct
chan |	else|	goto|	package|	switch
const |	fallthrough|	if|	range|	type
continue |	for|	import|	return|	var

**Go预定义标识符**

||||||||
-|-|-|-|-|-|-|-|-
append |	bool	|byte	|cap |	close|	complex|	complex64|	complex128|	uint16
copy   |	false	|float32	|float64|	imag|	int|	int8|	int16|	uint32
int32  |	int64	|iota	|len|	make|	new|	nil|	panic|	uint64|
print  |	println	|real	|recover|	string|	true|	uint|	uint8|	uintptr

**Go程序组成**

程序一般由关键字、常量、变量、运算符、类型和函数组成。

程序中可能会使用用到这些分隔符：括号(),中括号[]和大括号{}

程度中可能会使用到这些标点符号: . , ; : 和 ...

**Go语言的空格**

Go 语言中变量的声明必须使用空格隔开，如：

```go
var age int
```

**Go语言数据类型**

在Go编程语言中，数据类型用于声明函数和变量

数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才需要申请大内存，就可以充分利用内存

Go语言按类别有以下几种数据类型：

序号|类型和描述
-|-
1|**布尔型**布尔型的值可以是常量true或者false.一个简单的例子: var b bool = true
2|**数字类型**整型int和浮点型float32、float64，Go语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码
3|**字符串类型**字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。Go语言的字符串的字节使用UTF-8编码标识Unicode文本
4|**派生类型** 包括 指针类型 数组类型 结构化类型 Channel类型 函数类型 切片类型 接口类型 Map类型

**Go数字类型**

Go也有基于架构的类型，例如int,uint,uintptr

**整型**

序号|类型和描述
-|-
1|**uint8**-无符号 8 位整型 (0 到 255)
2|**uint16**-无符号 16 位整型 (0 到 65535)
3|**uint32**-无符号 32 位整型 (0 到 4294967295)
4|**uint64**-无符号 64 位整型 (0 到 18446744073709551615)
5|**int8**-有符号 8 位整型 (-128 到 127)
6|**int16**-有符号 16 位整型 (-32768 到 32767)
7|**int32**-有符号 32 位整型 (-2147483648 到 2147483647)
8|**int64**-有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)

**浮点型**

序号|类型和描述
-|-
1|**float32**-IEEE-754 32位浮点型数
2|**float64**-IEEE-754 64位浮点型数
3|**complex64**-32 位实数和虚数
4|**complex128**-64 位实数和虚数

**其他数字类型**

序号|类型和描述
-|-
1|**byte**-类似 uint8
2|**rune**-类似 int32
3|**uint**-32 或 64 位
4|**int**-与 uint 一样大小
5|**uintptr**-无符号整型，用于存放一个指针


**Go语言变量**

Go语言变量名由字母、数字、下划线组成，其中首个字符不能为数字

```go
var identifier type;
```

可以一次声明多个变量

```go
var identifier1, identifier2 type;
```

```go
package main
import "fmt"
func main() {
    var a string = "DuGu"
    fmt.Println(a)
    var b, c int = 1, 2
    fmt.Println(b, c)
}
```

**变量声明**

* 指定变量类型，如果没有初始化，则变量默认为零值
* 根据值自行判定变量类型
* 省略**var**,使用:=运算符，如可以将var f string = "Runoob" 简写为 f := "Runoob"

```go
var v_name = v_type
v_name = value
```

```go
package main
import "fmt"
func main() {
    var a = "DuGu"
    fmt.Println(a)

    var b int
    fmt.Println(b)

    var c bool
    fmt.Println(c)
}
```

* 数值类型(包括complex64/128)为0
* 布尔类型为false
* 字符串为""(空字符串)
* 一下几种类型为nil

```go
var a *int
var a []int
var a map[string] int
var c chan int
var a func(string) int
var a error // error 是接口
```

```go
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)
```

*多变量声明*-

```go
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误,并且这种不带格式声明的只能在函数体中出现

// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)
```

**Go变量值类型和引用类型**

所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值

当使用等号 = 将一个变量的值赋值给另一个变量时间，j = i, 实际是在内存中将i的值进行了拷贝。

可以通过&i来获取变量i的内存地址，例如0xf840000040

*注意:*如果在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明,若果在定义变量a之前使用它，则会得到编译错误undefined: a;如果声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误

**Go语言常量**

常量是一个简单值的标识符，在程序运行时，不会被修改的量。

常量的定义格式：const identifier \[type\] = value 

可以省略类型说明符，因为编译器可以根据变量的值来推断

```go
package main

import "fmt"

func main() {
   const LENGTH int = 10
   const WIDTH int = 5  
   var area int
   const a, b, c = 1, false, "str" //多重赋值

   area = LENGTH * WIDTH
   fmt.Printf("面积为 : %d", area)
   println()
   println(a, b, c)  
}
```


*iota*-特殊常量，可以认为是一个可以被编译器修改的常量,iota在const关键字出现时将被重置为0，const中每新增一行常量声明将使iota计数一次(iota可理解为cons语句块中的行索引)

```go
package main

import "fmt"

func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
}
```

运行结果

> 0 1 2 ha ha 100 100 7 8

**Go语言运算符**

Go语言内置的运算符

* **算术运算符**

A = 10, B = 20

运算符|描述|实例
-|-|-
+|	相加|	A + B 输出结果 30
-|	相减|	A - B 输出结果 -10
*|	相乘|	A * B 输出结果 200
/|	相除|	B / A 输出结果 2
%|	求余|	B % A 输出结果 0
++|	自增|	A++ 输出结果 11
--|	自减|	A-- 输出结果 9

```go
package main

import "fmt"

func main() {

   var a int = 21
   var b int = 10
   var c int

   c = a + b
   fmt.Printf("第一行 - c 的值为 %d\n", c )
   c = a - b
   fmt.Printf("第二行 - c 的值为 %d\n", c )
   c = a * b
   fmt.Printf("第三行 - c 的值为 %d\n", c )
   c = a / b
   fmt.Printf("第四行 - c 的值为 %d\n", c )
   c = a % b
   fmt.Printf("第五行 - c 的值为 %d\n", c )
   a++
   fmt.Printf("第六行 - a 的值为 %d\n", a )
   a=21   // 为了方便测试，a 这里重新赋值为 21
   a--
   fmt.Printf("第七行 - a 的值为 %d\n", a )
}
```

* **关系运算符**

A = 10, B = 20

运算符|描述|实例
-|-|-
== |	检查两个值是否相等，如果相等返回 True 否则返回 False。	|(A == B) 为 False
!= |	 检查两个值是否不相等，如果不相等返回 True 否则返回 False。|	(A != B) 为 True
\>  |	检查左边值是否大于右边值，如果是返回 True 否则返回 False。|	(A > B) 为 False
\<	| 检查左边值是否小于右边值，如果是返回 True 否则返回 False。 |	(A < B) 为 True
\>=	| 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。|	(A >= B) 为 False
\<=	| 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。|	(A <= B) 为 True

```go
package main

import "fmt"

func main() {
   var a int = 21
   var b int = 10

   if( a == b ) {
      fmt.Printf("第一行 - a 等于 b\n" )
   } else {
      fmt.Printf("第一行 - a 不等于 b\n" )
   }
   if ( a < b ) {
      fmt.Printf("第二行 - a 小于 b\n" )
   } else {
      fmt.Printf("第二行 - a 不小于 b\n" )
   }
   
   if ( a > b ) {
      fmt.Printf("第三行 - a 大于 b\n" )
   } else {
      fmt.Printf("第三行 - a 不大于 b\n" )
   }
   /* Lets change value of a and b */
   a = 5
   b = 20
   if ( a <= b ) {
      fmt.Printf("第四行 - a 小于等于 b\n" )
   }
   if ( b >= a ) {
      fmt.Printf("第五行 - b 大于等于 a\n" )
   }
}
```

* **逻辑运算符**

A = True, B = False

运算符|描述|实例
-|-|-
&& |	逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。	(A && B) 为 False
\|\| |	逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。	(A \|\| B) 为 True
! |	逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。	!(A && B) 为 True

```go
package main

import "fmt"

func main() {
   var a bool = true
   var b bool = false
   if ( a && b ) {
      fmt.Printf("第一行 - 条件为 true\n" )
   }
   if ( a || b ) {
      fmt.Printf("第二行 - 条件为 true\n" )
   }
   /* 修改 a 和 b 的值 */
   a = false
   b = true
   if ( a && b ) {
      fmt.Printf("第三行 - 条件为 true\n" )
   } else {
      fmt.Printf("第三行 - 条件为 false\n" )
   }
   if ( !(a && b) ) {
      fmt.Printf("第四行 - 条件为 true\n" )
   }
}
```

* **位运算符**

p |	q |	p & q |	p \| q | p ^ q
-|-|-|-|-
0 |	0 |	0 |	0 |	0
0 |	1 |	0 |	1 |	1
1 |	1 |	1 |	1 |	0
1 |	0 |	0 |	1 |	1

```go
A = 0011 1100

B = 0000 1101

-----------------

A & B = 0000 1100

A | B = 0011 1101

A ^ B = 0011 0001
```

运算符|描述|实例
-|-|-
& |	按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。	(A & B) 结果为 12, 二进制为 0000 1100
\| |	按位或运算符"\|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或 |	(A \| B) 结果为 61, 二进制为 0011 1101
^ |	按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。 |	(A ^ B) 结果为 49, 二进制为 0011 0001
<< |	左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。|	A << 2 结果为 240 ，二进制为 1111 0000
\>\>	右移运算符"\>\>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，"\>\>"右边的数指定移动的位数。	A >> 2 结果为 15 ，二进制为 0000 1111

```go
package main

import "fmt"

func main() {

   var a uint = 60      /* 60 = 0011 1100 */  
   var b uint = 13      /* 13 = 0000 1101 */
   var c uint = 0          

   c = a & b       /* 12 = 0000 1100 */
   fmt.Printf("第一行 - c 的值为 %d\n", c )

   c = a | b       /* 61 = 0011 1101 */
   fmt.Printf("第二行 - c 的值为 %d\n", c )

   c = a ^ b       /* 49 = 0011 0001 */
   fmt.Printf("第三行 - c 的值为 %d\n", c )

   c = a << 2     /* 240 = 1111 0000 */
   fmt.Printf("第四行 - c 的值为 %d\n", c )

   c = a >> 2     /* 15 = 0000 1111 */
   fmt.Printf("第五行 - c 的值为 %d\n", c )
}
```

* **赋值运算符**

运算符|描述|实例
-|-|-
= |	简单的赋值运算符，将一个表达式的值赋给一个左值 | C = A + B 将 A + B 表达式结果赋值给 C
+= |	相加后再赋值 | C += A 等于 C = C + A
-= |	相减后再赋值 | C -= A 等于 C = C - A
*= |	相乘后再赋值 | C *= A 等于 C = C * A
/= |	相除后再赋值 | C /= A 等于 C = C / A
%= |	求余后再赋值 | C %= A 等于 C = C % A
<<=	 | 左移后赋值 |	C <<= 2 等于 C = C << 2
\>\>= |	右移后赋值 | C >>= 2 等于 C = C >> 2
&= |	按位与后赋值 | C &= 2 等于 C = C & 2
^= |	按位异或后赋值 | C ^= 2 等于 C = C ^ 2
|= |	按位或后赋值 |	C |= 2 等于 C = C | 2

```go
package main

import "fmt"

func main() {
   var a int = 21
   var c int

   c =  a
   fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c )
   c +=  a
   fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c )
   c -=  a
   fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c )
   c *=  a
   fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c )
   c /=  a
   fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c )
   c  = 200;

   c <<=  2
   fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c )
   c >>=  2
   fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c )
   c &=  2
   fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c )
   c ^=  2
   fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c )
   c |=  2
   fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c )
}
```

* **其他运算符**

运算符|描述|实例
-|-|-
& |	返回变量存储地址 |	\&a; 将给出变量的实际地址。
\* |	指针变量。 |	\*a; 是一个指针变量

```go
package main

import "fmt"

func main() {
   var a int = 4
   var b int32
   var c float32
   var ptr *int

   /* 运算符实例 */
   fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
   fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
   fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );

   /*  & 和 * 运算符实例 */
   ptr = &a     /* 'ptr' 包含了 'a' 变量的地址 */
   fmt.Printf("a 的值为  %d\n", a);
   fmt.Printf("*ptr 为 %d\n", *ptr);
}
```

**Go运算符优先级**

优先级|运算符
-|-
5|	\* / % << \>\> & &^
4|	+ - \| ^
3|	== != < <= \> \>=
2|	&&  逻辑与
1|	\| \| 逻辑或


```go

package main

import "fmt"

func main() {
   var a int = 20
   var b int = 10
   var c int = 15
   var d int = 5
   var e int;

   e = (a + b) * c / d;      // ( 30 * 15 ) / 5
   fmt.Printf("(a + b) * c / d 的值为 : %d\n",  e );

   e = ((a + b) * c) / d;    // (30 * 15 ) / 5
   fmt.Printf("((a + b) * c) / d 的值为  : %d\n" ,  e );

   e = (a + b) * (c / d);   // (30) * (15/5)
   fmt.Printf("(a + b) * (c / d) 的值为  : %d\n",  e );

   e = a + (b * c) / d;     //  20 + (150/5)
   fmt.Printf("a + (b * c) / d 的值为  : %d\n" ,  e );  
}

```

指针变量 **\*** 和地址值 **&** 的区别：指针变量保存的是一个地址值，会分配独立的内存来存储一个整型数字。当变量前面有 **\*** 标识时，才等同于 **&** 的用法，否则会直接输出一个整型数字。

```go
func main() {
    var a int = 4
    var ptr *int
    ptr = &a
    println("a的值为", a); // 4
    println("*ptr为", *ptr); // 4
    println("ptr为", ptr)  // 824633794744
}
```

*注意：Go的自增只能作为表达式使用，而不能用于赋值语句*

**Go条件语句**

语句|描述
-|-
if | 语句	if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
if...else | 语句	if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
if 嵌套语句 |	可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
switch 语句	 | switch 语句用于基于不同条件执行不同动作。
select 语句	| select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。

```go

   num := 10
   if num < 20 {
    println("num < 20")
   } else {
    println("num >= 20")
   }

   var marks int = 90

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"  
   }

```

**Go语言selec语句**

select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。

select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。

*select语法*

```go
select {
    case communication clause  :
       statement(s);      
    case communication clause  :
       statement(s); 
    /* 你可以定义任意数量的 case */
    default : /* 可选 */
       statement(s);
}
```

* 每个 case 都必须是一个通信
* 所有 channel 表达式都会被求值
* 所有被发送的表达式都会被求值
* 如果任意某个通信可以进行，它就执行，其他被忽略。
* 如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。否则
1. 如果有 default 子句，则执行该语句。
2. 如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。

```go
package main

import "fmt"

func main() {
   var c1, c2, c3 chan int
   var i1, i2 int
   select {
      case i1 = <-c1:
         fmt.Printf("received ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
      default:
         fmt.Printf("no communication\n")
   }    
}
```

*注意：Go没有三目运算符?:*

**Go语言循环语句**

*Go的for循环*
```go
package main

import "fmt"

func main() {

   var b int = 15
   var a int

   numbers := [6]int{1, 2, 3, 5}

   /* for 循环 */
   for a := 0; a < 10; a++ {
      fmt.Printf("a 的值为: %d\n", a)
   }

   for a < b {
      a++
      fmt.Printf("a 的值为: %d\n", a)
   }

   for i,x := range numbers {
      fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
   }  
}
```

**Go循环控制语句**

控制语句|描述
-|-
break语句 | 经常用于中断当前 for 循环或跳出 switch 语句
continue 语句 | 跳过当前循环的剩余语句，然后继续进行下一轮循环。
goto语句 | 将控制转移到被标记的语句。

```go
var a int = 10
for a < 20 {
   println("a 的值为：", a)
   a++
   if a == 12 {
      println("a continue")
      continue
   }
   if a > 15 {
      println("a break")
      break
   }
}

a = 10
LOOP: for a < 20 {
   if a == 15 {
      /* 跳过迭代 */
      a = a + 1
      goto LOOP
   }
   fmt.Printf("a的值为 : %d\n", a)
   a++    
}
```

**Go语言函数**

函数是基本的代码块，用于执行一个任务，Go语言最少有个main()函数，可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。

函数声明高速了编译器函数的名称，返回类型和参数。

Go语言标准库提供了多种可动用的内置函数。例如，len()函数可以接受不同类型参数并返回该类型的长度。如果传入字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含元素的个数。

Go语言函数定义格式如下：

```go
func function_name( [paramater list] ) [return types] {
   // 函数体
}
```

函数定义解析：
* func：函数由 func 开始声明
* function_name：函数名称，函数名和参数列表一起构成了函数签名。
* parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
* return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
* 函数体：函数定义的代码集合。

```go
func max(num1, num2 int) int {
   var result int
   if num1 > num2 {
      result = num1
   } else {
      result = num2
   }
   return result
}
```

*Go函数可以返回多个值*

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x //函数可以返回多个值
}

func main() {
   a, b := swap("DuGu", "GuDu")
   fmt.Println(a, b)
}

```

**Go函数参数**

函数如果使用参数，该变量可称为函数的形参。

形参就像定义在函数体内的局部变量。

调用函数，可以通过两种方式来传递函数

传递类型|描述
-|-
值传递|值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
引用传递|引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

*注意：默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。*

**Go函数变量**

Go 语言可以很灵活的创建函数，并作为另外一个函数的实参。

```go
package main

import (
   "fmt"
   "math"
)

func main() {
   getSquareRoot := func(x float64) float64 {
      return math.Sqrt(x)
   }  
   println(getSquareRoot(9))
}
```

**Go匿名函数和闭包**

Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。

```go
package main

import "fmt"

func getSequence() func() int {
   i := 0
   return func() int {
      i += 1
      return i  
   }
}

func main(){
   /* nextNumber 为一个函数，函数 i 为 0 */
   nextNumber := getSequence()  

   /* 调用 nextNumber 函数，i 变量自增 1 并返回 */
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   
   /* 创建新的函数 nextNumber1，并查看结果 */
   nextNumber1 := getSequence()  
   fmt.Println(nextNumber1())
   fmt.Println(nextNumber1())
}
```

**Go语言函数方法**

Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下：

```go
package main

import (
   "fmt"  
)

/* 定义结构体 */
type Circle struct {
  radius float64
}

func main() {
  var c1 Circle
  c1.radius = 10.00
  fmt.Println("圆的面积 = ", c1.getArea())
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
  //c.radius 即为 Circle 类型对象中的属性
  return 3.14 * c.radius * c.radius
}
```

**Go语言变量作用域**

作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。

Go语言中变量可以在三个地方声明：

* 函数内定义的变量称为局部变量
* 函数外定义的变量称为全局变量
* 函数定义中的变量称为形式参数

**局部变量**

```go
package main

import "fmt"

func main() {
   var a, b, c int

   a = 10
   b = 20
   c = a + b

   println("a = %d, b = %d, c = %d", a, b, c)
}
```

**全局变量**

Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。实例如下：

```go

package main

import "fmt"

func main() {
   var g int = 10
   println("g = %d\n", g)
}

```

**形式参数**

形式参数会作为函数的局部变量来使用。实例如下：

```go

package main

import "fmt"

/* 声明全局变量 */
var a int = 20;

func main() {
   /* main 函数中声明局部变量 */
   var a int = 10
   var b int = 20
   var c int = 0

   fmt.Printf("main()函数中 a = %d\n",  a);
   c = sum( a, b);
   fmt.Printf("main()函数中 c = %d\n",  c);
}

/* 函数定义-两数相加 */
func sum(a, b int) int {
   fmt.Printf("sum() 函数中 a = %d\n",  a);
   fmt.Printf("sum() 函数中 b = %d\n",  b);

   return a + b;
}

```

**初始化局部和全局变量**

不同类型的局部和全局变量默认值为：

数据类型|初始化默认值
-|-
int|0
float32|0
pointer|nil


**Go语言数组**

Go语言提供了数组类型的数据结构

数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整型、字符串或者自定义类型。

相对于去声明number0, number1, number2, ..., number99的变量，更加方便且易于扩展。

**声明数组**

Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：

```go
var variable_name [SIZE] variable_type
```

以上为一维数组的定义方式。例如以下定义了数组 balance 长度为 10 类型为 float32：

```go
var balance [10] float32
```

**初始化数组**

```go
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

初始化数组中 \{\} 中的元素个数不能大于 \[\] 中的数字。

如果忽略 \[\] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：

```go
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

**访问数组元素**

数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。例如：

```go
var salary float32 = balance[9]
```

```go
func arrays() {
   var n [10]int /* n 是一个长度为 10 的数组 */
   var i, j int

   /* 为数组 n 初始化元素 */        
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* 设置元素为 i + 100 */
   }

   /* 输出每个数组元素的值 */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d ", j, n[j] )
   }
   println("")
}
```

**二维数组**

```go
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
```

```go
var threedim [5][10][4]int
```

**二维数组**

二维数组是最简单的多维数组，二维数组本质上是由一维数组组成的。二维数组定义方式如下

```go
var arrayName [x][y]variable_type
```

**初始化二维数组**

多维数组可通过大括号来初始值。

```go
a = [3][4]int {
   {0, 1, 2, 3}, 
   {4, 5, 6, 7}, 
   {8, 9, 10, 11}, 
}
```

```go
package main

import "fmt"

func main() {
   /* 数组 - 5 行 2 列*/
   var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
   var i, j int

   /* 输出数组元素 */
   for  i = 0; i < 5; i++ {
      for j = 0; j < 2; j++ {
         fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
      }
   }
}
```

**Go语言指针**

Go语言的取地址符是&，放到一个变量前使用就会返回相应变量的内存地址。

```go
package main

import "fmt"

func main() {
   var a int = 10  
   fmt.Printf("变量的地址: %x\n", &a)
}
```
