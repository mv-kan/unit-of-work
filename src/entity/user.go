package entity

type User struct {
	Base
	Username string
	Todos    []Todo //`gorm:"foreignKey:TodoID"`
}
