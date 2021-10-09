package main

import (
	"net/http"

	"github.com/shreyatr03/Instagram/routes"
)

func main() {
	http.HandleFunc("/users", routes.UsersHandlers)
	http.HandleFunc("/users/", routes.UserHandlers)
	http.HandleFunc("/posts", routes.PostsHandlers)
	http.HandleFunc("/posts/", routes.PostHandlers)
	// http.HandleFunc("/users/{id}", routes.UserHandlers)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
