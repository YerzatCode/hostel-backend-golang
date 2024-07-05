package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YerzatCode/sql_/internal/model"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) AdditionalCreate(w http.ResponseWriter, r *http.Request) {
	var additionalService model.AdditionalServicesModel
	err := json.NewDecoder(r.Body).Decode(&additionalService)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.service.AdditionalServices.Create(additionalService)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusCreated, id)
}

func (h *Handler) AdditionalGetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.AdditionalServices.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusAccepted, data)
}

func (h *Handler) AdditionalGetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	additionalService, err := h.service.AdditionalServices.GetOne(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusAccepted, additionalService)
}

func (h *Handler) AdditionalDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	msg, err := h.service.AdditionalServices.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusAccepted, msg)
}
