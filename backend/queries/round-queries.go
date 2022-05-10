package queries

import (
	model "sdte-backend/models"
)

func GetAllRound() ([]model.Round, error) {
	sqlStatement := "SELECT * FROM round_ava;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	// release connection resource when finish this function
	defer rows.Close()
	list := make([]model.Round, 0)
	for rows.Next() {
		var oneUser model.Round
		if err := rows.Scan(&oneUser.RoundID, &oneUser.RoundName); err != nil {
			return []model.Round{}, err
		}
		list = append(list, oneUser)
	}
	if err = rows.Err(); err != nil {
		return []model.Round{}, err
	}
	return list, nil
}
