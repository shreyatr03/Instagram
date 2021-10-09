package main

import (
	"net/http"

	"github.com/shreyatr03/Instagram/routes"
)

func main() {
	http.HandleFunc("/users", routes.UsersHandler)
	http.HandleFunc("/users/", routes.UserHandler)
	http.HandleFunc("/posts", routes.PostsHandler)
	http.HandleFunc("/posts/", routes.PostHandler)
	http.HandleFunc("/posts/users/", routes.UserPostsHandler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
