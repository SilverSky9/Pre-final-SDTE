package services

import (
	"sdte-backend/models"
	"sdte-backend/queries"
)

func GetAllFaculty() ([]models.Faculty, error) {
	resp, err := queries.GetAllFaculty()
	if err != nil {
		return []models.Faculty{}, err
	}
	return resp, nil
}
