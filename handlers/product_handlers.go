package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/services"
	"kasir-api/libs"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// HandleProducts - GET /api/products
func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			h.GetAll(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			libs.HandleError(http.StatusMethodNotAllowed, w, "Method not allowed")
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	products, err := h.service.GetAll(name)
	if err != nil {
		libs.HandleError(http.StatusInternalServerError, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusOK, w, products)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateProductRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		libs.HandleError(http.StatusBadRequest, w, "Invalid request body")
		return
	}

	product, err := h.service.Create(req)
	if err != nil {
		libs.HandleError(http.StatusInternalServerError, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusCreated, w, product)
}

// HandleProductByID - GET/PUT/DELETE /api/products/{id}
func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			h.GetByID(w, r)
		case http.MethodPut:
			h.Update(w, r)
		case http.MethodDelete:
			h.Delete(w, r)
		default:
			libs.HandleError(http.StatusMethodNotAllowed, w, "Method not allowed")
	}
}

// GetByID - GET /api/products/{id}
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		libs.HandleError(http.StatusBadRequest, w, "Invalid product ID")
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		libs.HandleError(http.StatusNotFound, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusOK, w, product)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		libs.HandleError(http.StatusBadRequest, w, "Invalid product ID")
		return
	}

	var req models.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		libs.HandleError(http.StatusBadRequest, w, "Invalid request body")
		return
	}

	product, err := h.service.Update(id, req)
	if err != nil {
		libs.HandleError(http.StatusInternalServerError, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusOK, w, product)
}


// Delete - DELETE /api/products/{id}
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		libs.HandleError(http.StatusBadRequest, w, "Invalid product ID")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		libs.HandleError(http.StatusInternalServerError, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusOK, w, nil, "Product deleted successfully")
}