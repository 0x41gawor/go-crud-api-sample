package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0x41gawor/go-crud-api-sample/pkg/model"
	"github.com/0x41gawor/go-crud-api-sample/pkg/repo"
)

type CountryApiHandler struct {
	repo repo.CountryRepository
}

func NewCountryApiHandler(repo repo.CountryRepository) *CountryApiHandler {
	return &CountryApiHandler{
		repo: repo,
	}
}

// handles "/country" endpoint
func (this *CountryApiHandler) handleCountry(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return this.List(w, r)
	case "POST":
		return this.Create(w, r)
	default:
		return WriteJSON(w, http.StatusOK, "error: method not allowed")
	}
}

func (this *CountryApiHandler) handleCountryId(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return this.Read(w, r)
	default:
		return WriteJSON(w, http.StatusOK, "error: method not allowed")
	}
}

func (this *CountryApiHandler) List(w http.ResponseWriter, r *http.Request) error {
	res, err := this.repo.List()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, res)
}

func (this *CountryApiHandler) Create(w http.ResponseWriter, r *http.Request) error {
	model := new(model.Country)

	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}
	id, err := this.repo.Create(model)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	return WriteJSON(w, http.StatusOK, fmt.Sprintf("createdId: %d", id))
}

func (this *CountryApiHandler) Read(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	model, err := this.repo.Read(int64(id))

	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	return WriteJSON(w, http.StatusOK, model)
}
