package handler

import (
	"github.com/YerzatCode/sql_/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *chi.Mux {
	route := chi.NewRouter()

	route.Route("/guest", func(r chi.Router) {
		r.Post("/", h.GuestCreate)
		r.Get("/", h.GuestGetAll)
		r.Get("/{id}", h.GuestGetOne)
		r.Delete("/{id}", h.GuestDelete)
	})
	route.Route("/room", func(r chi.Router) {
		r.Post("/", h.RoomCreate)
		r.Get("/", h.RoomGetAll)
		r.Get("/{id}", h.RoomGetOne)
		r.Delete("/{id}", h.RoomDelete)
	})

	route.Route("/category", func(r chi.Router) {
		r.Post("/", h.CategoriesCreate)
		r.Get("/", h.CategoriesGetAll)
		r.Get("/{id}", h.CategoriesGetOne)
		r.Delete("/{id}", h.CategoriesDelete)
	})

	route.Route("/occupied", func(r chi.Router) {
		r.Post("/", h.OccupiedCreate)
		r.Get("/", h.OccupiedGetAll)
		r.Get("/{id}", h.OccupiedGetOne)
		r.Delete("/{id}", h.OccupiedDelete)
	})

	route.Route("/additional", func(r chi.Router) {
		r.Post("/", h.AdditionalCreate)
		r.Get("/", h.AdditionalGetAll)
		r.Get("/{id}", h.AdditionalGetOne)
		r.Delete("/{id}", h.AdditionalDelete)
	})

	route.Route("/payment", func(r chi.Router) {
		r.Post("/", h.PaymentCreate)
		r.Get("/", h.PaymentGetAll)
		r.Get("/{id}", h.PaymentGetOne)
		r.Delete("/{id}", h.PaymentDelete)
	})

	route.Route("/employees", func(r chi.Router) {
		r.Post("/", h.EmployeesCreate)
		r.Get("/", h.EmployeesGetAll)
		r.Get("/{id}", h.EmployeesGetOne)
		r.Delete("/{id}", h.EmployeesDelete)
	})

	route.Route("/position", func(r chi.Router) {
		r.Post("/", h.PositionsCreate)
		r.Get("/", h.PositionsGetAll)
		r.Get("/{id}", h.PositionsGetOne)
		r.Delete("/{id}", h.PositionsDelete)
	})

	route.Route("/children", func(r chi.Router) {
		r.Post("/", h.ChildrenCreate)
		r.Get("/", h.ChildrenGetAll)
		r.Get("/{id}", h.ChildrenGetOne)
		r.Delete("/{id}", h.ChildrenDelete)
	})

	return route
}
