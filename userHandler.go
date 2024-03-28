package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// @Summary Get all user
// @Description Get all user
// @ID get-all-users
// @Produce  json
// @Success 200 {object} User
// @Router /getUsers [get]
func getUsers(w http.ResponseWriter, r *http.Request) {
	var userList []User

	log.Println("Init query")
	db := GetDB()
	log.Println("Finish Getting DB")
	err := db.Find(&userList).Error
	log.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(userList)
}

// @Summary Get user by ID
// @Description Get user by ID
// @ID get-user-by-id
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /getUser/{id} [get]
func getUser(w http.ResponseWriter, r *http.Request) {
	var user User
	vars := mux.Vars(r)
	id := vars["id"]

	db := GetDB()
	err := db.First(&user, id).Error

	if err != nil {
		log.Println(err)
		if user.Id == 0 {
			fmt.Fprintln(w, "The requested user does not exist")
		} else {
			fmt.Fprintln(w, "Something's wrong")
		}
		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}

// @Summary Create-new-user
// @Description Create new user base on the json passed
// @ID Create-new-user
// @Produce  json
// @Success 200 {object} User
// @Router /createUser [post]
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	db := GetDB()

	json.NewDecoder(r.Body).Decode(&newUser)
	log.Println(newUser)
	err := db.Create(&newUser).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something's wrong")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Success")
}

// @Summary Update-user
// @Description Update user info based on the json passed
// @ID Update-user
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /updateUser [post]
func updateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var oldUser User
	db := GetDB()

	json.NewDecoder(r.Body).Decode(&user)
	log.Println(user)
	err1 := db.First(&oldUser, user.Id).Error

	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Somethings's wrong upon retrieving the data")
		return
	} else {
		oldUser = user
		err := db.Save(&oldUser).Error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if user.Id == 0 {
				fmt.Fprintln(w, "The requested user does not exist")
			} else {
				fmt.Fprintln(w, "Something's wrong")
			}
			return
		}
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprint(w, "Success")
}

// @Summary Delete user by ID
// @Description Delete the user which have the id passed through the get request
// @ID Delete-user-by-id
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /deleteUser/{id} [get]
func deleteUser(w http.ResponseWriter, r *http.Request) {
	var user User
	vars := mux.Vars(r)
	id := vars["id"]
	db := GetDB()
	err := db.Delete(&user, id).Error

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something's wrong")
		return
	}
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, "Success")
}
