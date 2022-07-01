package main

import (
	"fmt"
	"log"

	"github.com/mv-kan/unit-of-work/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgrespw dbname=postgres port=49153 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	todo := entity.Todo{Name: "Hello golang"}
	result := db.Create(&todo)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	fmt.Println(fmt.Sprintf("Todo: id=%s name=%s", todo.ID, todo.Name))

}
