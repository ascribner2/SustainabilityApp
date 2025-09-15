package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ascribner/sustainabilityapp/internal/entity"
	"github.com/ascribner/sustainabilityapp/internal/services"
)

type Handler struct {
	s services.AuthService
}

func NewAuthHandler(as services.AuthService) *Handler {
	return &Handler{
		s: as,
	}
}

func (h *Handler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("/login", h.login)
}

func (h *Handler) login(rw http.ResponseWriter, r *http.Request) {
	user := &entity.User{}

	dec := json.NewDecoder(r.Body)
	dec.Decode(user)

	authenticated, err := h.s.UserLogin(user)
	if authenticated && err == nil {
		enc := json.NewEncoder(rw)
		if err = enc.Encode(map[string]string{"response": "JWT"}); err != nil {
			log.Print(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		rw.WriteHeader(http.StatusOK)
	} else {
		log.Print(err)
		rw.WriteHeader(http.StatusBadRequest)
	}
}
