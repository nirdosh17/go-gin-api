package storage

import (
	"gorm.io/gorm"
)

// if a new field is added. auto migration will run and will create or remove field
type Product struct {
	gorm.Model
	Category string
	Name     string
	Price    uint
}

type ApiKey struct {
	gorm.Model
	ApiKey string
}

type CreateRequest struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    uint   `json:"price"`
}
