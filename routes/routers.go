package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tqoliver/grogar/helpers"
	"log"
	"net/http"
	"time"
)

//Route As the name implies these are the API's incoming legitimate routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes slice of Route
type Routes []Route

//NewRouter function
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

//Index function
func Index(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	day := t.Format("2006 January _2 03:04:05PM MST")
	fmt.Fprintf(w, "<h1>Hello Demo Crowd<br>You've reached the home page of test App.<br>The current time is: "+day+"</h1>")
}

//EmployeeDb handle the employee function
func EmployeeDb(w http.ResponseWriter, r *http.Request) {
	s, err := helpers.GetEmployees()
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Encountered an error %g", err))
		return
	}
	fmt.Fprintf(w, s)
}

//SystemInfo handle the system info function
func SystemInfo(w http.ResponseWriter, r *http.Request) {
	s, err := helpers.GetSystemInfo()
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error encountered: %s", err))
		return
	}
	fmt.Fprintf(w, s)
}

//DvdRentalDb handle the DVD Rental Info
func DvdRentalDb(w http.ResponseWriter, r *http.Request) {
	s, err := helpers.GetRentals()
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error encountered: %s", err))
		return
	}

	fmt.Fprintf(w, s)
}

//DvdInfo handle the DVD Info Requests
func DvdInfo(w http.ResponseWriter, r *http.Request) {
	s, err := helpers.GetFilms()
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error encountered: %s", err))
		return
	}
	fmt.Fprintf(w, s)
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"Home",
		"GET",
		"/v1/home",
		Index,
	},

	Route{
		"Employee List",
		"GET",
		"/v1/employee/list",
		EmployeeDb,
	},

	Route{
		"System/Container Info",
		"GET",
		"/v1/info/system",
		SystemInfo,
	},

	Route{
		"DVD Rental",
		"GET",
		"/v1/info/dvd/rental",
		DvdRentalDb,
	},

	Route{
		"DVD Films",
		"GET",
		"/v1/info/dvd/films",
		DvdInfo,
	},
}

//Logger function
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
