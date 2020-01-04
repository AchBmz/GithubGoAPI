package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// the repository owner
type Owner struct {
	Login string
}

// infos about repository
type Item struct {
	Name            string
	Owner           Owner
	Description     string
	HTMLURL         string `json:"html_url"`
	StargazersCount int    `json:"stargazers_count"`
}

// JSONData contains the GitHub API response
type JSONData struct {
	Items []Item
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			fmt.Println(err)
		}

		data := JSONData{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(data)

		c, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		// c is now a []byte containing the encoded JSON
		//fmt.Print(string(c))
		w.Write([]byte(c)) //c
	})

	erro := http.ListenAndServe(":3000", nil)

	if erro != nil {
		log.Fatal(erro)
	}

}
