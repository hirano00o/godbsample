package adapter

import (
	"strconv"

	"github.com/hirano00o/godbsample/entity/entity"
)

type DB interface {
	Set(map[string]string) error
	Get(string) ([][]interface{}, error)
}

type UserAdapter struct {
	DB
}

func (a *UserAdapter) Store(u entity.User) error {
	name := u.Name
	age, err := strconv.Itoa(u.Age)
	if err != nil {
		return err
	}
	return a.Set(map[string]string{
		"Name": name,
		"Age":  age,
	})
}

func (a *Adapter) Find(u entity.User) ([]entity.User, error) {
	users, err := a.Get(u.Name)
	if err != nil {
		return nil, err
	}
	ret := make([]entity.User)
	for i, user := range users {
		for _, c := range user {
			ret = append(entity.User{
				Name: c[0].(string),
				Age:  c[1].(int),
			})
		}
	}
	return ret, err
}
