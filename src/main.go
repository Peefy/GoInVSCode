package main

import "fmt"
import "unsafe"
import "./package"

// show usage of the go variable
func variables()  {
  var age = 18
  num := 12
  sum := age + num
  var str = "123"
  fmt.Printf("%d %s %d\n", age, str, sum)
  var a int = 1
  fmt.Println(a)

  // 没有初始化就为零值
  var b int = 2
  a, b = b, a
  _, str = "", "123"
  fmt.Println(b)

  // bool 零值为 false
  var c bool
  fmt.Println(c)

  var i int
  var f float64
  var bb bool
  var s string
  fmt.Println(i, f, bb, s)

  var j, k int

  var (
    vname1 int = 1
    vname2 string = ""
  )

  fmt.Println(j, k, vname1, vname2)
  fmt.Println(multirfunc())

  const WIDTH = 12
  const LENGTH = 12

  fmt.Println(WIDTH * LENGTH)
  fmt.Println("sizeof:", unsafe.Sizeof(WIDTH))

  const (
    ca = 1 << iota
    cb = 3 << iota
    ck
    cj
  )
  fmt.Println(ca, cb, ck, cj)
}

func strings() {
  fmt.Printf("123" + "456" + "789" + "\n")
}

func multirfunc() (int, int, string) {
  return 1, 2, "abc"
}

func operators() {
  fmt.Println("go operators")
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

   num := 10
   if num < 20 {
    println("num < 20")
   } else {
    println("num >= 20")
   }

   var marks int = 90
   grade := "None"

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"  
   }
   println(grade)

}

func controls() {
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
     fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
  }  

  for i := range numbers {
    fmt.Printf("第 %d 位 x 的值 = %d\n", i, i)
  }  

  a = 10
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

}

func getSequence() func() int {
   i := 0
   return func() int {
      i += 1
      return i  
   }
}

type Circle struct {
   radius float64
}

func (c Circle) getArea() float64 {
   //c.radius 即为 Circle 类型对象中的属性
   return 3.14 * c.radius * c.radius
}

func functions() {
   maxfunc := func (num1, num2 int) int {
      var result int
      if num1 > num2 {
         result = num1
      } else {
         result = num2
      }
      return result
   }

   swap := func (x, y string) (string, string) {
      return y, x //函数可以返回多个值
   }

   println(maxfunc(10, 20))
   a, b := swap("DuGu", "GuDu")
   println(a, b)
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

   var c Circle
   c.radius = 10
   println("area of c is", c.getArea())

}

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

   /* 数组 - 5 行 2 列*/
   var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}

   /* 输出数组元素 */
   for  i = 0; i < 5; i++ {
      for j = 0; j < 2; j++ {
         fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
      }
   }

   println("")
}

func pointers() {
   var aptr int = 10  
   fmt.Println("变量的地址: %x\n", &aptr)
   var ptrnil *int
   if ptrnil == nil {
      fmt.Println("ptr 的值为 : %x\n", ptrnil)
   }

   const MAX int = 3
   a := []int{10,100,200}
   var i int
   var ptr [MAX]*int;

   for i = 0; i < MAX; i++ {
      ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
   }

   for i = 0; i < MAX; i++ {
      fmt.Printf("a[%d] = %d\n", i,*ptr[i] )
   }

}

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func (b *Books) GetTitle() string {    
   return b.title    
}

func structs() {
   // 创建一个新的结构体
   fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

   // 也可以使用 key => value 格式
   fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

   // 忽略的字段为 0 或 空
   fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

   var Book1 Books
   Book1.title = "dugu"
   println("book.title is", Book1.title)
   println("book.title is", Book1.GetTitle())
}

func slices() {
   /* 打印切片函数 */
   printSlice := func (x []int) {
      fmt.Println("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
   }
   /* 创建切片 */
   numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
   printSlice(numbers)
   fmt.Println("numbers ==", numbers)
   fmt.Println("")
   /* 打印子切片从索引1(包含) 到索引4(不包含)*/
   fmt.Println("numbers[1:4] ==", numbers[1:4])

   /* 默认下限为 0*/
   fmt.Println("numbers[:3] ==", numbers[:3])

   /* 默认上限为 len(s)*/
   fmt.Println("numbers[4:] ==", numbers[4:])

   numbers1 := make([]int,0,5)
   printSlice(numbers1)

   /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
   number2 := numbers[:2]
   printSlice(number2)

   /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
   number3 := numbers[2:5]
   printSlice(number3)


   /* 向切片添加一个元素 */
   numbers = append(numbers, 1)
   /* 同时添加多个元素 */
   numbers = append(numbers, 2,3,4)
   /* 创建切片 numbers1 是之前切片的两倍容量*/
   numbers4 := make([]int, len(numbers), (cap(numbers))*2)

   /* 拷贝 numbers 的内容到 numbers1 */
   copy(numbers4, numbers)

}

func ranges() {
   nums := []int{1, 2, 3}
   sum := 0
   for _, num := range nums {
      sum += num
   }
   println("the sum is", sum)
   //在数组上使用range将传入index和值两个变量。
   //上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
   for i, num := range nums {
      if num == 3 {
         println("the element 3 index:", i)
      }
   }
   //range可以用在map的键值对上。
   kvs := map[string]string{"a" : "apple", "d" : "dugu"}
   for k, v := range kvs {
      println("%s -> %s", k, v)
   }
   //range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
   for i, c := range "DuGu" {
      println(i, c)
   }
}

func main() {
  // main.go
  fmt.Println("Hello Go in Vs Code!")
  // Go变量
  variables()
  // Go字符串
  strings()
  // Go操作符
  operators()
  // Go运算符
  controls() 
  // Go函数
  functions()
  // Go数组
  arrays()
  // Go指针
  pointers()
  // Go结构体
  structs()
  // Go切片
  slices()
  // Go range 关键字
  ranges()
  // Go包导入
  mylib.PrintInfo()
  println("Hello End Go in Vs Code!")
}


