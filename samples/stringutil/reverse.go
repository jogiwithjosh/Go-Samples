package stringutil

import (
  "fmt"
)

func Reverse(input string) string {
  length := len(input)
  var output string

  for i := length-1; i >= 0; i-- {
    output = output + string(input[i])
  }

  fmt.Println(output)
  return output;
}
