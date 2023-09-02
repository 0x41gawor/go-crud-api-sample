package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0x41gawor/go-crud-api-sample/pkg/model"
	"github.com/0x41gawor/go-crud-api-sample/pkg/repo"
)

type ApiHandlerContinent struct {
	repo repo.RepositoryContinent
}

func NewContinentApiHandler(repo repo.RepositoryContinent) *ApiHandlerContinent {
	return &ApiHandlerContinent{
		repo: repo,
	}
}

// handles "/continent" endpoint
func (h *ApiHandlerContinent) handleContinent(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return h.list(w, r)
	case "POST":
		return h.create(w, r)
	default:
		return WriteJSON(w, http.StatusOK, "error: method not allowed")
	}
}

// handles "/continent/{id}" endpoint
func (h *ApiHandlerContinent) handleContinentId(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return h.read(w, r)
	case "PUT":
		return h.update(w, r)
	case "DELETE":
		return h.delete(w, r)
	default:
		return WriteJSON(w, http.StatusOK, "error: method not allowed")
	}
}

func (h *ApiHandlerContinent) list(w http.ResponseWriter, r *http.Request) error {
	res, err := h.repo.List()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, res)
}

func (h *ApiHandlerContinent) create(w http.ResponseWriter, r *http.Request) error {
	model := new(model.Continent)

	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	id, err := h.repo.Create(model)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	return WriteJSON(w, http.StatusOK, fmt.Sprintf("createdId: %d", id))
}

func (h *ApiHandlerContinent) read(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	model, err := h.repo.Read(int64(id))

	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	return WriteJSON(w, http.StatusOK, model)
}

func (h *ApiHandlerContinent) update(w http.ResponseWriter, r *http.Request) error {
	model := new(model.Continent)

	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	id, err := getID(r)

	model.Id = id

	err = h.repo.Update(int64(id), model)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	updatedModel, err := h.repo.Read(int64(id))
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	return WriteJSON(w, http.StatusOK, updatedModel)
}

func (h *ApiHandlerContinent) delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	res, err := h.repo.Delete(int64(id))
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	if res == false {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("res: no item deleted "))
	}

	return WriteJSON(w, http.StatusOK, fmt.Sprintf("res: item with %d deleted ", id))
}
