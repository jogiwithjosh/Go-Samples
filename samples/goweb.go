package main

import (
	"fmt"
	"net/http"
)

//main will be called when an application start executing
func main() {
	fmt.Println("Hello World !!!")

	//Dealing with Web
	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)
	http.ListenAndServe(":8093", nil)

}

//Main ending here !!!

//Dealing with Web
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth\n")
}
