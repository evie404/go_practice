package main

import "fmt"

var slice1 = []int{1,2,3}
var slice2 = []int{4,5,6}

func sum(ary []int) (total int) {
  total = 0
  for _, value := range ary {
    total += value
  }
  return
}

func main() {
  fmt.Println(sum(slice1)) // 6
  fmt.Println(sum(slice2)) // 15
}