package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tqoliver/grogar/handlers"
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
	fmt.Fprintf(w, "You've reached the home page of Grogar the App")
}

//EmployeeDb handle the employee function
func EmployeeDb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, handlers.EmployeeDb(10))
}

//SystemInfo handle the system info function
func SystemInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, handlers.SystemInfo())
}

//DvdRentalDb handle the DVD Rental Info
func DvdRentalDb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, handlers.DvdRentalDB())
}

//DvdInfo handle the DVD Info Requests
func DvdInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, handlers.DvdInfo())
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
		"/home",
		Index,
	},

	Route{
		"Employee List",
		"GET",
		"/employee/list",
		EmployeeDb,
	},

	Route{
		"System/Container Info",
		"GET",
		"/info/system",
		SystemInfo,
	},

	Route{
		"DVD Rental",
		"GET",
		"/info/dvd/rental",
		DvdRentalDb,
	},

	Route{
		"DVD Films",
		"GET",
		"/info/dvd/films",
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
