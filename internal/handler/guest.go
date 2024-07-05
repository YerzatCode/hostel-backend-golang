package handler

import (
	"net/http"

	"github.com/YerzatCode/sql_/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/goccy/go-json"
)

func (h *Handler) GuestCreate(w http.ResponseWriter, r *http.Request) {
	var guest model.GuestModel
	err := json.NewDecoder(r.Body).Decode(&guest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg, err := h.service.Guests.Create(guest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondwithJSON(w, http.StatusCreated, msg)
}

func (h *Handler) GuestGetAll(w http.ResponseWriter, r *http.Request) {

	quests, err := h.service.Guests.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondwithJSON(w, http.StatusAccepted, quests)
}

func (h *Handler) GuestGetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	guest, err := h.service.Guests.GetOne(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondwithJSON(w, http.StatusAccepted, guest)
}

func (h *Handler) GuestDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	msg, err := h.service.Guests.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondwithJSON(w, http.StatusAccepted, msg)
}
