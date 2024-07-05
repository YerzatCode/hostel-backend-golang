package model

type PaymentMethodsModel struct {
	ID          int     `json:"id"`
	GuestID     int     `json:"GuestID"`
	PaymentForm string  `json:"paymentForm"`
	PaidAmount  float64 `json:"paidAmount"`
}
