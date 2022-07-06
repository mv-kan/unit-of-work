package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/mv-kan/unit-of-work/entity"
	"github.com/mv-kan/unit-of-work/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgrespw dbname=postgres port=49153 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	todoRepo := repo.NewTodoRepo(db)
	// todo := entity.Todo{Base: entity.Base{ID: uuid.New()}, Name: "Hello golang", Subs: []entity.Sub{{Base: entity.Base{ID: uuid.New()}, Name: "Sub1"}, {Base: entity.Base{ID: uuid.New()}, Name: "Sub2"}}}
	// todo, err = todoRepo.Create(todo)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Todo created: ", todo)
	todo := entity.Todo{Base: entity.Base{ID: uuid.MustParse("451bffc6-207f-4e86-8e29-ee8cb6254a8e")}}
	todoNew, err := todoRepo.Read(todo.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Todo readed: ", todoNew)

	todoNew.Name = "Update golang"
	err = todoRepo.Update(todoNew)
	if err != nil {
		log.Fatal(err)
	}

	todoNew, err = todoRepo.Read(todo.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Todo updated: ", todoNew)

	// err = todoRepo.Delete(todo.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// todoNew, err = todoRepo.Read(todo.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Todo readed: ", todoNew)

	// result := db.Create(&todo)
	// fmt.Println(result.Error)
	// fmt.Println(result.RowsAffected)
	// fmt.Println(fmt.Sprintf("Todo: id=%s name=%s", todo.ID, todo.Name))
}
