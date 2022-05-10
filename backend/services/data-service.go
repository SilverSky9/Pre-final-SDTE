package services

import (
	"sdte-backend/models"
	"sdte-backend/queries"
)

func GetAllData() ([]models.Data, error) {
	resp, err := queries.GetAllData()
	if err != nil {
		return []models.Data{}, err
	}
	return resp, nil
}
