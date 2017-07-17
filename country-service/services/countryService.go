package services

import (
  dto "country-service/dto"
  errors "errors"
)

var countries []dto.Country

func Save(newcountry dto.Country) (string, error) {
  for _,country := range countries {
    if country.Code == newcountry.Code {
      return "", errors.New("Unable to save, Country already exists!")
    }
  }

  countries = append(countries, newcountry)
  return newcountry.Code, nil
}

func FindOne(code string) (dto.Country, error) {
  for _,country := range countries {
    if country.Code == code {
      return country, nil
    }
  }
  return dto.Country{}, errors.New("Sorry, No Country found with Code: " + code)
}
