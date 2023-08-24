package api

import (
	"fmt"
	"go/types"
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
	if r.Method == "GET" {
		return this.list(w, r)
	}
	if r.Method == "POST" {
		return this.create(w, r)
	}

	return WriteJSON(w, http.StatusOK, "res: method not allowed")
}

// handles "/continent/{id}" endpoint
func (this *ContinentApiHandler) handleContinentId(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (this *ContinentApiHandler) list(w http.ResponseWriter, r *http.Request) error {
	res, err := this.repo.List()
	if err != nil {
		fmt.Println(err.Error())
	}
	return WriteJSON(w, http.StatusOK, res)
}

func (this *ContinentApiHandler) create(w http.ResponseWriter, r *http.Request) error {
	model := new()

	return nil
}

func (this *ContinentApiHandler) read(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (this *ContinentApiHandler) update(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (this *ContinentApiHandler) delete(w http.ResponseWriter, r *http.Request) error {
	return nil
}
