package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Email         string
	EmailVerified bool
	Username      string
	PasswordHash  string
	GivenName     string
	FamilyName    string
	Name          string
	Nickname      string
	Picture       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
