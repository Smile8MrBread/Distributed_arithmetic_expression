// Папка с хэндлерами
package handler

import (
	"net/http"
	"fmt"
	"regexp"
)

// /home
func PostMainHandler(w http.ResponseWriter, r *http.Request) {
	// Ввод выражения
	fmt.Println("Введите выражение: ")
	str := ""
	fmt.Scanln(&str)

	//Проверка на правильность
	limit := regexp.MustCompile(`^[0-9]+(?:[\+\-\*/][0-9]+)*$`)
	if !limit.MatchString(str) {
		fmt.Println("Введенно неверное выражение")
	} else {
		fmt.Println("Ваше выражение принято к рассмотрению")
	}
}

func AllExpressions(w http.ResponseWriter, r *http.Request) {
	//...
}

func AllOperationsAndTiming(w http.ResponseWriter, r *http.Request) {
	//...
}

