package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // no-lint
	"github.com/spf13/viper"
)

type Person struct {
	ID              int
	Name            string
	Description     string
	Email           string
	Phone           string
	Relation        string
	AlternateEmails string
	AlternatePhones string
	Notes           string
	Todos           string
}

func GetPersons() ([]Person, error) {
	connectionString := viper.GetString("CONNECTION_STRING")
	driver := viper.GetString("DB_DRIVER")
	db, err := sql.Open(driver, connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection to database with error:%w", err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT Id, Name FROM Person")
	if err != nil {
		return nil, fmt.Errorf("failed to query from database with error:%w", err)
	}
	defer rows.Close()

	var id int
	var name string
	var persons []Person
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan next row with error:%w", err)
		}
		person := Person{ID: id, Name: name}
		persons = append(persons, person)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("the cursor gave error:%w", err)
	}
	return persons, nil
}
