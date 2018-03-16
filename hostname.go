package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	{
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Es funktioniert und ich laufe auf: ")
			fmt.Fprintf(w, name)
		})
		http.ListenAndServe(":8000", nil)
	}
}
