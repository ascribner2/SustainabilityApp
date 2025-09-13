package user

import (
	"encoding/json"
	"fmt"
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
	r.HandleFunc("/register", h.registerUser)
}

func (h *Handler) registerUser(rw http.ResponseWriter, r *http.Request) {
	newUser := &entity.NewUser{}

	dec := json.NewDecoder(r.Body)
	dec.Decode(newUser)

	fmt.Print(newUser)

	// h.s.RegisterUser()

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User Registered"))
}
