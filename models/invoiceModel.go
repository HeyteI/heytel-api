package models

import "github.com/google/uuid"

type Invoice struct {
	Base           `mapstructure:",squash"`
	RoomId         uuid.UUID `json:"room_id" binding:"required"`    // ID of room
	UserId         uuid.UUID `json:"user_id" binding:"required"`    // ID of user that invoiced
	Date           string    `json:"date_range" binding:"required"` // Date format: "22/08/2022-29/08/2022"
	Cancelled      bool      `json:"cancelled" gorm:"default:false"`
	Paid           bool      `json:"paid" gorm:"default:false"`
	PaymentMethod  string    `json:"payment_method" binding:"required"` // cash, credit_card
	RegisterNumber string    `json:"register_number" binding:"required"`
	PeopleCount    int       `json:"people_count" binding:"required"`
}
