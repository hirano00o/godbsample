package controller

import (
	"github.com/hirano00o/godbsample/infrastructure/database"
	"github.com/hirano00o/godbsample/interface/adapter"
)

type UserController struct {
	Interactor usecase.UserUsecase
}

func NewUserController(s database.Server) *UserController {
	return &UserController{
		Interactor: usecase.UserUsecase{
			adp: &adapter.UserAdapter{
				DB: s,
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
