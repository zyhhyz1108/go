package main

import (
	"net/http"
	"text/template"
)

type Game struct {
	Board [][]string
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	game := Game{
		Board: make([][]string, 15),
	}
	for i := range game.Board {
		game.Board[i] = make([]string, 15)
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, game)
}