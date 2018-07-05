package helpers

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

//EmpSvcClient this is the API client
type EmpSvcClient struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

//GetEmployees will return a list of films from the database
func GetEmployees() (string, error) {

	url := "http://employeeservices:8000/v1/employees"
	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "apiClient")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("err: %s", err)
		return "", err
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), err
}
