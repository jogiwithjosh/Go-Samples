package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

//main will be called when an application start executing
func main() {
	fmt.Println("Hello World !!!")

	var age int = 50               // int variable
	var favNum float64 = 1.1680339 // float variable

	// randNum := 1 // automatically creates int type variable

	fmt.Println(age, " ", favNum) // or same fmt.Println(age, favNum)

	var numOne = 1.000
	var num99 = .999

	fmt.Println("Substraction --> ", numOne-num99)

	fmt.Println("6 + 4 = ", 6+4)
	fmt.Println("6 - 4 = ", 6-4)
	fmt.Println("6 * 4 = ", 6*4)
	fmt.Println("6 / 4 = ", 6/4)
	fmt.Println("6 % 4 = ", 6%4)

	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)

	//const pi float64 = 3.1452456 // declaring connstant
	var myName string = "Vinod Paliwal"
	var isOver30 bool = true

	fmt.Println(len(myName))

	fmt.Println("I like \n \n")
	fmt.Println("Newlines")
	fmt.Println("Boolean => ", isOver30)

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("for loop ==> ")
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	yourAge := 18 // yourAge becomes int as assigning int to it

	if yourAge >= 16 {
		fmt.Println("You can Drive")
	} else if yourAge >= 18 {
		fmt.Println("You can Vote")
	} else {
		fmt.Println("You can have fun")
	}

	switch yourAge {
	case 16:
		fmt.Println("Go Drive!!")
	case 18:
		fmt.Println("Go Vote!!")
	default:
		fmt.Println("Go Have Fun!!")
	}

	//Array
	var favNums2 [5]float64

	favNums2[0] = 163
	favNums2[1] = 7449
	favNums2[2] = 844
	favNums2[3] = 3.22
	favNums2[4] = 1.383
	fmt.Println("Array => ", favNums2[3])

	favNum3 := [5]float64{1, 2, 3, 4, 5}

	for i, value := range favNum3 {
		fmt.Println("Index=> ", i, "Value =>", value)
	}

	//If Index not required then
	for _, value := range favNum3 {
		fmt.Println("Value =>", value)
	}

	numSlice := []int{5, 4, 3, 2, 1}
	numSlice2 := numSlice[3:5]
	fmt.Println("numSlice2 =", numSlice2)       //numSlice2 = [2 1]
	fmt.Println("numSlice2[1] =", numSlice2[1]) //numSlice2[1] = 1
	fmt.Println("numSlice[:2] =", numSlice[:2]) //numSlice[:2] = [5 4]
	fmt.Println("numSlice[2:] =", numSlice[2:]) //numSlice[2:] = [3 2 1]

	numSlice3 := make([]int, 5, 10) // Creting int type array, which would have first 5 values, Max size of array would be 10.
	copy(numSlice3, numSlice)
	fmt.Println("numSlice3 => ", numSlice3)
	fmt.Println("numSlice3[0] => ", numSlice3[0])

	numSlice3 = append(numSlice3, 0, -1, 9)
	fmt.Println("numSlice3 after append=> ", numSlice3)

	//Collections
	presAge := make(map[string]int) //Key is string, value is int
	presAge["Hello"] = 99
	fmt.Println("presAge[Hello] ==>", presAge["Hello"])
	fmt.Println("Length ==> ", len(presAge))

	presAge["Hello Welcome"] = 77
	fmt.Println("Length11 ==> ", len(presAge))
	fmt.Println("presAge ==>", presAge)

	delete(presAge, "Hello") //Delete the data for which key was passed
	fmt.Println("Length22 ==> ", len(presAge))

	//Function calling .....
	listNums := []float64{1, 2, 3, 4, 5}
	fmt.Println("Sum : ", addThemUp(listNums))

	//Retruning two values from function
	num1, num2 := next2Values(5)
	fmt.Println("Retruning two values ==> ", num1, num2)

	//Without defining the no of arguments to function
	fmt.Println("Any no of arguments ==>", subtractThem(1, 2, 3, 4, 5))

	// Declaring function inside function and calling
	num3 := 3
	doubleNum := func() int {
		num3 *= 2

		return num3
	}
	fmt.Println("Func inside func : ", doubleNum())

	defer printTwo()
	//Recusive function
	fmt.Println("Recursive call : ", fact(3))

	//defer concept
	defer printTwo()
	printOne()

	// recover -- Save Div
	fmt.Println("Div with zero :", safeDiv(3, 0))
	fmt.Println("Div with number : ", safeDiv(3, 1))

	//Catching Exception
	demPanic()

	//Without Pointers uses
	x := 0
	changeXVal(x)
	fmt.Println("Just X = ", x)

	//Pointers uses
	x1 := 0
	changeXValNow(&x1)
	fmt.Println("Pointers X1 = ", x1)
	fmt.Println("Memory Address for X1 = ", &x1)

	//Creating & Using Pointer
	yPtr := new(int)
	changeYValNow(yPtr)
	fmt.Println("Y = ", *yPtr)

	//struct uses....
	//	rect1 := Rectangle{leftX: 0, topY: 50, height: 100, width: 200}
	//or -- if we know the order
	rect1 := Rectangle{10, 50}
	fmt.Println("Rectangle is --->>>>>>>   ", rect1.width, "wide")
	fmt.Println("Area of Rectangle  --->>>>>>>   ", rect1.area())

	//calculating Area of Rectangle and Circle
	rect2 := Rectangle{height: 10, width: 50}
	cir2 := Circle{radius: 10}

	fmt.Println("Rectangle Area =  ", getArea(rect2))
	fmt.Println("Rectangle Circle =  ", getArea(cir2))

	// String Functions
	tempString := "Hello World"
	fmt.Println(strings.Contains(tempString, "lo"))
	fmt.Println(strings.Index(tempString, "lo"))
	fmt.Println(strings.Count(tempString, "l"))
	fmt.Println(strings.Replace(tempString, "l", "X", 3)) //replace first 3

	csvString := "1,2,3,4,5"
	fmt.Println("Splited String := ", strings.Split(csvString, ","))

	listOfLetters := []string{"c", "a", "b"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters := ", listOfLetters)

	listOfNums := strings.Join([]string{"3", "2", "1"}, ",")
	fmt.Println("listOfNums: ==", listOfNums)

	// Dealing with File System
	file, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("This is some random text")
	file.Close()

	stream, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)
	fmt.Println("readString := ", readString)

	//Data type conversion
	randInt := 5
	randFloat := 10.9
	randString := "100"
	randString2 := "250.5"

	fmt.Println("Integer0 := ", float64(randInt))
	fmt.Println("Float0 := ", int(randFloat))

	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println("Integer1 := ", newInt)

	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println("Float1 := ", newFloat)

}

//Main ending here !!!

//Defining interface
type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

//func retuning area of Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

//func retuning area of Circle
func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {
	fmt.Println("Shape --> ", shape)
	return shape.area()
}

// ******************************************************************

func addThemUp(numbers []float64) float64 {

	sum := 0.0

	for _, val := range numbers {
		sum += val
	}
	return sum
}

func next2Values(number int) (int, int) {

	return number + 1, number + 2
}

////Without defining the no of arguments to function
func subtractThem(args ...int) int {

	finalValue := 0

	for _, value := range args {
		finalValue += value
	}
	return finalValue
}

//Recusive function
func fact(no int) int {

	if no == 0 {
		return 1
	}
	return no * fact(no-1)
}

//defer concept
func printOne() { fmt.Println("Print1 : ", 1) }
func printTwo() { fmt.Println("Print2 : ", 2) }

// recover -- Save Div
func safeDiv(num1, num2 int) int {

	defer func() {
		recover() //Program wouldn`t fail when deviser is 0
	}()
	solution := num1 / num2
	return solution
}

//Catching Exception
func demPanic() {
	defer func() {
		fmt.Println("---->>> ", recover())
	}()
	panic("Exception Handling : PANIC")
}

func changeXVal(x int) {

	fmt.Println("X = ", x)
}

//Pointers == putting * with datatype
func changeXValNow(x *int) {

	*x = 2
}

//Pointers == putting * with datatype
func changeYValNow(yPt *int) {

	*yPt = 100
}

/*
func makeDough(stringChan chan string) {
	pizzaName++
	pizzaName = "Pizza # " + strconv.I
}
*/
