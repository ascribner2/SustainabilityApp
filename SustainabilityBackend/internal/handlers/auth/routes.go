package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ascribner/sustainabilityapp/internal/entity"
	"github.com/ascribner/sustainabilityapp/internal/services"
)

type Handler struct {
	s services.UserService
}

func NewUserHandler(us services.UserService) *Handler {
	return &Handler{
		s: us,
	}
}

func (h *Handler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("/login", h.login)
}

func (h *Handler) login(rw http.ResponseWriter, r *http.Request) {
	user := &entity.User{}

	dec := json.NewDecoder(r.Body)
	dec.Decode(user)

}
