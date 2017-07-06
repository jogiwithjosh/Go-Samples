package main

import (
  "fmt"
  sutil "samples/stringutil"
)

func main() {
  fmt.Println("Hello, Go Language.")
  fmt.Println(reverse("Hello, Go Language."))
  fmt.Println(sutil.Reverse("Hello, Go Language."))
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
