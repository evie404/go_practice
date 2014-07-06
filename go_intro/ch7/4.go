package main

import "fmt"

func makeOddGenerator() func() uint {
  i := uint(1)
  return func () (val uint) {
    val = i
    i += 2
    return
  }
}

func main() {
  nextOdd := makeOddGenerator()
  nextOdd1 := makeOddGenerator()

  fmt.Println(nextOdd()) // 1
  fmt.Println(nextOdd1()) // 1
  fmt.Println(nextOdd()) // 3
  fmt.Println(nextOdd1()) // 3
  fmt.Println(nextOdd()) // 5
  fmt.Println(nextOdd()) // 7
  fmt.Println(nextOdd()) // 9
  fmt.Println(nextOdd1()) // 5
}