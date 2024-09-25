package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Arc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []struct {
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
		// Create a template and fill in the values with the json data
		tmpl, err := template.ParseFiles("cmd/template.html")
		if err != nil {
			http.Error(response, "Error parsing template", 500)
		}

		tmpl.Execute(response, story.Intro)
	})
	mux.HandleFunc("/new-york", func(response http.ResponseWriter, request *http.Request) {
		// Create a template and fill in the values with the json data
		tmpl, err := template.ParseFiles("cmd/template.html")
		if err != nil {
			http.Error(response, "Error parsing template", 500)
		}

		tmpl.Execute(response, story.NewYork)
	})
	mux.HandleFunc("/debate", func(response http.ResponseWriter, request *http.Request) {
		// Create a template and fill in the values with the json data
		tmpl, err := template.ParseFiles("cmd/template.html")
		if err != nil {
			http.Error(response, "Error parsing template", 500)
		}

		tmpl.Execute(response, story.Debate)
	})
	mux.HandleFunc("/sean-kelly", func(response http.ResponseWriter, request *http.Request) {
		// Create a template and fill in the values with the json data
		tmpl, err := template.ParseFiles("cmd/template.html")
		if err != nil {
			http.Error(response, "Error parsing template", 500)
		}

		tmpl.Execute(response, story.SeanKelly)
	})

	mux.HandleFunc("/mark-bates", func(response http.ResponseWriter, request *http.Request) {
		// Create a template and fill in the values with the json data
		tmpl, err := template.ParseFiles("cmd/template.html")
		if err != nil {
			http.Error(response, "Error parsing template", 500)
		}

		tmpl.Execute(response, story.MarkBates)
	})

	mux.HandleFunc("/denver", func(response http.ResponseWriter, request *http.Request) {
		// Create a template and fill in the values with the json data
		tmpl, err := template.ParseFiles("cmd/template.html")
		if err != nil {
			http.Error(response, "Error parsing template", 500)
		}

		tmpl.Execute(response, story.Denver)
	})

	mux.HandleFunc("/home", func(response http.ResponseWriter, request *http.Request) {
		// Create a template and fill in the values with the json data
		tmpl, err := template.ParseFiles("cmd/template.html")
		if err != nil {
			http.Error(response, "Error parsing template", 500)
		}

		tmpl.Execute(response, story.Home)
	})

	fmt.Println("Starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
