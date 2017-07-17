package main

import (
	"fmt"
	"strconv"
	"time"
)

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChain chan string) {

	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
	fmt.Println("Make Dough and Send for Sauce")
	stringChain <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChain chan string) {

	pizza := <-stringChain

	fmt.Println("Add Sauce and Send ", pizza, "for toppings")
	stringChain <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addToppings(stringChain chan string) {

	pizza := <-stringChain

	fmt.Println("Add Toppings to", pizza, "and ship ")
	stringChain <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

//main will be called when an application start executing
func main() {
	fmt.Println("Hello World !!!")

	stringChain := make(chan string)

	for i := 0; i < 3; i++ {
		go makeDough(stringChain)
		go addSauce(stringChain)
		go addToppings(stringChain)

		time.Sleep(time.Millisecond * 5000)
	}

}

//Main ending here !!!
