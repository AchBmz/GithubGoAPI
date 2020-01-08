package service

import (
	"io/ioutil"
	"log"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	languages := r.URL.Query()["language"]
	var language string
	if len(languages) == 1 {
		language = languages[0]
	}
	res, err := http.Get("https://api.github.com/search/repositories?q=language:" + language + "&sort=stars&order=desc")

	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal("Unexpected status code", res.StatusCode)
	}

	log.Printf("List of Repositories: %s\n", body)
	w.Write([]byte(body))
}

// function to allow access from the frontend server
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
