package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"webDB/internal/models"
)

func Init() *gorm.DB {
	// localhost или название контейнера с бд - database в данном случае
	dns := "postgres://pg:pass@database:5432/crud"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
