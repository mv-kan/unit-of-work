package entity

type Todo struct {
	Base
	Name string
	Subs []Sub
}
