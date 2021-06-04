package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Text      string
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
