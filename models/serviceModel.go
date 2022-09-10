package models

type Service struct {
	Base
	Name string `json:"name" binding:"required"`
	Price int `json:"suggested_price" binding:"required"`
	Description string `json:"description"`
}