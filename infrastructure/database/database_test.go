package database

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestOKNewDB(t *testing.T) {
	conf := Config{
		User:     "testuser",
		Password: "password",
		Host:     "localhost",
		Port:     "3306",
		DBName:   "USER",
	}
	db := NewDB(conf)
	if db == nil {
		t.Errorf("was expecting an error, but there was none")
	}
}

func TestOKSet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO USERS").ExpectExec().WithArgs("Bob", 10).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	s := new(Server)
	s.Conn = db
	m := map[string]string{
		"Name": "Bob",
		"Age":  "10",
	}
	if err := s.Set(m); err != nil {
		t.Errorf("error was not expected while insert stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestNGSetBegin(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s := new(Server)
	s.Conn = db
	err = s.Set(nil)
	if err == nil {
		t.Errorf("was expecting an error, but there was none")
	}
	if strings.Contains(err.Error(), "Begin") == false {
		t.Errorf("was not expecting an error, but there was Begin")
	}
}

func TestNGSetStmt(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()

	s := new(Server)
	s.Conn = db
	err = s.Set(nil)
	if err == nil {
		t.Errorf("was expecting an error, but there was none")
	}
	if strings.Contains(err.Error(), "Prepare") == false {
		t.Errorf("was not expecting an error, but there was Prepare")
	}
}

func TestNGSetStrconv(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO USERS").ExpectExec().WithArgs("Bob").WillReturnError(fmt.Errorf("error strconv.Atoi"))
	mock.ExpectRollback()

	s := new(Server)
	s.Conn = db
	m := map[string]string{
		"Name": "Bob",
	}
	err = s.Set(m)
	if err == nil {
		t.Errorf("was expecting an error, but there was none")
	}
	if strings.Contains(err.Error(), "strconv.Atoi") == false {
		t.Errorf("was not expecting an error, but there was strconv.Atoi")
	}
}

func TestOKGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT NAME, AGE FROM USERS WHERE NAME = ?`),
	).
		WithArgs("Bob").
		WillReturnRows(sqlmock.NewRows([]string{"ID", "NAME", "AGE"}).
			AddRow(1, "Bob", 15),
		)

	s := new(Server)
	s.Conn = db
	res, err := s.Get("Bob")
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}

	if len(res) != 1 {
		t.Errorf("was not expecting count %d, but there was count 1", len(res))
	}

	id := res[0][0].(int64)
	if id != 1 {
		t.Errorf("was not expecting id %d, but there was id 1", id)
	}

	name := res[0][1].(string)
	if name != "Bob" {
		t.Errorf("was not expecting name %s, but there was name 'Bob'", name)
	}

	age := res[0][2].(int64)
	if age != 15 {
		t.Errorf("was not expecting age %d, but there was age 15", age)
	}
}

func TestNGGetConn(t *testing.T) {
	_, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT NAME, AGE FROM USERS WHERE NAME = ?`),
	).
		WithArgs("Bob").
		WillReturnRows(sqlmock.NewRows([]string{"ID", "NAME", "AGE"}).
			AddRow(1, "Bob", 15),
		)

	s := new(Server)
	_, err = s.Get("Bob")
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
}
