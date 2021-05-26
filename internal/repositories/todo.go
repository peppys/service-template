package repositories

import (
	"context"
	"errors"
	"github.com/google/uuid"
	todo "github.com/peppys/service-template/gen/go/proto"
	"time"
)

var todos = map[string]*todo.Todo{}

type TodoRepository struct {
}

func (r *TodoRepository) FindAll(ctx context.Context) ([]*todo.Todo, error) {
	values := make([]*todo.Todo, 0, len(todos))
	for _, val := range todos {
		values = append(values, val)
	}

	return values, nil
}

func (r *TodoRepository) Create(ctx context.Context, text string, author string) (*todo.Todo, error) {
	t := &todo.Todo{
		Id:        uuid.New().String(),
		Text:      text,
		Author:    author,
		Timestamp: time.Now().String(),
	}
	todos[t.GetId()] = t

	return t, nil
}

func (r *TodoRepository) FindById(ctx context.Context, id string) (*todo.Todo, error) {
	t := todos[id]
	if t == nil {
		return nil, errors.New("does not exist")
	}

	return t, nil
}
