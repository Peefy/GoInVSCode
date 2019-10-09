package main

import "fmt"
import "./package"

// show usage of the go variable
func variables()  {
  var age = 18
  num := 12
  sum := age + num
  var str = "123"
  fmt.Printf("%d %s %d\n", age, str, sum)
}

func main() {
  // main.go
  fmt.Println("Hello Go in Vs Code!")
  variables()
  mylib.PrintInfo()
  // fmt.Println(Add(1, 2), Sub(1, 2))
}


