package main

import (
	"github.com/AchBmz/GithubGoAPI.git/service"
	"log"
	"net/http"
)

//main function is the entry of the ptoject
func main() {
	http.HandleFunc("/", service.ServeHTTP)
	error := http.ListenAndServe(":3000", nil) //app to be started on port 3000
	if error != nil {
		log.Fatal(error)
	}
}
