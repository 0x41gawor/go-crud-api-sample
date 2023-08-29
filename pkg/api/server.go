package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/0x41gawor/go-crud-api-sample/pkg/repo"
	"github.com/gorilla/mux"
)

type Server struct {
	listenPort string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenPort: listenAddr,
	}
}

func (this *Server) Run() {
	router := mux.NewRouter()

	dbConnectionHelper, err := repo.NewDatabaseConnectionHelper()
	if err != nil {
		fmt.Print(err.Error())
	}
	continentRepository := repo.NewContinentRepository(dbConnectionHelper.DB)
	continentApiHandler := NewContinentApiHandler(*continentRepository)

	countryRepository := repo.NewCountryRepository(dbConnectionHelper.DB)
	countryApiHandler := NewCountryApiHandler(*countryRepository)

	router.HandleFunc("/continent", makeHTTPHandleFunc(continentApiHandler.handleContinent))
	router.HandleFunc("/continent/", makeHTTPHandleFunc(continentApiHandler.handleContinent))
	router.HandleFunc("/continent/{id}", makeHTTPHandleFunc(continentApiHandler.handleContinentId))

	router.HandleFunc("/country", makeHTTPHandleFunc(countryApiHandler.handleCountry))
	router.HandleFunc("/country/", makeHTTPHandleFunc(countryApiHandler.handleCountry))
	router.HandleFunc("/country/{id}", makeHTTPHandleFunc(countryApiHandler.handleCountryId))

	log.Println("JSON API server running on port: ", this.listenPort)
	http.ListenAndServe(this.listenPort, router)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle error here
		}
	}
}
