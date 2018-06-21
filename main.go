package main

import (
	"github.com/tqoliver/grogar/routes"
	"log"
	"net/http"
)

func main() {

	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", r))

}
