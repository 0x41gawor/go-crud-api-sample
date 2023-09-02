package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiHandlerImage struct {
	path string
}

func NewApiHandlerImage() *ApiHandlerImage {
	return &ApiHandlerImage{
		path: "web/test.jpg",
	}
}

// handles "/image" endpoint
func (h *ApiHandlerImage) handleImage(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		fileBytes, err := ioutil.ReadFile(h.path)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(fileBytes)
		return nil
	default:
		return WriteJSON(w, http.StatusOK, "error: method not allowed")
	}
}
