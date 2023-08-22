package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spearexit/dropick.core/v2/pkg/apis"
	"github.com/spearexit/dropick.core/v2/pkg/models"
)

func main() {
	db, err := models.SetConnection()
	if err != nil {
		log.Fatal(err)
	}

	app := &apis.App{
		Db:       db,
		Validate: validator.New(),
	}

	app.Start()
}