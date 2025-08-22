package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

var userCache = make(map[int]User)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	mux.HandleFunc("POST /users", createUser)

	fmt.Println("Server listening to :8080")
	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	if user.Name == "" {
		http.Error(
			w,
			"name is required",
			http.StatusBadRequest,
		)
		return
	}

	// not thread safe, but we only make it for testing purposes
	userCache[len(userCache)+1] = user

	w.WriteHeader(http.StatusNoContent)
}