package models

import "github.com/google/uuid"

type Room struct {
	Base         `mapstructure:",squash"`
	Number       string    `json:"number" binding:"required"` // physical room number
	Class        string    `json:"class" binding:"required"`
	People_range string    `json:"people" binding:"required"` // e.g. 1-4
	Description  string    `json:"description" binding:"required"`
	Price        int       `json:"price" binding:"required"` // day price
	Floor        int       `json:"floor" binding:"required"`
	Title        string    `json:"title" binding:"required"`
	Available    bool      `json:"available" default:"true"`
	InvoiceId    uuid.UUID `json:"invoice_id"`
}
