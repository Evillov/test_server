package main

import (
	"fmt"
	"net/http"
	"os"
)

type Words struct {
	Id   int    `json:"id"`
	Word string `json:"word"`
}

func (word *Words) HandlerWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", word.Word)

}

func main() {

	hostname, _ := os.Hostname()
	word := Words{
		Id:   1,
		Word: hostname,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/word", word.HandlerWord)

	http.ListenAndServe(":666", mux)

}
