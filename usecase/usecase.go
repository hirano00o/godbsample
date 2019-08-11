package usecase

import (
	"log"
	"strconv"

	"github.com/hirano00o/godbsample/entity"
)

type UserUsecase struct {
	Adp UserAdapter
}

type UserAdapter interface {
	Store(entity.User) error
	Find(entity.User) ([]entity.User, error)
}

func (u *UserUsecase) WriteUser(name string, age int) {
	err := u.Adp.Store(entity.User{
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
	users, err := u.Adp.Find(*user)
	if err != nil {
		log.Println("Read Error: " + err.Error())
	}
	log.Println("Name: " + name + " found " + string(len(users)) + " users.")
	for _, us := range users {
		log.Println("User: " + us.Name + "(" + strconv.Itoa(us.Age) + ")")
	}
}
