package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=kiss password=kiss dbname=sum_go port=5432"
var DB *gorm.DB

func DBConection()  {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN),&gorm.Config{})
		if error != nil {
			log.Fatal(error)
		}else{
			log.Println("DB Connected")
		}
	
}