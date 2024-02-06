package handler

import (
	"net/http"
	"fmt"
	"regexp"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Введите выражение: ")
	str := ""
	fmt.Scanln(&str)

	//Проверка на правильность
	limit := regexp.MustCompile(`^[0-9]+(?:[\+\-\*/][0-9]+)*$`)
	if !limit.MatchString(str) {
		fmt.Println("Введенно неверное выражение")
	}
} 