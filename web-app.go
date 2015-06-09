package main

import (
    "log"
    "net/http"

    Ctrl "github.com/facundoj/web-abm/controller"
)

func main() {
    mux := http.NewServeMux()

    Ctrl.InitUsers()

    mux.HandleFunc("/", Ctrl.ListHandler)
    mux.HandleFunc("/create/", Ctrl.CreateHandler)
    mux.HandleFunc("/delete/", Ctrl.DeleteHandler)
    mux.HandleFunc("/edit/", Ctrl.EditHandler)

    err := http.ListenAndServe(":8080", mux)

    if err != nil {
        log.Print("Error")
    }
}
