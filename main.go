package main

import (
	"github.com/vlad-marlo/example/config"
	"github.com/vlad-marlo/example/controller"
	"github.com/vlad-marlo/example/storage"
)

func main() {
	store := storage.NewStorage()
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	app := controller.New(store, cfg)
	if err = app.Run(); err != nil {
		panic(err)
	}
}
