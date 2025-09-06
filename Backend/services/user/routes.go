package user

import "net/http"

type Handler struct {
}

func NewUserHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("/register", registerUser)
}

func registerUser(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("test"))
}
