package api

import (
	"net/http"

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
		return this.list(w, r)
	default:
		return WriteJSON(w, http.StatusOK, "error: method not allowed")
	}
}

func (this *CountryApiHandler) list(w http.ResponseWriter, r *http.Request) error {
	res, err := this.repo.List()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, res)
}
