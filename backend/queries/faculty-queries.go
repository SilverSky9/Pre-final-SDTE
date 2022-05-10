package queries

import (
	model "sdte-backend/models"
)

func GetAllFaculty() ([]model.Faculty, error) {
	sqlStatement := "SELECT * FROM faculty;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	// release connection resource when finish this function
	defer rows.Close()
	list := make([]model.Faculty, 0)
	for rows.Next() {
		var oneUser model.Faculty
		if err := rows.Scan(&oneUser.FacultyID, &oneUser.FacultyName); err != nil {
			return []model.Faculty{}, err
		}
		list = append(list, oneUser)
	}
	if err = rows.Err(); err != nil {
		return []model.Faculty{}, err
	}
	return list, nil
}
