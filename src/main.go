package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/mv-kan/unit-of-work/entity"
	"github.com/mv-kan/unit-of-work/uow"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=db user=postgres password=secret dbname=postgres port=5432 sslmode=disable"
	//dsn := "host=localhost user=postgres password=secret dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// auto migrate to create db scheme
	db.AutoMigrate(&entity.User{}, &entity.Todo{}, &entity.Sub{})

	unitOfWork := uow.New(db)
	// this function creates user and then creates to user first greeting todo
	// and all of that is in one transaction
	createUserAndAddGreetingTodo := func(store uow.UnitOfWorkStore) error {
		username := "cool user " + uuid.NewString() // to create uniqueness for username

		user := entity.User{Base: entity.Base{ID: uuid.New()}, Username: username}

		user, err := store.Users().Create(user)
		if err != nil {
			return err
		}
		fmt.Printf("User 'username=%s' has been created, id=%s\n", username, user.ID.String())

		todo := entity.Todo{Base: entity.Base{ID: uuid.New()}, UserID: user.ID, Name: "Greetings" + username, Subs: []entity.Sub{{Base: entity.Base{ID: uuid.New()}, Name: "Sub1"}, {Base: entity.Base{ID: uuid.New()}, Name: "Sub2"}}}
		todo, err = store.Todos().Create(todo)
		if err != nil {
			return err
		}
		fmt.Printf("Todo 'id=%s' has been created, todo.UserID=%s\n", todo.ID.String(), todo.UserID.String())
		return nil
	}
	unitOfWork.Do(createUserAndAddGreetingTodo)
	return
	// // Below is the code that I used for testing todo and user repos
	// // USER REPOSITORY
	// // CREATE
	// todoRepo := repo.NewTodoRepo(db)
	// userRepo := repo.NewUserRepo(db)
	// user, err = userRepo.Create(user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("User created: ", user)
	// // READ
	// // todo := entity.Todo{Base: entity.Base{ID: uuid.MustParse("451bffc6-207f-4e86-8e29-ee8cb6254a8e")}}
	// userNew, err := userRepo.Read(user.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("User readed: ", userNew)
	// // UPDATE
	// // UPDATE
	// userNew.Todos = append(userNew.Todos, entity.Todo{Base: entity.Base{ID: uuid.New()}, Name: "Added updated todo user repo"})
	// err = userRepo.Update(userNew)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// userNew, err = userRepo.Read(user.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("User updated: ", userNew)
	// // DELETE
	// err = userRepo.Delete(user.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("User Deleted: without errors")
	// userNew, err = userRepo.Read(user.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("User readed: ", userNew)
	// return
	// // TODO REPOSITORY
	// // CREATE
	// todo, err = todoRepo.Create(todo)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Todo created: ", todo)
	// // READ
	// // todo := entity.Todo{Base: entity.Base{ID: uuid.MustParse("451bffc6-207f-4e86-8e29-ee8cb6254a8e")}}
	// todoNew, err := todoRepo.Read(todo.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Todo readed: ", todoNew)
	// // UPDATE
	// todoNew.Name = "Update golang"
	// todoNew.Subs = append(todoNew.Subs, entity.Sub{Base: entity.Base{ID: uuid.New()}, Name: "Added updated sub"})
	// todoNew.Subs[0].Name = "Updated updated sub"
	// err = todoRepo.Update(todoNew)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// todoNew, err = todoRepo.Read(todo.ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Todo updated: ", todoNew)
	// // DELETE
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
