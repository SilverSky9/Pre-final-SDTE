package config

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"fmt"
)

const (
	host = "localhost"
	port = 5432
)

var db *sql.DB = Connect_db()

func init() {
	Connect_db()
}

func Connect_db() *sql.DB {
	//Create psql Info for connect your Postgres DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, "postgres", "magical", "final")
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		// panic(err)
		fmt.Println(NotConnectDatabaseWarning())
	}
	// err = db.Ping()
	if err != nil {
		// panic(err)
		fmt.Println(NotConnectDatabaseWarning())
	}
	return db
}

func GetDB() *sql.DB {
	return db
}
func NotConnectDatabaseWarning() error {
	return errors.New("Go Service isn't connect database, But if you want to test unitest is normal")
}
