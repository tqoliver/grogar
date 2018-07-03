package helpers

import (
	"io/ioutil"
	"net/http"
	"time"
)

//GetSystemInfo will return a list of films from the database
func GetSystemInfo() (string, error) {
	url := "http://svcsystem:8000/v1/info/system"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	client := http.Client{
		Timeout: time.Second * 2,
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "API Client System Info")

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
