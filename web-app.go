package main

import (
    "net/http"
    "log"
    "text/template"
    "github.com/facundoj/web-abm/entity"
    "strconv"
)

var users map[int64]entity.User

func main() {
    mux := http.NewServeMux()

    users = make(map[int64]entity.User)

    mux.HandleFunc("/create/", createHandler)
    mux.HandleFunc("/list/", listHandler)

    err := http.ListenAndServe(":8080", mux)

    if err != nil {
        log.Print("Error")
    }
}

func createHandler(res http.ResponseWriter, req *http.Request) {
    log.Print(req.Method)
    log.Print("Create")

    if req.Method == "GET" {
        t, err := template.ParseFiles("views/create.tpl")
        if err != nil {
            log.Print("Error loading template")
        }

        t.Execute(res, nil)

    } else if req.Method == "POST" {

        name := req.FormValue("name")
        surname := req.FormValue("surname")
        dni, _ := strconv.Atoi(req.FormValue("dni"))

        _, ok := users[int64(dni)]

        if !ok {
            users[int64(dni)] = entity.User{Name: name, Surname: surname, Dni: int64(dni)}
            log.Println("User created", users[int64(dni)])
        } else {
            log.Println("User already exists", dni)
        }

        http.Redirect(res, req, "/list/", http.StatusFound)
    }
}

func listHandler(res http.ResponseWriter, req *http.Request) {
    log.Print(req.Method)
    log.Print("List")

    t, err := template.ParseFiles("views/list.tpl")
    if err != nil {
        log.Print("Error loading template")
    }

    t.Execute(res, users)
}