package index

import (
	"fmt"
	"net/http"
)

func NewHandler() http.Handler {
	return &handler{}
}

type handler struct {

}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		h.get(w, req)
	case http.MethodPost:
		h.post(w, req)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// get godoc
// @Summary Site index
// @Success 200 {object} string
// @Router / [get]
func (h *handler) get(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintln(w, "hello world")
}

// post godoc
// @Summary Site index
// @Router / [post]
func (h *handler) post(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintln(w, "hello world")
}