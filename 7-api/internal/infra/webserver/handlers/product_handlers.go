package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/mateus-sousa/goexpert/7-api/internal/dto"
	"github.com/mateus-sousa/goexpert/7-api/internal/entity"
	"github.com/mateus-sousa/goexpert/7-api/internal/infra/database"
	entityPkg "github.com/mateus-sousa/goexpert/7-api/pkg/entity"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	repository database.ProductInterface
}

func NewProductHandler(repository database.ProductInterface) *ProductHandler {
	return &ProductHandler{repository: repository}
}

// Create Product godoc
//
//	@Summary      Create product
//	@Description  Create products
//	@Tags         products
//	@Accept       json
//	@Produce      json
//	@Param        request    body     dto.CreateProductInput  true  "product request"
//	@Success      201
//	@Failure      500  {object}  Error
//	@Router       /products [post]
//
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&productDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := entity.NewProduct(productDto.Name, productDto.Price)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	err = h.repository.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// List Products godoc
//
//	@Summary      List products
//	@Description  get all products
//	@Tags         products
//	@Accept       json
//	@Produce      json
//	@Param        page    query     string  false  "page number"
//	@Param        limit    query     string  false  "limit"
//	@Param        sort    query     string  false  "sort"
//	@Success      200  {array}  entity.Product
//	@Failure      404  {object}  Error
//	@Failure      500  {object}  Error
//	@Router       /products [get]
//
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limit := r.URL.Query().Get("limit")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}
	sort := r.URL.Query().Get("sort")
	products, err := h.repository.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Get Product godoc
//
//	@Summary      Get a  product
//	@Description  Get a product
//	@Tags         products
//	@Accept       json
//	@Produce      json
//	@Param        id    path     string  true  "product ID" Format(uuid)
//	@Success      200  {object}  entity.Product
//	@Failure      404
//	@Failure      500  {object}  Error
//	@Router       /products/{id} [get]
//
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	product, err := h.repository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// Update Product godoc
//
//	@Summary      Update a product
//	@Description  Update a product
//	@Tags         products
//	@Accept       json
//	@Produce      json
//	@Param        id    path     string  true  "product ID" Format(uuid)
//	@Param        request    body     dto.CreateProductInput  true  "product request"
//	@Success      200
//	@Success      404
//	@Failure      500  {object}  Error
//	@Router       /products/{id} [put]
//
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.repository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.repository.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete Products godoc
//
//	@Summary      Delete a product
//	@Description  Delete a product
//	@Tags         products
//	@Accept       json
//	@Produce      json
//	@Param        id    path     string  true  "product ID" Format(uuid)
//	@Success      200
//	@Failure      404
//	@Failure      500  {object}  Error
//	@Router       /products/{id} [delete]
//
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := h.repository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.repository.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
