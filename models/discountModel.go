package models

type Discount struct {
	Base
	Name string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
	Description string `json:"description"`
	Active bool `json:"active" default:"true"`
}