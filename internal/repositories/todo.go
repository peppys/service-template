package repositories

import (
	"context"
	"github.com/google/uuid"
	"github.com/peppys/service-template/internal/entities"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db,
	}
}

func (r *TodoRepository) FindAll(ctx context.Context) ([]*entities.Todo, error) {
	var todos []*entities.Todo
	result := r.db.Debug().Find(&todos)
	return todos, result.Error
}

func (r *TodoRepository) Create(ctx context.Context, todo *entities.Todo) (*entities.Todo, error) {
	result := r.db.Debug().Create(todo)
	return todo, result.Error
}

func (r *TodoRepository) FindById(ctx context.Context, id string) (*entities.Todo, error) {
	recordUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	todo := &entities.Todo{ID: recordUuid}
	result := r.db.Debug().First(&todo)
	return todo, result.Error
}
