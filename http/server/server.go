package server

import (
	"Distributed_arithmetic_expression_evaluator/http/server/handler"
	"fmt"
	"net/http"
)

func StartServer() {
	fmt.Println("Server is starting")
	server := http.Server{
		Addr: "127.0.0.1:8123",
	}
	server.ListenAndServe()
}

func Home() {
	http.HandleFunc("/", handler.HomeHandler)
}

func PostHome() {
	http.HandleFunc("/home", handler.MainForm)
}

func GetTimings() {
	http.HandleFunc("/timing", handler.AllExpressionsTime)
}

func GetExpressions() {
}