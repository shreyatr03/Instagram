package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/users", routes.usersHandlers)
	http.HandleFunc("/users/{id}", routes.userHandlers)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
