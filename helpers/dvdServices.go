package helpers

import (
	"io/ioutil"
	"net/http"
	"time"
)

//GetFilms will return a list of films from the database
func GetFilms() (string, error) {
	url := "http://dvdservice:8000/v1/dvd/rentals"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	client := http.Client{
		Timeout: time.Second * 2,
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "API Client")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), err
}

//GetRentals will return a list of films from the database
func GetRentals() (string, error) {
	url := "http://dvdservice:8000/v1/dvd/films"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	client := http.Client{
		Timeout: time.Second * 2,
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "API Client DVD Services")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), err
}
