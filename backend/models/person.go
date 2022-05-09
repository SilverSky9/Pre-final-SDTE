package models

type Person struct {
	PersonID  int    `json:"personid"`
	LastName  string `json:"lastname"`
	FirstName string `json:"firstname"`
	Address   string `json:"address"`
	City      string `json:"city"`
}
