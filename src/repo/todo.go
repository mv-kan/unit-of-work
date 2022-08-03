package repo

import (
	"github.com/google/uuid"
	"github.com/mv-kan/unit-of-work/entity"
	"gorm.io/gorm"
)

type TodoRepo interface {
	Create(entity.Todo) (entity.Todo, error)
	Read(uuid.UUID) (entity.Todo, error)
	Update(entity.Todo) error
	Delete(uuid.UUID) error
}

func NewTodoRepo(db *gorm.DB) TodoRepo {
	return &todoRepo{db: db}
}

type todoRepo struct {
	db *gorm.DB
}

func (repo *todoRepo) Create(todo entity.Todo) (entity.Todo, error) {
	result := repo.db.Create(&todo)
	return todo, result.Error
}

func (repo *todoRepo) Read(id uuid.UUID) (entity.Todo, error) {
	var todo entity.Todo
	// result := repo.db.First(&todo, id)
	result := repo.db.Model(&todo).Preload("Subs").Find(&todo)
	return todo, result.Error
}

func (repo *todoRepo) Update(todo entity.Todo) error {
	result := repo.db.Model(&todo).Updates(todo)
	return result.Error
}

func (repo *todoRepo) Delete(id uuid.UUID) error {
	var todo = entity.Todo{Base: entity.Base{ID: id}}
	result := repo.db.Delete(&todo, id)
	return result.Error
}
