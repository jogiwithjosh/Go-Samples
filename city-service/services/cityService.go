package services

import (
  dto "city-service/dto"
  errors "errors"
)

var cities []dto.City

func Save(newcity dto.City) (string, error) {
  for _,city := range cities {
    if city.Code == newcity.Code {
      return "", errors.New("Unable to save, City already exists!")
    }
  }

  cities = append(cities, newcity)
  return newcity.Code, nil
}

func FindOne(code string) (dto.City, error) {
  for _,city := range cities {
    if city.Code == code {
      return city, nil
    }
  }
  return dto.City{}, errors.New("Sorry, No City found with Code: " + code)
}
