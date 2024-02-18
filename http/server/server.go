// серверный пакет для удобного и простого запуска всех хэндлеров
package server

import (
	"Distributed_arithmetic_expression_evaluator/http/server/handler"
	"fmt"
	"net/http"
)

// запуск локалки по адресу 127.0.0.1:8123
func StartServer() {
	fmt.Println("Server is started")
	server := http.Server{
		Addr: "127.0.0.1:8123",
	}
	server.ListenAndServe()
}

func Calculator() {
	http.HandleFunc("/", handler.Calculator) // хэндлер для загрузки шаблона начальной страницы
	http.HandleFunc("/homeForm", handler.HomeForm) // хэндлер получения формы выражений
}

func TimeOfOperations() {
	http.HandleFunc("/timeOfOperations", handler.AllOperationsAndTiming) // хэндлер для загрузки шаблона страницы с операциями
	
	// хэндлеры получения времени каждого знака с формы
	http.HandleFunc("/AllOperationFormPlus", handler.OperationFormPlus)
	http.HandleFunc("/AllOperationFormMinus", handler.OperationFormMinus)
	http.HandleFunc("/AllOperationFormMult", handler.OperationFormMult)
	http.HandleFunc("/AllOperationFormDivision", handler.OperationFormDivision)
}

func ResultPage() {
	http.HandleFunc("/Results", handler.ResultsPage) // хэндлер для загрузки шаблона страницы с результатами
	http.HandleFunc("/Search", handler.Searcher) // реализация поиска
	http.HandleFunc("/SearchReset", handler.SearchReset) // реализация сброса поиска
}