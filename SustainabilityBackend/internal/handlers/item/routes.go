package item

import (
	"net/http"

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
	r.HandleFunc("/additem", h.addItem)
	r.HandleFunc("/getitems", h.getItems)
}

func (h *Handler) addItem(rw http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getItems(rw http.ResponseWriter, r *http.Request) {

}
