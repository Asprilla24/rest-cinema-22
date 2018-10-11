package store

import (
	"database/sql"
	"errors"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Repository struct{}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Movie struct {
	ID        string          `json:"id"`
	Title     string          `json:"title"`
	Publisher string          `json:"publisher"`
	Rating    float64         `json:"rating"`
	Cover     string          `json:"cover"`
	Category  []MovieCategory `json:"category"`
}

type MovieCategory struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	AgeFrom  int    `json:"age_from"`
	AgeTo    int    `json:"age_to"`
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

	query, err := db.Query("SELECT Id, Title, Publisher, Rating, Cover, CategoryId FROM Movie")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	for query.Next() {
		var movie Movie
		var categoryIDs string
		err := query.Scan(&movie.ID, &movie.Title, &movie.Publisher, &movie.Rating, &movie.Cover, &categoryIDs)
		if err != nil {
			panic(err.Error())
		}

		categoryID := strings.Split(categoryIDs, ",")

		println(categoryIDs)
		println(categoryID)

		movie.Category = getCategory(r, categoryID)

		result = append(result, movie)
	}

	return result
}

func (r Repository) getAllCategory() []MovieCategory {
	result := []MovieCategory{}

	db := connect()
	defer db.Close()

	query, err := db.Query("SELECT Id, Category, AgeFrom, AgeTo FROM MovieCategory")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	for query.Next() {
		var category MovieCategory
		err := query.Scan(&category.ID, &category.Category, &category.AgeFrom, &category.AgeTo)
		if err != nil {
			panic(err.Error())
		}

		result = append(result, category)
	}

	return result
}

func (r Repository) getCategory(ID string) MovieCategory {
	result := MovieCategory{}

	db := connect()
	defer db.Close()

	query, err := db.Query("SELECT Id, Category, AgeFrom, AgeTo FROM MovieCategory WHERE Id = ?", ID)
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	for query.Next() {
		err := query.Scan(&result.ID, &result.Category, &result.AgeFrom, &result.AgeTo)
		if err != nil {
			panic(err.Error())
		}
	}

	return result
}

func getCategory(r Repository, IDs []string) []MovieCategory {
	var result []MovieCategory

	for _, id := range IDs {
		category := r.getCategory(id)
		result = append(result, category)
	}

	return result
}
