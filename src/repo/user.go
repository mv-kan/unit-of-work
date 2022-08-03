package repo

import (
	"github.com/google/uuid"
	"github.com/mv-kan/unit-of-work/entity"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(entity.User) (entity.User, error)
	Read(uuid.UUID) (entity.User, error)
	Update(entity.User) error
	Delete(uuid.UUID) error
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db}
}

type userRepo struct {
	db *gorm.DB
}

func (repo *userRepo) Create(user entity.User) (entity.User, error) {
	result := repo.db.Create(&user)
	return user, result.Error
}

func (repo *userRepo) Read(id uuid.UUID) (entity.User, error) {
	var user entity.User
	// result := repo.db.First(&user, id)
	result := repo.db.Model(&user).Preload("Todos").Find(&user)
	for i, todo := range user.Todos {
		repo.db.Model(&todo).Preload("Subs")
		user.Todos[i] = todo
	}
	return user, result.Error
}

func (repo *userRepo) Update(user entity.User) error {
	result := repo.db.Model(&user).Updates(user)
	return result.Error
}

func (repo *userRepo) Delete(id uuid.UUID) error {
	var user = entity.User{Base: entity.Base{ID: id}}
	result := repo.db.Delete(&user, id)
	return result.Error
}
