package database

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/hirano00o/godbsample/interface/adapter"
)

type Server struct {
	Conn *sql.DB
}

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewDB(cnf Config) adapter.DB {
	dbconf := mysql.Config{
		User:   cnf.User,
		Passwd: cnf.Password,
		Net:    "tcp",
		Addr:   cnf.Host + ":" + cnf.Port,
		DBName: cnf.DBName,
	}
	db, err := sql.Open("mysql", dbconf.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	handler := new(Server)
	handler.Conn = db
	return handler
}

func (s *Server) Set(m map[string]string) error {
	tx, err := s.Conn.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println(r)
		} else if err != nil {
			tx.Rollback()
			log.Println("Database Rollback: " + err.Error())
		} else {
			err = tx.Commit()
			log.Println("Commit")
		}
	}()

	stmt, err := tx.Prepare("INSERT INTO USERS (NAME, AGE) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	age, err := strconv.Atoi(m["Age"])
	if err != nil {
		return err
	}
	_, err = stmt.Exec(m["Name"], age)

	return err
}

func (s *Server) Get(name string) ([][]interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	rows, err := s.Conn.Query("SELECT NAME, AGE FROM USERS WHERE NAME = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	valuePtr := make([]interface{}, len(columns))
	ret := make([][]interface{}, 0)

	for rows.Next() {
		values := make([]interface{}, len(valuePtr))
		for i := range columns {
			valuePtr[i] = &values[i]
		}
		rows.Scan(valuePtr...)

		ret = append(ret, values)
	}
	return ret, nil
}
