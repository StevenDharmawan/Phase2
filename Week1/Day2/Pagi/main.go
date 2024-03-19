package main

import (
	"fmt"
	"go-web-server/config"
	"go-web-server/handlers"
	"net/http"
)

func main() {
	//Basic
	//var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer, "Hello World")
	//}
	//
	//server := http.Server{
	//	Addr:    "localhost:8080",
	//	Handler: handler,
	//}
	//fmt.Println("Running server on port: 8080")
	//err := server.ListenAndServe()
	//if err != nil {
	//	return
	//}

	//ServeMux
	config.InitDB("root:@tcp(localhost:3306)/Books")
	defer config.DB.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/create/book", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateBook(w, r)
	})
	mux.HandleFunc("/get/books", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetBooks(w, r)
	})

	fmt.Println("Running server on port :8080")

	// running web server on local env
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server :", err.Error())
	}
}
