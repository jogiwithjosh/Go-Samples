package dto

type Country struct {
  Name string `json:"name, omitempty"`
  Code string `json:"code, omitempty"`
  Capital string `json:"capital, omitempty"`
}
