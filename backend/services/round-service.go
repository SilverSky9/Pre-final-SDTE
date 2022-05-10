package services

import (
	"sdte-backend/models"
	"sdte-backend/queries"
)

func GetAllRound() ([]models.Round, error) {
	resp, err := queries.GetAllRound()
	if err != nil {
		return []models.Round{}, err
	}
	return resp, nil
}
