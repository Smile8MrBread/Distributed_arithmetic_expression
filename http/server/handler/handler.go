// Папка с хэндлерами
package handler

import (
	"Distributed_arithmetic_expression_evaluator/database"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"

	"github.com/google/uuid"
)

// /home
func MainForm(w http.ResponseWriter, r *http.Request) {
	//Получаем данные из формы
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Что-то не так в бэке")
		w.WriteHeader(500)
	}

	form := r.FormValue("expression")

	//Проверка на правильность
	limit := regexp.MustCompile(`^[0-9]+(?:[\+\-\*/][0-9]+)*$`)

	// Если нет - код 400
	if !limit.MatchString(form) {
		w.WriteHeader(400)
		fmt.Println("Введенно неверное выражение:", form)

		// Если да - добавляем в мапу, код 200
	} else {
		w.WriteHeader(200)
		fmt.Println("Ваше выражение принято к рассмотрению")
		database.SetDB(uuid.New().String()[:9], form) // uuid делает id с ограничением длинны в 8 символов
	}
}

func AllExpressionsTime(w http.ResponseWriter, r *http.Request) {
	//создаем html-шаблон
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	tmpl, err := template.ParseFiles(wd + "\\templates\\AllTimings.tmpl")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	} else {
		w.WriteHeader(200)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Distributed_arithmetic_expression_evaluator\\html\\homePage.html"))

	tmpl.Execute(w, nil)
}

func AllOperationsAndTiming(w http.ResponseWriter, r *http.Request) {
	//
}
