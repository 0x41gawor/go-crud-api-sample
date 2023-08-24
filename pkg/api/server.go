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

	continentRepository, err := repo.NewContinentRepository()
	if err != nil {
		fmt.Print(err.Error())
	}
	continentApiHandler := NewContinentApiHandler(*continentRepository)

	router.HandleFunc("/continent", makeHTTPHandleFunc(continentApiHandler.handleContinent))
	router.HandleFunc("/continent/", makeHTTPHandleFunc(continentApiHandler.handleContinent))
	router.HandleFunc("/continent/{id}", makeHTTPHandleFunc(continentApiHandler.handleContinentId))

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
