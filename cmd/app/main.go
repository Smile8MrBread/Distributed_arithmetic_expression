// точка входа в программу
package main

import "Distributed_arithmetic_expression_evaluator/http/server"

func main() {
	server.Calculator() 
	server.TimeOfOperations()
	server.ResultPage()

	server.StartServer()
}