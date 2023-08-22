package main

import (
	"log"

	"github.com/spearexit/dropick.core/v2/pkg/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func main() {
	db, err := models.SetConnection()
	if err != nil {
		log.Fatal(err)
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{})

	m.InitSchema(func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&models.Workflow{}, &models.Document{}, &models.Group{}, &models.Tag{}, &models.User{}); err != nil {
			return err
		}
		return nil
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration did run successfully")
}