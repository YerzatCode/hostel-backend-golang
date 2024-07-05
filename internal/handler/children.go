package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YerzatCode/sql_/internal/model"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) ChildrenCreate(w http.ResponseWriter, r *http.Request) {
	var children model.ChildrenModel
	err := json.NewDecoder(r.Body).Decode(&children)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := h.service.Children.Create(children)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusCreated, id)
}

func (h *Handler) ChildrenGetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.Children.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusAccepted, data)
}

func (h *Handler) ChildrenGetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data, err := h.service.Children.GetOne(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusAccepted, data)
}

func (h *Handler) ChildrenDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	msg, err := h.service.Children.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusAccepted, msg)
}
