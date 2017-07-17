package main

import (
	dto "country-service/dto"
	"encoding/json"
	fmt "fmt"
	http "net/http"

	consul "city-service/consul"
)

func main() {
	client, err := consul.NewConsulClient("http://10.225.241.195:8500")
	if err != nil {
		panic(err)
	}
	err = client.Register("country-service", 8081)
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8081", NewRouter())
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(dto.Country{})
	fmt.Fprintf(writer, "<h1>%s</h1>", "Hi, Country-Service is UP and Running!")
}

func HealthHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	var respone map[string]string
	respone["output"] = "Country-Service is UP and Runnig!!"
	if err := json.NewEncoder(writer).Encode(respone); err != nil {
		panic(err)
	}
}
