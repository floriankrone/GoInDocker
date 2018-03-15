//Go Server 

package main

import (
        "io"
        "net/http"
)

func text(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "Es funktioniert!")
}

func main() {
        http.HandleFunc("/", text)
        http.ListenAndServe(":8000", nil)
}
