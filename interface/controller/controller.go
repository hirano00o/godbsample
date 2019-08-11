package controller

import (
	"github.com/hirano00o/godbsample/interface/adapter"
	"github.com/hirano00o/godbsample/usecase"
)

type UserController struct {
	Interactor usecase.UserUsecase
}

func NewUserController(db adapter.DB) *UserController {
	return &UserController{
		Interactor: usecase.UserUsecase{
			Adp: &adapter.UserAdapter{
				DB: db,
			},
		},
	}
}

func (u *UserController) Write(name string, age int) {
	u.Interactor.WriteUser(name, age)
}

func (u *UserController) Read(name string) {
	u.Interactor.ReadUser(name)
}
