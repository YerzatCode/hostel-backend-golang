package service

import (
	"github.com/YerzatCode/sql_/internal/model"
	"github.com/YerzatCode/sql_/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Guests interface {
	Create(guest model.GuestModel) (string, error)
	GetAll() ([]model.GuestModel, error)
	GetOne(guestId string) (model.GuestModel, error)
	Delete(guestId string) (string, error)
	Update(guestId string, guest model.GuestModel) error
}

type Rooms interface {
	Create(room model.RoomModel) (string, error)
	GetAll() ([]model.RoomModel, error)
	GetOne(roomId string) (model.RoomModel, error)
	Delete(roomId string) (string, error)
}

type Categories interface {
	Create(category model.CategoriesModel) (string, error)
	GetAll() ([]model.CategoriesModel, error)
	GetOne(categoryId string) (model.CategoriesModel, error)
	Delete(categoryId string) (string, error)
}

type OccupiedRooms interface {
	Create(occupiedRoom model.OccupiedRoomsModel) (string, error)
	GetAll() ([]model.OccupiedRoomsModel, error)
	GetOne(occupiedRoomId string) (model.OccupiedRoomsModel, error)
	Delete(occupiedRoomId string) (string, error)
}

type AdditionalServices interface {
	Create(additionalService model.AdditionalServicesModel) (string, error)
	GetAll() ([]model.AdditionalServicesModel, error)
	GetOne(additionalServiceId string) (model.AdditionalServicesModel, error)
	Delete(additionalServiceId string) (string, error)
}

type PaymentMethods interface {
	Create(paymentMethod model.PaymentMethodsModel) (string, error)
	GetAll() ([]model.PaymentMethodsModel, error)
	GetOne(paymentMethodId string) (model.PaymentMethodsModel, error)
	Delete(paymentMethodId string) (string, error)
}

type Employees interface {
	Create(employee model.EmployeesModel) (string, error)
	GetAll() ([]model.EmployeesModel, error)
	GetOne(employeeId string) (model.EmployeesModel, error)
	Delete(employeeId string) (string, error)
}

type Positions interface {
	Create(position model.PositionModel) (string, error)
	GetAll() ([]model.PositionModel, error)
	GetOne(positionId string) (model.PositionModel, error)
	Delete(positionId string) (string, error)
}

type Children interface {
	Create(children model.ChildrenModel) (string, error)
	GetAll() ([]model.ChildrenModel, error)
	GetOne(childrenId string) (model.ChildrenModel, error)
	Delete(childrenId string) (string, error)
}

type Service struct {
	Guests
	Rooms
	Categories
	OccupiedRooms
	AdditionalServices
	PaymentMethods
	Employees
	Positions
	Children
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Guests:             NewGuestService(repo.Guests),
		Rooms:              NewRoomService(repo.Rooms),
		Categories:         NewCategoriesService(repo.Categories),
		OccupiedRooms:      NewOccupiedService(repo.OccupiedRooms),
		AdditionalServices: NewAddService(repo.AdditionalServices),
		PaymentMethods:     NewPayMetService(repo.PaymentMethods),
		Employees:          NewEmployeesService(repo.Employees),
		Positions:          NewPositionsService(repo.Positions),
		Children:           NewChildrenService(repo.Children),
	}
}
