package main

import "fmt"

func smallest(numbers ...int) (smallest int) {
  smallest = numbers[0]
  for _, v := range numbers {
    if v < smallest {
      smallest = v
    }
  }
  return
}

func main() {
  fmt.Println(smallest(1,2,3)) // 1
  fmt.Println(smallest(3,2,3,1)) // 1
  fmt.Println(smallest(-3,2,3,1)) // -3
}