package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type todoItem struct {
	Id         int       `gorm:"primaryKey"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	IdUser     int       `column:"id_user"`
	RemindTime time.Time `column:"remind_time"`
}

func getTodoByUserId(w http.ResponseWriter, r *http.Request) {
	var todoItem []todoItem
	vars := mux.Vars(r)
	id := vars["idUser"]
	log.Println("test1")
	db := GetDB()
	err := db.Where("\"id_user\" = ?", id).Find(&todoItem).Error

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	json.NewEncoder(w).Encode(todoItem)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todoItem todoItem
	log.Println(todoItem)
	db := GetDB()

	err2 := json.NewDecoder(r.Body).Decode(&todoItem)
	if err2 != nil {
		log.Println(err2)
	}
	log.Println(todoItem)
	err := db.Create(&todoItem).Error

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success")
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	var todoItem todoItem
	vars := mux.Vars(r)
	id := vars["id"]
	db := GetDB()

	err := db.Delete(&todoItem, id).Error

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success")
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodoItem todoItem
	var todoItem todoItem
	db := GetDB()

	json.NewDecoder(r.Body).Decode(&newTodoItem)
	err := db.Find(&todoItem, newTodoItem.Id).Error

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something went wrong")
		return
	} else {
		todoItem = newTodoItem
		log.Println(newTodoItem)
		err1 := db.Save(&todoItem).Error
		if err1 != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Something went wrong")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Success")
	}
}
