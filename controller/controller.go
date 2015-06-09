package controller

import (
    "log"
    "net/http"
    "strconv"
    "text/template"

    UsersManager "github.com/facundoj/web-abm/usersManager"
)

// InitController prepara el controlador para funcionar
func InitController()  {
    UsersManager.Init()
}

// CreateHandler sirve para crear usuarios. POST para agregar el usuario y GET para obtener el formulario
func CreateHandler(res http.ResponseWriter, req *http.Request) {
    if req.Method == "GET" {
        // renderizamos el formulario de creacion de usuario
        t, err := template.ParseFiles("views/create.tpl")
        if err != nil {
            log.Print("Error loading template")
        }

        t.Execute(res, nil)

    } else if req.Method == "POST" {
        // creacion del usuario con data ingresada
        name := req.FormValue("name")
        surname := req.FormValue("surname")
        dni, _ := strconv.Atoi(req.FormValue("dni"))

        UsersManager.CreateUser(name, surname, int64(dni))

        // redirigimos a root para ver la lista
        http.Redirect(res, req, "/", http.StatusFound)
    }
}

// ListHandler sirve para mostrar la lista de usuarios creados
func ListHandler(res http.ResponseWriter, req *http.Request) {
    // renderizamos la lista completa de usuarios
    t, err := template.ParseFiles("views/list.tpl")
    if err != nil {
        log.Print("Error loading template")
        return
    }
    t.Execute(res, UsersManager.GetUsers())
}

// DeleteHandler sirve para eliminar un usuario
func DeleteHandler(res http.ResponseWriter, req *http.Request)  {
    // eliminacion del usuario
    userDniStr := req.URL.Path[len("/delete/"):]
    userDni, _ := strconv.Atoi(userDniStr)

    UsersManager.DeleteUser(int64(userDni))

    // redirigimos a root para ver la lista
    http.Redirect(res, req, "/", http.StatusFound)
}

// EditHandler sirve para editar un usuario existente. GET para obtener el formulario y POST para modificarlo
func EditHandler(res http.ResponseWriter, req *http.Request)  {
    if req.Method == "GET" {
        // renderizamos el form para edicion
        userDniStr := req.URL.Path[len("/edit/"):]
        userDni, _ := strconv.Atoi(userDniStr)

        // obtenemos el usuario
        user, err := UsersManager.GetUser(int64(userDni))

        if err != nil {
            // si no se encontro, redirijo a la lista de usuarios
            http.Redirect(res, req, "/", http.StatusFound)
        }

        t, err := template.ParseFiles("views/edit.tpl")
        if err != nil {
            log.Print("Error loading template")
            return
        }
        t.Execute(res, user)

    } else if req.Method == "POST" {
        // modificacion del usuario con nueva data
        dniStr := req.FormValue("dni")
        newName := req.FormValue("name")
        newSurname := req.FormValue("surname")

        dni, _ := strconv.Atoi(dniStr)

        UsersManager.EditUser(int64(dni), newName, newSurname)

        // redirijo a la lista de usuarios
        http.Redirect(res, req, "/", http.StatusFound)
    }
}
