package main

import "fmt"

func fib(i int) (int) {
  if i == 0 {
    return 1
  } else if i <= 2 {
    return i
  } else {
    return fib(i - 1) + fib (i - 2)
  }
}

func main() {
  for i := 0; i < 10; i ++ {
    fmt.Println(fib(i))
  }
}