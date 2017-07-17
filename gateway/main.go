package main

import (
	fmt "fmt"
	consul "gateway/consul"
	http "net/http"
)

func main() {
	client, err := consul.NewConsulClient("http://10.225.241.195:8500")
	if err != nil {
		panic(err)
	}
	address, meta, err := client.Service("city-service", "")

	for i, service := range address {
		fmt.Println(service.Service.Address, service.Service.Port, service.Service.ID, meta.LastContact, i)
	}
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", homeHandler)
	err = http.ListenAndServe(":8083", nil)
	if err != nil {
		panic(err)
	}
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>%s</h1>", "Hi, Service is UP and Running!")
}
