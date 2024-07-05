package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YerzatCode/sql_/internal/model"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) OccupiedCreate(w http.ResponseWriter, r *http.Request) {
	var occupiedRoom model.OccupiedRoomsModel
	err := json.NewDecoder(r.Body).Decode(&occupiedRoom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg, err := h.service.OccupiedRooms.Create(occupiedRoom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondwithJSON(w, http.StatusCreated, msg)
}

func (h *Handler) OccupiedGetAll(w http.ResponseWriter, r *http.Request) {
	occupiedRooms, err := h.service.OccupiedRooms.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusAccepted, occupiedRooms)
}

func (h *Handler) OccupiedGetOne(w http.ResponseWriter, r *http.Request) {
	occupiedRoomId := chi.URLParam(r, "id")
	occupiedRoom, err := h.service.OccupiedRooms.GetOne(occupiedRoomId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondwithJSON(w, http.StatusAccepted, occupiedRoom)
}

func (h *Handler) OccupiedDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	msg, err := h.service.OccupiedRooms.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusAccepted, msg)
}
