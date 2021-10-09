package main

import (
	"net/http"

	"github.com/shreyatr03/Instagram/routes"
)

func main() {
	http.HandleFunc("/users", routes.UsersHandlers)
	http.HandleFunc("/users/:id", routes.UserHandlers)
	http.HandleFunc("/posts", routes.PostsHandlers)
	// http.HandleFunc("/users/{id}", routes.UserHandlers)
	// http.HandleFunc("/users/{id}", routes.UserHandlers)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
