package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/slogsdon/b/models"
)

var DB gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", "user=mobdfugjswmyiz password=Orj0APRPUBg6eSzF9U504Idt_X host=ec2-54-204-40-140.compute-1.amazonaws.com port=5432 dbname=d4os48l5viar5r sslmode=require")

	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	DB.AutoMigrate(models.Category{})
	DB.AutoMigrate(models.Post{})
}
