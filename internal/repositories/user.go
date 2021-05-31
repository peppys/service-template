package repositories

import (
	"context"
	"github.com/peppys/service-template/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) Save(ctx context.Context, user *entities.User) (*entities.User, error) {
	result := r.db.Debug().Create(user)
	return user, result.Error
}

func (r *UserRepository) FindFirst(ctx context.Context, query entities.User) (*entities.User, error) {
	var user *entities.User
	result := r.db.Debug().Where(query).First(&user)
	return user, result.Error
}
