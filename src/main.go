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
}

func main() {
  // main.go
  fmt.Println("Hello Go in Vs Code!")
  variables()
  strings()
  operators()
  mylib.PrintInfo()
  // fmt.Println(Add(1, 2), Sub(1, 2))
  println("Hello End Go in Vs Code!")
}


