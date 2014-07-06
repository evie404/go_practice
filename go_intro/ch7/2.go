package main

import "fmt"

func half(num int) (halfVal int, isEven bool) {
  return num / 2, num % 2 == 0
}

func main() {
  fmt.Println(half(1)) // 0 false
  fmt.Println(half(2)) // 1 true
  fmt.Println(half(3)) // 1 false
  fmt.Println(half(4)) // 2 true
}