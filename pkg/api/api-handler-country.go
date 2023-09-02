package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0x41gawor/go-crud-api-sample/pkg/model"
	"github.com/0x41gawor/go-crud-api-sample/pkg/repo"
)

type ApiHandlerCountry struct {
	repo repo.CountryRepository
}

func NewCountryApiHandler(repo repo.CountryRepository) *ApiHandlerCountry {
	return &ApiHandlerCountry{
		repo: repo,
	}
}

// handles "/country" endpoint
func (h *ApiHandlerCountry) handleCountry(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return h.list(w, r)
	case "POST":
		return h.create(w, r)
	default:
		return WriteJSON(w, http.StatusOK, "error: method not allowed")
	}
}

// handles "/country/{id}" endpoint
func (h *ApiHandlerCountry) handleCountryId(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return h.read(w, r)
	case "POST":
		return h.update(w, r)
	case "DELETE":
		return h.delete(w, r)
	default:
		return WriteJSON(w, http.StatusOK, "error: method not allowed")
	}
}

func (h *ApiHandlerCountry) list(w http.ResponseWriter, r *http.Request) error {
	res, err := h.repo.List()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, res)
}

func (h *ApiHandlerCountry) create(w http.ResponseWriter, r *http.Request) error {
	model := new(model.Country)

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

func (h *ApiHandlerCountry) read(w http.ResponseWriter, r *http.Request) error {
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

func (h *ApiHandlerCountry) update(w http.ResponseWriter, r *http.Request) error {
	model := new(model.Country)

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

func (h *ApiHandlerCountry) delete(w http.ResponseWriter, r *http.Request) error {
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
