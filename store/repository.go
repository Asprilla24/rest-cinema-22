package store

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type Repository struct{}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Movie struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Publisher string  `json:"publisher"`
	Rating    float64 `json:"rating"`
	Cover     string  `json:"cover"`
}

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1)/cinema22")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (r Repository) login(Username string, Password string) (User, error) {
	var result User

	db := connect()
	defer db.Close()

	query, err := db.Query("SELECT Id, Username, Email FROM User WHERE Username = ? AND Password = ?", Username, Password)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	for query.Next() {
		if err := query.Scan(&result.ID, &result.Username, &result.Email); err != nil {
			panic(err.Error())
		}
	}

	if result.ID == "" {
		return result, errors.New("User not found")
	}

	return result, nil
}

func (r Repository) getAllMovie() []Movie {
	result := []Movie{}

	db := connect()
	defer db.Close()

	query, err := db.Query("SELECT Id, Title, Publisher, Rating, Cover FROM Movie")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	for query.Next() {
		var movie Movie
		err := query.Scan(&movie.ID, &movie.Title, &movie.Publisher, &movie.Rating, &movie.Cover)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, movie)
	}

	return result
}
