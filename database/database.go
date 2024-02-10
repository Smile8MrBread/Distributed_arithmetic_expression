package database

import (
	"errors"
)

type DataBase struct {
	id 		   string
	expression string
	value 	   int
}

var db = []DataBase{}

func SetDB(id, expression string) {
	db = append(db, DataBase{id: id, expression: expression, value: 0})
}

func GetDB(id string) (int, error) {
	for i := range db {
		if db[i].id == id {
			return db[i].value, nil
		}
	}

	return 0, errors.New("id не найден")
}