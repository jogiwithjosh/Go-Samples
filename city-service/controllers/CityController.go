package controllers

import (
  http "net/http"
  json "encoding/json"
  cityService "city-service/services"
  mux "github.com/gorilla/mux"
  dto "city-service/dto"
  io "io"
	ioutil "io/ioutil"
  //strconv "strconv"
)

func Save(writer http.ResponseWriter, request *http.Request) {

    var city dto.City

    //read body from request. this data will be []byte
  	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
  	if err != nil {
  		panic(err)
  	}
  	if err := request.Body.Close(); err != nil {
  		panic(err)
  	}

    //convert []byte to Country struct
    if err := json.Unmarshal(body, &city); err != nil {
		  writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		  writer.WriteHeader(422) // unprocessable entity
		  if err := json.NewEncoder(writer).Encode(err); err != nil {
			  panic(err)
		  }
	}

  code, err := cityService.Save(city)
  if err != nil {
    if err = json.NewEncoder(writer).Encode(err); err != nil {
  			panic(err)
  		}
  }
  writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
  if err = json.NewEncoder(writer).Encode(code); err != nil {
			panic(err)
		}
}

func FindOne(writer http.ResponseWriter, request *http.Request) {
  var code string
  var err error
  vars := mux.Vars(request)
  code = string(vars["code"])

  city, err := cityService.FindOne(code)
  if err != nil {
    if err = json.NewEncoder(writer).Encode(err); err != nil {
  			panic(err)
  		}
      return
  }
  writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
  if err = json.NewEncoder(writer).Encode(city); err != nil {
			panic(err)
		}
}
