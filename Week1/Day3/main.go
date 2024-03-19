package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-routing/config"
	"go-routing/handlers"
	"net/http"
)

func main() {
	config.InitDB("root:@tcp(localhost:3306)/books?parseTime=true")
	defer config.DB.Close()

	router := httprouter.New()
	router.POST("/book/create", handlers.Create)
	router.GET("/books", handlers.GetAll)
	router.GET("/book/:id", handlers.GetBookByID)
	router.PUT("/book/:id", handlers.Update)
	router.DELETE("/book/:id", handlers.Delete)

	fmt.Println("Running server on port :8080")

	// running web server on local env
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server :", err.Error())
	}
}
