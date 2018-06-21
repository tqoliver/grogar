package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	//github.com/lib/pq has a comment because the editor wants it
	_ "github.com/lib/pq"
	"os"
)

//DvdInfo will return data on films from a PostgreSQL microservice
func DvdInfo() string {

	var (
		dbUser     = os.Getenv("PG_DBUSER")     //postgres
		dbPassword = os.Getenv("PG_DBPASSWORD") //postgres
		dbName     = os.Getenv("PG_DATABASE")   //dvdrental
		dbHost     = os.Getenv("PG_DBHOST")     //"192.168.64.3"
		dbPort     = os.Getenv("PG_DBPORT")     //"32072"
	)

	type DvdData struct {
		CategoryName    string `json:"categoryName"`
		FilmID          int    `json:"filmID"`
		FilmTitle       string `json:"filmTitle"`
		FilmDescription string `json:"filmDescription"`
		FilmFullText    string `json:"filmFullText"`
	}

	var dds []DvdData
	var dbinfo string

	dbinfo = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query(
		"SELECT category.name, film_category.film_id, film.title, film.description, film.fulltext " +
			"FROM category " +
			"INNER JOIN film_category on category.category_id = film_category.category_id " +
			"INNER JOIN film on film_category.film_id = film.film_Id LIMIT 200")
	checkErr(err)

	for rows.Next() {
		var dd DvdData
		rows.Scan(&dd.CategoryName, &dd.FilmID, &dd.FilmTitle, &dd.FilmDescription, &dd.FilmFullText)
		dds = append(dds, dd)
	}
	strJSON, err := json.Marshal(dds)
	checkErr(err)
	return string(strJSON)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
