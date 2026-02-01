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

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// HandleCategories - GET/POST /api/categories
func (h *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			h.GetAll(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			libs.HandleError(http.StatusMethodNotAllowed, w, "Method not allowed")
	}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		libs.HandleError(http.StatusInternalServerError, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusOK, w, products)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {		
		libs.HandleError(http.StatusBadRequest, w, "Invalid request body")
		return
	}

	err = h.service.Create(&category)
	if err != nil {
		libs.HandleError(http.StatusInternalServerError, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusCreated, w, category)
}

// HandleProductByID - GET/PUT/DELETE /api/categories/{id}
func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
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

// GetByID - GET /api/categories/{id}
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		libs.HandleError(http.StatusBadRequest, w, "Invalid category ID")
		return
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		libs.HandleError(http.StatusNotFound, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusOK, w, category)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		libs.HandleError(http.StatusBadRequest, w, "Invalid category ID")
		return
	}

	var category models.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		libs.HandleError(http.StatusBadRequest, w, "Invalid request body")
		return
	}

	category.ID = id
	err = h.service.Update(&category)
	if err != nil {
		libs.HandleError(http.StatusInternalServerError, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusOK, w, category)
}


// Delete - DELETE /api/categories/{id}
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		libs.HandleError(http.StatusBadRequest, w, "Invalid category ID")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		libs.HandleError(http.StatusInternalServerError, w, err.Error())
		return
	}

	libs.HandleResponse(http.StatusOK, w, nil, "Category deleted successfully")
}