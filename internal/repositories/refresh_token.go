package repositories

import (
	"context"
	"github.com/peppys/service-template/internal/entities"
	"gorm.io/gorm"
	"time"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db,
	}
}

func (r *RefreshTokenRepository) Save(ctx context.Context, refreshToken *entities.RefreshToken) (*entities.RefreshToken, error) {
	result := r.db.Debug().Create(&refreshToken)
	return refreshToken, result.Error
}

func (r *RefreshTokenRepository) FindFirst(ctx context.Context, query entities.RefreshToken) (*entities.RefreshToken, error) {
	var refreshToken *entities.RefreshToken
	result := r.db.Debug().Where(query).Where("expires_at > ?", time.Now()).First(&refreshToken)
	return refreshToken, result.Error
}

func (r *RefreshTokenRepository) Delete(ctx context.Context, query entities.RefreshToken) error {
	result := r.db.Debug().Where(query).Delete(&entities.RefreshToken{})
	return result.Error
}
