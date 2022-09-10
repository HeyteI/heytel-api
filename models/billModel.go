package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Bill struct {
	Base      `mapstructure:",squash"`
	InvoiceId uuid.UUID `json:"invoice_id" binding:"required"`
	Services  pq.StringArray  `json:"services" gorm:"type:text[]" binding:"required" default:"[]"` // format ["service:price"]
}
