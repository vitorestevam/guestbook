package main

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"text/template"
	"time"
)

type Message struct {
	Content string
	Moment  string
}

var Messages []Message

func saveMessages() {
	m, _ := json.Marshal(Messages)
	ioutil.WriteFile("data.txt", m, 0644)
}

func loadMessages() {
	data, _ := ioutil.ReadFile("data.txt")
	json.Unmarshal(data, &Messages)
}

func main() {

	loadMessages()

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			responseMessage := r.FormValue("message")
			if responseMessage != "" {
				moment := time.Now()
				message := Message{responseMessage, moment.Format("2006-01-02 15:04:05")}
				Messages = append([]Message{message}, Messages...)
				saveMessages()
			}
		}

		pageData := struct{ Messages []Message }{Messages}
		tmpl.Execute(w, pageData)
	})

	http.ListenAndServe(":8080", nil)
}
