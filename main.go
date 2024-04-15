package main

import (
	"github.com/vlad-marlo/example/controller"
	"github.com/vlad-marlo/example/storage"
)

func main() {
	store := storage.NewStorage()
	app := controller.New(store)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
