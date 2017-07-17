package main

import (
  "fmt"
)

func main() {

  c := Circle{1,2,3}
  fmt.Println(area(c))
}

type Circle struct {
   x int
   y int
   z int
}

func area(c Circle) int {
  return c.x * c.y * c.z
}
