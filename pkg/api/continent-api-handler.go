package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0x41gawor/go-crud-api-sample/pkg/model"
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

	if r.Method == "GET" {
		return this.read(w, r)
	}
	if r.Method == "DELETE" {
		return this.delete(w, r)
	}

	return WriteJSON(w, http.StatusOK, "res: method not allowed")
}

func (this *ContinentApiHandler) list(w http.ResponseWriter, r *http.Request) error {
	res, err := this.repo.List()
	if err != nil {
		fmt.Println(err.Error())
	}
	return WriteJSON(w, http.StatusOK, res)
}

func (this *ContinentApiHandler) create(w http.ResponseWriter, r *http.Request) error {
	model := new(model.Continent)

	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	id, err := this.repo.Create(model)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return WriteJSON(w, http.StatusOK, fmt.Sprintf("createdId: %d", id))
}

func (this *ContinentApiHandler) read(w http.ResponseWriter, r *http.Request) error {
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

func (this *ContinentApiHandler) update(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (this *ContinentApiHandler) delete(w http.ResponseWriter, r *http.Request) error {
	return nil
}
