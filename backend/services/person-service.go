package services

import (
	"sdte-backend/models"
	"sdte-backend/queries"
)

func GetAllPerson() ([]models.Person, error) {
	resp, err := queries.GetAllPerson()
	if err != nil {
		return []models.Person{}, err
	}
	return resp, nil
}
