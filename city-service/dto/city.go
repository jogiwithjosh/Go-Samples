package dto

type City struct {
  Name string `json:"name, omitempty"`
  CountryId int `json:"countryId, omitempty"`
  Code string `json:"code, omitempty"`
}
