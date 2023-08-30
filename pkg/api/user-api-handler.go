package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/0x41gawor/go-crud-api-sample/pkg/model"
	"github.com/0x41gawor/go-crud-api-sample/pkg/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserApiHandler struct {
	repo repo.UserRepository
}

func NewUserApiHandler(repo repo.UserRepository) *UserApiHandler {
	return &UserApiHandler{
		repo: repo,
	}
}

// handles "/register" endpoint
func (this *UserApiHandler) handleRegister(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return this.create(w, r)
	default:
		return WriteJSON(w, http.StatusOK, "error: method not allowed")
	}
}

// handles "/login" endpoint
func (this *UserApiHandler) handleLogin(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "POST":
		return this.login(w, r)
	default:
		return WriteJSON(w, http.StatusOK, "error: method not allowed")
	}
}

func (this *UserApiHandler) create(w http.ResponseWriter, r *http.Request) error {
	model := new(model.User)

	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}
	_, err = this.repo.Create(model)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	return WriteJSON(w, http.StatusOK, "res: registered sucessfully")
}

func (this *UserApiHandler) login(w http.ResponseWriter, r *http.Request) error {
	modelAttemp := new(model.User)

	err := json.NewDecoder(r.Body).Decode(modelAttemp)

	modelFromRepo, err := this.repo.FindByLogin(modelAttemp.Login)
	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	if bcrypt.CompareHashAndPassword([]byte(modelFromRepo.Password), []byte(modelAttemp.Password)) != nil {
		return WriteJSON(w, http.StatusOK, "res: permission denied")
	}

	tokenStr, err := CreateJWT(modelFromRepo.Login)

	if err != nil {
		return WriteJSON(w, http.StatusOK, fmt.Sprintf("error: %s", err.Error()))
	}

	return WriteJSON(w, http.StatusOK, fmt.Sprintf("{res: login successful, Bearer: %s}", tokenStr))
}
