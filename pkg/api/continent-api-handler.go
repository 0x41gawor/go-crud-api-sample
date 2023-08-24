package api

import (
	"fmt"
	"net/http"

	"github.com/0x41gawor/go-crud-api-sample/pkg/repo"
)

type ContinentApiHandler struct {
	repo repo.ContinentRepository
}

func NewContinentApiHandler(repo repo.ContinentRepository) *ContinentApiHandler {
	return &ContinentApiHandler{
		repo: repo,
	}
}

// handles "/continent" endpoint
func (this *ContinentApiHandler) handleContinent(w http.ResponseWriter, r *http.Request) error {
	res, err := this.repo.List()
	if err != nil {
		fmt.Println(err.Error())
	}
	return WriteJSON(w, http.StatusOK, res)
}

// handles "/continent/{id}" endpoint
func (this *ContinentApiHandler) handleContinentId(w http.ResponseWriter, r *http.Request) error {
	return nil
}
