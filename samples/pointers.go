package main

import (
  "fmt"
)

func main() {
  x := 5
  zero(&x)

  fmt.Println(x)

  y := new(int)
  *y = 20
  zero(y)

  fmt.Println(y)
  fmt.Println(*y)

  z := 1.5
  fmt.Println(square(&z))
  fmt.Println(square(&z))
}

func zero(number *int) {
  *number = 0
}

func square(x *float64) float64 {
    *x = *x * *x
    return *x
}
