package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=glow password=1234 dbname=gorm port=5434 sslmode=disable" //DatabaseStriNg
var DB *gorm.DB

func DBconnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}
}

//docker run --name rrhh-postgres -e POSTGRES_USER=glow -e POSTGRES_PASSWORD=1234 -p 5434:5432 -d postgres
