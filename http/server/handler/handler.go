// пакет реализации всех хэндлеров
package handler

import (
	"Distributed_arithmetic_expression_evaluator/database"
	"Distributed_arithmetic_expression_evaluator/database/operationTime"
	"Distributed_arithmetic_expression_evaluator/logic"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func Calculator(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/homePage.html")) // шаблон
	
	data := database.GetDB()
	err := tmpl.Execute(w, data) // заполнение шаблона 
	if err != nil {
		fmt.Println(err, "ошибка формы")
	}
}

func HomeForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err, "парс формы")
	}
	form := r.FormValue("expression-input") // получение выражения из формы

	// регулярка, по которой определяется - выражение введено правильно или с ошибкой
	// не реализована запись со скобками
	limit := regexp.MustCompile(`^-?[0-9]+(?:[\+\-\*/][0-9]+)*$`)
	id := uuid.New().String()[:8] // генерация id

	if !limit.MatchString(form) {
		fmt.Println("Введенно неверное выражение:", form)

		// добавление в массив с флагом false
		database.SetDB(id, form, time.Second * 0, false)
	} else {
		fmt.Println("Ваше выражение принято к рассмотрению")

		logic.ExecutionTime(id, form) // подробнее в пакете logic
		timeEnd := logic.GetExecTime(id) // подробнее в пакете logic
		database.SetDB(id, form, timeEnd, true) // добавление в массив с флагом true

		// go рутинка, которая ждет пока пройдет время для счета выражения,
		// а после утановит статус о готовности
		go func(id string) {
			for {
				if logic.GetExecTime(id) <= 0 {
					database.IsReadySet(id)

					break
				}
			}
		}(id)
	}

	// редирект на ту же страницу, т.к post форма перенаправляет на страницу, 
	// которая записана в поле actions
	http.Redirect(w, r, "http://127.0.0.1:8123/", http.StatusSeeOther)
}

func AllOperationsAndTiming(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/AllTimings.html")) // шаблон
	
	data := operationTime.GetOperators()
	err := tmpl.Execute(w, data) // заполнение шаблона
	if err != nil {
		fmt.Println(err, "загрузка шаблона")
	}
}

func OperationFormPlus(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err, "парс формы")
	}

	plus := r.FormValue("plus-time") // получение значения для "+"
	plusN, _ := strconv.Atoi(plus) // перевод в int
	// запись нового значения для "+" в массив
	operationTime.SetTimePlus(time.Duration(plusN) * time.Second)

	// редирект
	http.Redirect(w, r, "http://127.0.0.1:8123/timeOfOperations", http.StatusSeeOther)
}

// по аналогии
func OperationFormMinus(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err, "парс формы 2")
	}

	minus := r.FormValue("minus-time")
	minusN, _ := strconv.Atoi(minus)
	operationTime.SetTimeMinus(time.Duration(minusN) * time.Second)

	http.Redirect(w, r, "http://127.0.0.1:8123/timeOfOperations", http.StatusSeeOther)
}

// по аналогии
func OperationFormMult(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err, "парс формы 3")
	}

	mult := r.FormValue("multiplication-time")
	multN, _ := strconv.Atoi(mult)
	operationTime.SetTimeMult(time.Duration(multN) * time.Second)

	http.Redirect(w, r, "http://127.0.0.1:8123/timeOfOperations", http.StatusSeeOther)
}

// по аналогии
func OperationFormDivision(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err, "парс формы 4")
	}

	division := r.FormValue("division-time")
	divisionN, _ := strconv.Atoi(division)
	operationTime.SetTimeDivision(time.Duration(divisionN) * time.Second)

	http.Redirect(w, r, "http://127.0.0.1:8123/timeOfOperations", http.StatusSeeOther)
}

func ResultsPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/ResultsPage.html")) // шаблон
	
	data := database.GetDB()
	err := tmpl.Execute(w, data) // заполнение шаблона
	if err != nil {
		fmt.Println(err, "загрузка шаблона 2")
	}
}

func Searcher(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err, "парс формы 5")
	}

	s := r.FormValue("search-input") // получение id для поиска выражения


	data := database.GetDB()
	// долгий поиск введеного id
	for i := range data.Datas {
		if data.Datas[i].Id == s {
			
			// перевод в состояние "поиска" и добавление в поисковой массив
			database.IsInSearchSet(data.Datas[i]) 
			break
		}
	}
	
	// редирект
	http.Redirect(w, r, "http://127.0.0.1:8123/Results", http.StatusSeeOther)
}

func SearchReset(w http.ResponseWriter, r *http.Request) {
	database.IsNotInSearch() // сброс состояния "поиска"

	// редирект
	http.Redirect(w, r, "http://127.0.0.1:8123/Results", http.StatusSeeOther)
}