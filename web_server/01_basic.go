package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Go Web Server!")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome User!!!")
}

func main() {
	var port string = fmt.Sprintf(":%d", 5050)

	http.HandleFunc("/", index)
	http.HandleFunc("/user", user)

	fmt.Printf("Listening server at: http://localhost%s", port)
	http.ListenAndServe(port, nil)
}
