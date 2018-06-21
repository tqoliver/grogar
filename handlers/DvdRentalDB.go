package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	//github.com/lib/pq is a comment because the editor needed it
	_ "github.com/lib/pq"
	"os"
	"time"
)

//DvdRentalDB returns customer data from a PostgreSQL microservice database in a container
func DvdRentalDB() string {

	var (
		dbUser     = os.Getenv("PG_DBUSER")     //postgres
		dbPassword = os.Getenv("PG_DBPASSWORD") //postgres
		dbName     = os.Getenv("PG_DATABASE")   //dvdrental
		dbHost     = os.Getenv("PG_DBHOST")     //"192.168.64.3"
		dbPort     = os.Getenv("PG_DBPORT")     //"32072"
	)

	type DvdRental struct {
		CustomerID      string    `json:"customerID"`
		FirstName       string    `json:"firstName"`
		LastName        string    `json:"lastName"`
		Email           string    `json:"email"`
		RentalDate      time.Time `json:"rentalDate"`
		InventoryID     int       `json:"inventoryID"`
		FilmTitle       string    `json:"filmTitle"`
		FilmDescription string    `json:"filmDescription"`
		FilmRating      string    `json:"filmRating"`
		FilmReleaseYear int64     `json:"filmReleaseYear"`
		LanguageName    string    `json:"languageName"`
		CategoryName    string    `json:"categoryName"`
	}

	var dr DvdRental
	var drs []DvdRental
	var dbinfo string

	dbinfo = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	fmt.Printf("dbinfo: %s\n\n", dbinfo)

	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query(
		"SELECT customer.customer_id,customer.first_name, customer.last_name," +
			"customer.email,rental.rental_date,inventory.inventory_id,film.title," +
			"film.description, film.rating, film.release_year, language.name, category.name " +
			"FROM customer INNER JOIN rental ON customer.customer_id = rental.customer_id " +
			"INNER JOIN inventory ON rental.inventory_id = inventory.inventory_id " +
			"INNER JOIN film ON inventory.film_id = film.film_id " +
			"INNER JOIN language ON film.language_id = language.language_id " +
			"INNER JOIN film_category ON film.film_id = film_category.film_id " +
			"INNER JOIN category ON film_category.category_id = category.category_id LIMIT 100")

	checkErr(err)

	for rows.Next() {

		err := rows.Scan(&dr.CustomerID, &dr.FirstName, &dr.LastName, &dr.Email,
			&dr.RentalDate, &dr.InventoryID, &dr.FilmTitle, &dr.FilmDescription,
			&dr.FilmRating, &dr.FilmReleaseYear, &dr.LanguageName, &dr.CategoryName)
		checkErr(err)

		drs = append(drs, dr)

	}

	strJSON, err := json.Marshal(drs)
	checkErr(err)
	return string(strJSON)
}
