package queries

import (
	model "sdte-backend/models"
)

func GetAllData() ([]model.Data, error) {
	sqlStatement := "SELECT * FROM major_ava;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	// release connection resource when finish this function
	defer rows.Close()
	list := make([]model.Data, 0)
	for rows.Next() {
		var oneUser model.Data
		if err := rows.Scan(&oneUser.MajorID, &oneUser.MajorName, &oneUser.Major, &oneUser.MajorDes, &oneUser.OpenDate); err != nil {
			return []model.Data{}, err
		}
		list = append(list, oneUser)
	}
	if err = rows.Err(); err != nil {
		return []model.Data{}, err
	}
	return list, nil
}
