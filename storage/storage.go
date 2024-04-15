package storage

import (
	"context"
	"errors"
	"github.com/vlad-marlo/example/model"
	"sync"
)

type Storage struct {
	data      *sync.Map
	lastIndex int64
}

func NewStorage() *Storage {
	s := &Storage{
		data:      new(sync.Map),
		lastIndex: 0,
	}
	return s
}

func (s *Storage) Store(_ context.Context, item model.TodoCreateRequest) (*model.Todo, error) {
	todo := &model.Todo{
		ID:          s.lastIndex + 1,
		Name:        item.Name,
		Description: item.Description,
	}
	s.lastIndex += 1
	s.data.Store(todo.ID, todo)
	return todo, nil
}

func (s *Storage) GetByID(_ context.Context, id int64) (*model.Todo, error) {
	item, ok := s.data.Load(id)
	if !ok {
		return nil, errors.New("item does not exists in storage")
	}

	return item.(*model.Todo), nil
}

func (s *Storage) GetAll(_ context.Context) (res []*model.Todo, err error) {
	res = make([]*model.Todo, 0, s.lastIndex)
	s.data.Range(func(_, value any) bool {
		res = append(res, value.(*model.Todo))
		return true
	})
	return res, nil
}

func (s *Storage) Delete(_ context.Context, id int64) error {
	s.data.Delete(id)
	return nil
}
