package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YerzatCode/sql_/internal/model"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) RoomCreate(w http.ResponseWriter, r *http.Request) {
	var room model.RoomModel
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg, err := h.service.Rooms.Create(room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusCreated, msg)
}

func (h *Handler) RoomGetAll(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.service.Rooms.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusCreated, rooms)
}

func (h *Handler) RoomGetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	room, err := h.service.Rooms.GetOne(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusCreated, room)
}

func (h *Handler) RoomDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	msg, err := h.service.Rooms.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusCreated, msg)
}
