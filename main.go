package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "api/docs"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

type User struct {
	ID    int    `json: "id"`
	Name  string `json:	"name"`
	Email string `json: "email`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlerHomePage).Methods("GET")
	router.HandleFunc("/getUsers", getUsers).Methods("GET")
	router.HandleFunc("/getUser/{id}", getUser).Methods("GET")
	router.HandleFunc("/createUser", createUser).Methods("POST")
	router.HandleFunc("/updateUser", updateUser).Methods("POST")
	router.HandleFunc("/deleteUser/{id}", deleteUser).Methods("GET")
	log.Println("Started")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handlerHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your server is fucking running")
}

func GetDB() *sql.DB {
	var err error

	log.Println("Getting DB")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "123", "testdb")
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Err while getting DB")
		panic(err)
	}

	return db
}

// @Summary Get all user
// @Description Get all user
// @ID get-all-users
// @Produce  json
// @Success 200 {object} User
// @Router /getUsers [get]
func getUsers(w http.ResponseWriter, r *http.Request) {
	var user User
	var userList []User

	db := GetDB()
	log.Println(db.Ping())
	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println(rows)

	for rows.Next() {
		log.Println(1)
		rows.Scan(&user.ID, &user.Name, &user.Email)
		userList = append(userList, user)
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
	err := db.QueryRow("SELECT * FROM users WHERE \"Id\" = $1", id).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		log.Println(err)
		if user.ID == 0 {
			fmt.Fprintln(w, "The requested user does not exist")
		} else {
			fmt.Fprintln(w, "Something's wrong")
		}
		return
	}
	defer db.Close()
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
	_, err := db.Exec("INSERT INTO public.users(\"Name\", \"Email\") VALUES ($1, $2)", newUser.Name, newUser.Email)
	defer db.Close()

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
	db := GetDB()

	json.NewDecoder(r.Body).Decode(&user)
	log.Println(user)
	_, err := db.Exec("UPDATE public.users SET \"Name\"=$1, \"Email\"=$2	WHERE \"Id\" = $3;", user.Name, user.Email, user.ID)
	defer db.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if user.ID == 0 {
			fmt.Fprintln(w, "The requested user does not exist")
		} else {
			fmt.Fprintln(w, "Something's wrong")
		}
		return
	}
	w.WriteHeader(http.StatusOK)

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
	vars := mux.Vars(r)
	id := vars["id"]
	db := GetDB()
	_, err := db.Exec("DELETE FROM public.users	WHERE \"Id\"=$1;", id)
	defer db.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something's wrong")
		return
	}
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, "Success")
}
