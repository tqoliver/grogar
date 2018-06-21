SELECT 
customer.customer_id, 
customer.first_name, 
customer.last_name, 
customer.email, 
rental.rental_date, 
inventory.inventory_id, 
film.title, 
film.description, 
film.rating, 
film.release_year, 
language.name, 
category.name 
FROM 
customer 
INNER JOIN rental 
ON customer.customer_id = rental.customer_id
INNER JOIN inventory 
ON rental.inventory_id = inventory.inventory_id
INNER JOIN film 
ON inventory.film_id = film.film_id
INNER JOIN language
ON film.language_id = language.language_id
INNER JOIN film_category
ON film.film_id = film_category.film_id
INNER JOIN category
ON film_category.category_id = category.category_id
LIMIT 15;

rows, err := db.Query(
		"SELECT customer.customer_id,customer.first_name, customer.last_name,customer.email,rental.rental_date,inventory.inventory_id,film.title,film.description, film.rating, film.release_year, language.name, category.name FROM customer INNER JOIN rental ON customer.customer_id = rental.customer_id INNER JOIN inventory ON rental.inventory_id = inventory.inventory_id INNER JOIN film ON inventory.film_id = film.film_id INNER JOIN language ON film.language_id = language.language_id INNER JOIN film_category ON film.film_id = film_category.film_id INNER JOIN categoryN film_category.category_id = category.category_id LIMIT 15")
