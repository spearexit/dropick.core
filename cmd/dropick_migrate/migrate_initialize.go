package main

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spearexit/dropick.core/v2/internal/models"
	"github.com/spearexit/dropick.core/v2/internal/shared"
	"gorm.io/gorm"
)

func main() {
	db, err := shared.SetConnection(shared.Config.Database)
	if err != nil {
		log.Fatal(err)
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{})

	m.InitSchema(func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(&models.Workflow{}, &models.Document{}, &models.Group{}, &models.User{}); err != nil {
			return err
		}
		return nil
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration did run successfully")
}
