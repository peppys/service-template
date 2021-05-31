package repositories

import (
	"context"
	"github.com/peppys/service-template/internal/entities"
	"gorm.io/gorm"
)

type TodoMysqlRepository struct {
	db *gorm.DB
}

func NewTodoMysqlRepository(db *gorm.DB) *TodoMysqlRepository {
	return &TodoMysqlRepository{
		db,
	}
}

func (r *TodoMysqlRepository) FindAll(ctx context.Context) ([]*entities.Todo, error) {
	var todos []*entities.Todo
	result := r.db.Find(&todos)
	return todos, result.Error
}

func (r *TodoMysqlRepository) Create(ctx context.Context, todo *entities.Todo) (*entities.Todo, error) {
	result := r.db.Create(todo)
	return todo, result.Error
}

func (r *TodoMysqlRepository) FindById(ctx context.Context, id string) (*entities.Todo, error) {
	var todo *entities.Todo
	result := r.db.First(&todo, id)
	return todo, result.Error
}
