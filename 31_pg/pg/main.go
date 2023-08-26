package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"pg/model"
)

func main() {
	//database.Connect()
	dsn := "postgres://pg:pass@localhost:5432/crud"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.Tokens{})
	fmt.Println(111)
}


//fmt.Println("connecting")
//// these details match the docker-compose.yml file.
//postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//"postgres", 5432, "user", "mypassword", "user")
//db, err := sql.Open("postgres", postgresInfo)
//if err != nil {
//panic(err)
//}
//defer db.Close()
//
//start := time.Now()
//for db.Ping() != nil {
//if start.After(start.Add(10 * time.Second)) {
//fmt.Println("failed to connect after 10 secs.")
//break
//}
//}
//fmt.Println("connected:", db.Ping() == nil)
//_, err = db.Exec(`DROP TABLE IF EXISTS COMPANY;`)
//if err != nil {
//panic(err)
//}
//_, err = db.Exec(`CREATE TABLE COMPANY1 (ID INT PRIMARY KEY NOT NULL, NAME text);`)
//if err != nil {
//panic(err)
//}
//fmt.Println("table company is created")