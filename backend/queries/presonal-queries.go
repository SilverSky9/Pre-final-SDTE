package queries
 
import(
	model "sdte-backend/models"
)
func GetAllPerson() ([]model.Person, error) {
	sqlStatement := "SELECT * FROM persons;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	// release connection resource when finish this function
	defer rows.Close()
	list := make([]model.Person, 0)
	for rows.Next() {
		var oneUser model.Person
		if err := rows.Scan(&oneUser.PersonID, &oneUser.LastName, &oneUser.FirstName, &oneUser.Address, &oneUser.City); err != nil {
			return []model.Person{}, err
		}
		list = append(list, oneUser)
	}
	if err = rows.Err(); err != nil {
		return []model.Person{}, err
	}
	return list, nil
}
