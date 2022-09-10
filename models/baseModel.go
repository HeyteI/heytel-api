package models

import (
	"github.com/google/uuid"
	"time"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; in database!