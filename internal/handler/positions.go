package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YerzatCode/sql_/internal/model"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) PositionsCreate(w http.ResponseWriter, r *http.Request) {
	var position model.PositionModel
	err := json.NewDecoder(r.Body).Decode(&position)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.service.Positions.Create(position)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusCreated, id)
}

func (h *Handler) PositionsGetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.Positions.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusAccepted, data)
}

func (h *Handler) PositionsGetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data, err := h.service.Positions.GetOne(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusAccepted, data)
}

func (h *Handler) PositionsDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	msg, err := h.service.Positions.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusAccepted, msg)
}
