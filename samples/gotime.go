package main

import (
	"fmt"
	"time"
)

//main will be called when an application start executing
func main() {
	fmt.Println("Hello Main !!!")
	for i := 0; i < 10; i++ {
		go count(i)

	}
	time.Sleep(time.Microsecond * 11000)
}

//Main ending here !!!

func count(id int) {

	for i := 0; i < 10; i++ {
		fmt.Println(id, ": ", i)
		time.Sleep(time.Microsecond * 1000)
	}

}
