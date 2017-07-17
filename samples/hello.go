package main

import (
  "fmt"
  sutil "samples/stringutil"
)

func main() {
  fmt.Println("Hello, Go Language.")
  fmt.Println(reverse("Hello, Go Language."))
  fmt.Println(sutil.Reverse("Hello, Go Language."))

  var int1 int = 40
  var float1 float64 = 42.99

  fmt.Println(int1 + int(float1))

  fmt.Println(float64(int1) + float1)

  float2 := float64(int1) + float1

  fmt.Printf("is %v, %T \n" , int1, float2)
}

func reverse(input string) string {
  length := len(input)
  var output string

  for i := length-1; i >= 0; i-- {
    output = output + string(input[i])
  }

  fmt.Println(output)
  return output;
}
