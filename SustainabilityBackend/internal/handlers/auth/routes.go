package auth

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ascribner/sustainabilityapp/internal/entity"
	"github.com/ascribner/sustainabilityapp/internal/middleware"
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
	r.HandleFunc("/verify", middleware.AuthenticateRoute(h.verify))
	r.HandleFunc("/logout", h.logout)
}

func (h *Handler) login(rw http.ResponseWriter, r *http.Request) {
	user := &entity.User{}

	dec := json.NewDecoder(r.Body)
	dec.Decode(user)

	authenticated, err := h.s.UserLogin(user)

	// If the login matches and there are no errors encode and send JWT
	if authenticated && err == nil {
		token, err := CreateJWT(user.GetEmail())
		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(rw, &http.Cookie{
			Name:     "auth-token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Secure:   false,
		})
	} else if err == sql.ErrNoRows || !authenticated {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Invalid email or password"))
	} else {
		log.Print(err)
		rw.WriteHeader(http.StatusBadRequest)
	}
}

func (h *Handler) verify(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

func (h *Handler) logout(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Set-Cookie", "auth-token=; HttpOnly; Secure; Path=/; Max-Age=0")
	http.SetCookie(rw, &http.Cookie{
		Name:     "auth-token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   0,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})
}
