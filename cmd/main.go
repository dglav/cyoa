package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type cyoa struct {
	Intro     Arc `json:"intro"`
	NewYork   Arc `json:"new-york"`
	Debate    Arc `json:"debate"`
	SeanKelly Arc `json:"sean-kelly"`
	MarkBates Arc `json:"mark-bates"`
	Denver    Arc `json:"denver"`
	Home      Arc `json:"home"`
}

func main() {
	file, err := os.ReadFile("assets/gopher.json")
	if err != nil {
		fmt.Printf("Error opening story: %v", err)
		return
	}

	var story cyoa
	err = json.Unmarshal(file, &story)
	if err != nil {
		fmt.Printf("Error marshalling story: %v", err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(200)
		response.Header().Add("Content-Type", "text/html; charset=utf-8")
	})

	fmt.Println("Starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
