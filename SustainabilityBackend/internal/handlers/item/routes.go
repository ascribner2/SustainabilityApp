package item

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ascribner/sustainabilityapp/internal/entity"
	"github.com/ascribner/sustainabilityapp/internal/middleware"
	"github.com/ascribner/sustainabilityapp/internal/services"
)

type Handler struct {
	is services.ItemService
}

func NewItemHandler(is services.ItemService) *Handler {
	return &Handler{
		is: is,
	}
}

func (h *Handler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("/additem", middleware.AuthenticateRoute(h.addItem))
	r.HandleFunc("/getitems", middleware.AuthenticateRoute(h.getItems))
}

// Method for /additem route
func (h *Handler) addItem(rw http.ResponseWriter, r *http.Request) {
	item := entity.Item{}
	dec := json.NewDecoder(r.Body)
	dec.Decode(&item)

	email, err := GetEmailFromContext(r.Context())
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		return
	}

	if err := h.is.AddItem(&item, email); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

// method for /getitems route
func (h *Handler) getItems(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	email, err := GetEmailFromContext(r.Context())
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		return
	}

	items, totalOffset, err := h.is.GetItems(email, queryParams.Get("timespan"))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(rw)
	enc.Encode(map[string]any{"Items": items, "TotalOffset": totalOffset})
}
