package main

import (
	"fmt"
	"log"
	"net/http"

	_ "api/docs"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type User struct {
	Id    int    `column: "id"`
	Name  string `column: "name"`
	Email string `column: "email"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlerHomePage).Methods("GET")
	router.HandleFunc("/getUsers", getUsers).Methods("GET")
	router.HandleFunc("/getUser/{id}", getUser).Methods("GET")
	router.HandleFunc("/createUser", createUser).Methods("POST")
	router.HandleFunc("/updateUser", updateUser).Methods("POST")
	router.HandleFunc("/deleteUser/{id}", deleteUser).Methods("GET")

	router.HandleFunc("/getTodoByUser/{idUser}", getTodoByUserId).Methods("GET")
	router.HandleFunc("/createTodo", createTodo).Methods("POST")
	router.HandleFunc("/updateTodo", updateTodo).Methods("POST")
	router.HandleFunc("/deleteTodo/{id}", deleteTodo).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	handler := c.Handler(router)
	log.Println("Started")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func handlerHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your server is fucking running")
}
