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

	return WriteJSON(w, http.StatusOK, "error: method not allowed")
}

// handles "/continent/{id}" endpoint
func (this *ContinentApiHandler) handleContinentId(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return this.read(w, r)
	}
	if r.Method == "PUT" {
		return this.update(w, r)
	}
	if r.Method == "DELETE" {
		return this.delete(w, r)
	}

	return WriteJSON(w, http.StatusOK, "error: method not allowed")
}

func (this *ContinentApiHandler) list(w http.ResponseWriter, r *http.Request) error {
	res, err := this.repo.List()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, res)
}

func (this *ContinentApiHandler) create(w http.ResponseWriter, r *http.Request) error {
	model := new(model.Continent)

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
	model := new(model.Continent)

	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	id, err := getID(r)

	model.Id = id

	err = this.repo.Update(int64(id), model)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	updatedModel, err := this.repo.Read(int64(id))
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	return WriteJSON(w, http.StatusOK, updatedModel)
}

func (this *ContinentApiHandler) delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	res, err := this.repo.Delete(int64(id))
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	if res == false {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("res: no item deleted "))
	}

	return WriteJSON(w, http.StatusOK, fmt.Sprintf("res: item with %d deleted ", id))
}
