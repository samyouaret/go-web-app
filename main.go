package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

const (
	port     = 5432
	host     = "db"
	user     = "postgres"
	password = "password1234"
	dbname   = "go_app"
)

type User struct {
	// fields must be capitalized to be public
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	_, err := connectDB()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/users", userHandler)
	fmt.Println("http: ready")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	for name, headerList := range r.Header {
		fmt.Printf("%v %v\n", name, strings.Join(headerList, ", "))
	}
	fmt.Printf("%q", r.Method)
	w.Write([]byte("hello there"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsers(w, r)
		// case "POST":
		// createUser()
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{{
		Id:       1,
		Username: "samy",
		Email:    "sam@mail.com",
	}, {Id: 2,
		Username: "adam",
		Email:    "adam@mail.com",
	}}
	json, _ := json.Marshal(users)
	// you cannot write headers after calling this method
	// w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	w.Write(json)
}

func connectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	return db, err
}
