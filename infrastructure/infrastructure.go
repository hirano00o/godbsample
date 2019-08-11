package infrastructure

import (
	"os"

	"github.com/hirano00o/godbsample/infrastructure/database"
	"github.com/hirano00o/godbsample/interface/controller"
)

type user struct {
	name string
	age  int
}

func Route() {
	conf := database.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}

	c := controller.NewUserController(database.NewDB(conf))
	users := []user{
		user{"Alice", 10},
		user{"Bob", 15},
		user{"Carol", 20},
		user{"Dave", 25},
		user{"Ellen", 30},
		user{"Frank", 35},
		user{"Bobby", 18},
	}

	for _, u := range users {
		c.Write(u.name, u.age)
	}
	for _, u := range users {
		c.Read(u.name)
	}
}
