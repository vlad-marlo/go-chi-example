package main

type Todo struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TodoCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetByIDRequest struct {
	ID int64 `param:"id"`
}
