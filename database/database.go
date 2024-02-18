// пакет с импровизированной БД для хранения всей основной инфы
// о выражениях, которая используется в динамических html шаблонах 
// и не только
package database

import (
	"fmt"
	"time"

	"github.com/Knetic/govaluate"
)

// основная структура БД
type DataBase struct {
	InSearching bool // флажок, обозначающий поиск на странице Results
	Datas   	[]Data // массив всех выражений
	Search		[]Data // массив поиска
}

// структура каждого выражения в БД
type Data struct {
	Id         string
	Expression string // само выражение
	Time       time.Duration // время, которое потребуется для решения
	Flag       bool // флажок для html, содержит ли выражение ошибку
	IsReady    bool // флажок готовности для смены состояния
	Result	   float64 // результат выражения
}

// местный экземпляр ДБ
var db = DataBase { 
	Datas: []Data{},
	Search: []Data{},
}

// получение экземпляра
func GetDB() DataBase {
	return db
}

// запись нового выражения
func SetDB(id, expression string, time time.Duration, flag bool) {
	// если верно
	if flag {
		// блок, переводящий строку в выражения и считающий его значение
		express, err := govaluate.NewEvaluableExpression(expression)
		if err != nil {
			fmt.Println(err, "перевод из строки в вырвжение")
		}
		
		result, err := express.Evaluate(nil)
		if err != nil {
			fmt.Println(err, "решение выражения")
		}

		// запись result, если флаг true
		db.Datas = append(db.Datas, Data{Id: id, Expression: expression,
			Time: time, Flag: flag, IsReady: false, Result: result.(float64),})
	} else {
		db.Datas = append(db.Datas, Data{Id: id, Expression: expression,
			Time: time, Flag: flag, IsReady: false})
	}
}

// go рутинка для установки значения о готовности нужному выражению
func IsReadySet(id string) {
	go func() {
		// понимаю, что данный перебор долгий, но решил не заморачиваться
		for i := range db.Datas {
				if db.Datas[i].Id == id {
					db.Datas[i].IsReady = true
				}
			}
	}()
}

// добавление в массив поиска выражения и перевод состояния страницы в "поисковой"
func IsInSearchSet(d Data) {
	db.InSearching = true
	db.Search = append(db.Search, d)
}

// вывод из "поискового" состояния и опустошение массива
func IsNotInSearch() {
	db.InSearching = false
	db.Search = []Data{}
}