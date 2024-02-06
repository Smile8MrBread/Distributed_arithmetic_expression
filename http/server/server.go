package server

import (
	"net/http"
	"github.com/Smile8MrBread/Distributed_arithmetic_expression/http/server/handler"
)

func StartServer(port string) {
	server := http.Server{
		Addr: "127.0.0.1:8123",
	}
	server.ListenAndServe()
}

func PostHome() {
	http.HandleFunc("/home", handler.PostMainHandler)
}