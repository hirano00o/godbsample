package adapter

import (
	"strconv"

	"github.com/hirano00o/godbsample/entity"
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
	age := strconv.Itoa(u.Age)
	return a.Set(map[string]string{
		"Name": name,
		"Age":  age,
	})
}

func (a *UserAdapter) Find(u entity.User) ([]entity.User, error) {
	users, err := a.Get(u.Name)
	if err != nil {
		return nil, err
	}
	ret := make([]entity.User, 0)
	for _, user := range users {
		ret = append(ret, entity.User{
			Name: string(user[0].([]byte)),
			Age:  int(user[1].(int64)),
		})
	}
	return ret, err
}
