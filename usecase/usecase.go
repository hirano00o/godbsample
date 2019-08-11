package usecase

import (
	"log"

	"github.com/hirano00o/godbsample/entity/entity"
)

type UserUsecase struct {
	adp UserAdapter
}

type UserAdapter interface {
	Store(entity.User) error
	Find(entity.User) ([]entity.User, error)
}

func (u *UserUsecase) WriteUser(name string, age int) {
	err := u.adp.Store(&entity.User{
		Name: name,
		Age:  age,
	})
	if err != nil {
		log.Println("Write Error: " + err.Error())
	}
	log.Println("Write User: " + name)
}

func (u *UserUsecase) ReadUser(name string) {
	user := new(entity.User)
	user.Name = name
	users := u.adp.Find(user)
	if err != nil {
		log.Println("Read Error: " + err.Error())
	}
	log.Println("Name: " + name + " found " + string(len(users)) + " users.")
	for _, us := range users {
		var s string
		for _, v := range us {
			s = s + string(v) + " "
		}
		log.Println("User: " + s)
	}
}
