package main

import (
  "fmt"
  "errors"
)

func main() {

  fmt.Println(f3())

  i, j := multiReturn()

  fmt.Println(i, j)

  sum, err := varArgSum(1,2,3,4,5,6)
  if(err != nil) {
    fmt.Println(err)
  }
  fmt.Println(sum)
  //fmt.Println(varArgSum(1,2,3,4,5,6))
  closure()

  nextEven := makeEvenGenerator()
  fmt.Println(nextEven()) // 0
  fmt.Println(nextEven()) // 2
  fmt.Println(nextEven()) // 4

  fmt.Println(factorial(10))

  defer func() {
    str := recover()
    fmt.Println(str)
  }()
  panic("initial yet to finish")
  defer func() {
    str := recover()
    fmt.Println(str)
  }()
  panic("Yet to finish up.")
}

func f1() int {
  return f2()
}

func f2() int {
  return 1
}

func f3() (r int) {
  r = f1()
  return r
}

func multiReturn () (int,int) {
  return 5,6
}

func varArgSum (numbers ...int) (int, error) {
  var sum int
  for i := range numbers {
    sum =+ sum + i
  }

  if(sum < 20) {
     return sum, errors.New("Less value found than expected")
  }
  return sum, nil
}

//closure is function inside function
func closure() {
  sum := func() int {
    return 10
  }
  fmt.Println(sum())
}

func makeEvenGenerator() func() uint {
     i := uint(0)
     return func() (ret uint) {
           ret = i
           i += 2
           return
      }
}


//recursion
func factorial(x int) int {
  if x <= 0 {
    return 1;
  }
  return x * factorial(x-1)
}
