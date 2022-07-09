package uow

import (
	"github.com/mv-kan/unit-of-work/repo"
	"gorm.io/gorm"
)

type UnitOfWorkStore interface {
	Users() repo.UserRepo
	Todos() repo.TodoRepo
}

type uowStore struct {
	todos repo.TodoRepo
	users repo.UserRepo
}

func (store *uowStore) Users() repo.UserRepo {
	return store.users
}
func (store *uowStore) Todos() repo.TodoRepo {
	return store.todos
}

type UnitOfWorkBlock func(store UnitOfWorkStore) error

type UnitOfWork interface {
	Do(fn UnitOfWorkBlock) error
}

func New(db *gorm.DB) UnitOfWork {
	return &unitOfWork{db: db}
}

type unitOfWork struct {
	db *gorm.DB
}

func (uow *unitOfWork) Do(fn UnitOfWorkBlock) error {
	return uow.db.Transaction(func(tx *gorm.DB) error {
		newStore := uowStore{
			todos: repo.NewTodoRepo(tx),
			users: repo.NewUserRepo(tx),
		}
		return fn(&newStore)
	})
}
