package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Go Web Hello World!: %s\n", r.URL.Path)
    })

    http.ListenAndServe(":8081", nil)
}
