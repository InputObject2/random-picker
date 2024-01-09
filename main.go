package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type ItemList struct {
	Items []string `json:"items"`
}

func randomItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var list ItemList
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(list.Items) == 0 {
		http.Error(w, "List is empty", http.StatusBadRequest)
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(list.Items))
	selectedItem := list.Items[randomIndex]

	response, err := json.Marshal(map[string]string{"selectedItem": selectedItem})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/random", randomItemHandler)

	log.Println("Starting web server on http://localhost:8080/ and api is available for a json list format on http://localhost:8080/random")
	http.ListenAndServe(":8080", nil)
}
