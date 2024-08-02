package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"string"`
}

var users []User

func init() {
	users = append(users, User{ID: 1, Name: "Alice"})
	users = append(users, User{ID: 2, Name: "Bob"})
	users = append(users, User{ID: 3, Name: "Charlie"})
}

func main() {
	uh := userHandler{}
	http.Handle("/users", uh)
	http.ListenAndServe(":8080", nil)
}

type userHandler struct{}

func (uh userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsers(w, r)
	default:
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(b)
}
