package user

import (
	"encoding/json"
	"log"
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
	if err := dec.Decode(newUser); err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.s.RegisterUser(newUser)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("new user: %s %s", newUser.GetEmail(), newUser.GetPass())
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User Registered"))
}
