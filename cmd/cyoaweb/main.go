package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Story map[string]Arc

type Arc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func main() {
	file, err := os.ReadFile("assets/gopher.json")
	if err != nil {
		fmt.Printf("Error opening story: %v", err)
		return
	}

	var story Story
	err = json.Unmarshal(file, &story)
	if err != nil {
		fmt.Printf("Error marshalling story: %v", err)
		return
	}

	// fmt.Printf("%v\n", story)
	// fmt.Printf("%+v\n", story)

	tmpl, err := template.ParseFiles("cmd/cyoaweb/template.html")
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		tmpl.Execute(response, story["intro"])
	})
	mux.HandleFunc("/new-york", func(response http.ResponseWriter, request *http.Request) {
		arc, _ := strings.CutPrefix(request.URL.Path, "/")
		tmpl.Execute(response, story[arc])
	})
	mux.HandleFunc("/debate", func(response http.ResponseWriter, request *http.Request) {
		arc, _ := strings.CutPrefix(request.URL.Path, "/")
		tmpl.Execute(response, story[arc])
	})
	mux.HandleFunc("/sean-kelly", func(response http.ResponseWriter, request *http.Request) {
		arc, _ := strings.CutPrefix(request.URL.Path, "/")
		tmpl.Execute(response, story[arc])
	})

	mux.HandleFunc("/mark-bates", func(response http.ResponseWriter, request *http.Request) {
		arc, _ := strings.CutPrefix(request.URL.Path, "/")
		tmpl.Execute(response, story[arc])
	})

	mux.HandleFunc("/denver", func(response http.ResponseWriter, request *http.Request) {
		arc, _ := strings.CutPrefix(request.URL.Path, "/")
		tmpl.Execute(response, story[arc])
	})

	mux.HandleFunc("/home", func(response http.ResponseWriter, request *http.Request) {
		arc, _ := strings.CutPrefix(request.URL.Path, "/")
		tmpl.Execute(response, story[arc])
	})

	fmt.Println("Starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
