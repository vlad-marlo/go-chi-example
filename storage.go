package main

import (
	"context"
	"errors"
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

func (s *Storage) Store(_ context.Context, item TodoCreateRequest) (*Todo, error) {
	todo := &Todo{
		ID:          s.lastIndex + 1,
		Name:        item.Name,
		Description: item.Description,
	}
	s.lastIndex += 1
	s.data.Store(todo.ID, todo)
	return todo, nil
}

func (s *Storage) GetByID(_ context.Context, id int64) (*Todo, error) {
	item, ok := s.data.Load(id)
	if !ok {
		return nil, errors.New("item does not exists in storage")
	}

	return item.(*Todo), nil
}

func (s *Storage) GetAll(_ context.Context) (res []*Todo, err error) {
	res = make([]*Todo, 0, s.lastIndex)
	s.data.Range(func(_, value any) bool {
		res = append(res, value.(*Todo))
		return true
	})
	return res, nil
}

func (s *Storage) Delete(_ context.Context, id int64) error {
	s.data.Delete(id)
	return nil
}
