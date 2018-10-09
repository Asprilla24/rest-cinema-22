package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Repository struct{}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1)/cinema22")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (r Repository) login(Username string, Password string) User {
	db := connect()
	defer db.Close()

	query, err := db.Query("SELECT Id, Username, Email FROM User WHERE Username = ? AND Password = ?", Username, Password)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	var result User
	for query.Next() {
		if err := query.Scan(&result.Id, &result.Username, &result.Email); err != nil {
			panic(err.Error())
		}
	}

	return result
}
