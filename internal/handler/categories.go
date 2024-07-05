package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YerzatCode/sql_/internal/model"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) CategoriesCreate(w http.ResponseWriter, r *http.Request) {
	var category model.CategoriesModel
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg, err := h.service.Categories.Create(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondwithJSON(w, http.StatusCreated, msg)
}

func (h *Handler) CategoriesGetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.Categories.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusAccepted, data)
}

func (h *Handler) CategoriesGetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data, err := h.service.Categories.GetOne(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondwithJSON(w, http.StatusAccepted, data)
}

func (h *Handler) CategoriesDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	msg, err := h.service.Categories.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondwithJSON(w, http.StatusAccepted, msg)
}
