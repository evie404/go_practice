package main

import "fmt"

func swap(a *int, b *int) {
  var tmp = *a
  *a = *b
  *b = tmp
}

func main() {
  x := 1
  y := 2
  fmt.Println(x, y)
  swap(&x, &y)
  fmt.Println(x, y) // I think this only swaps the values and not the pointer reference?
}