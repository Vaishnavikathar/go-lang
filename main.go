package main

import "net/http"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello Vaishnavi !!!"))
    })

    http.ListenAndServe(":8080", nil) // HTTP, NOT HTTPS
}
