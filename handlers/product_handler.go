package handlers

import (
	"kasir-api/services"
	"net/http"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetProducts(w, r)
	case http.MethodPost:
		h.CreateProduct(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

}
