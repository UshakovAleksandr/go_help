package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"pg/model"
)

var DB *gorm.DB

func Connect() {
	dsn := "postgres://user:mypassword@localhost:5432/testdb?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model.Tokens{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

//func GetDB() *gorm.DB {
//	if DB == nil {
//		DB = Init()
//		var sleep = time.Duration(1)
//		for DB == nil {
//			sleep = sleep * 2
//			fmt.Printf("database is unavailable. Wait for %d sec...\n", sleep)
//			time.Sleep(sleep * time.Second)
//			DB = Init()
//		}
//	}
//	return DB
//}
