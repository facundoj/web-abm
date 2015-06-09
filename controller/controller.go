package controller

import (
    "log"
    "net/http"
    "strconv"
    "text/template"
)

// User es una estructura que agrupa la informacion de los usuarios
type User struct {
    Dni int64
    Name string
    Surname string
}

// coleccion de usuarios - closure
var users map[int64]User

// InitUsers crea la colleccion donde se alamacenaran los usuarios
func InitUsers()  {
	// usaremos el DNI como key
    users = make(map[int64]User)
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

        // recuperamos el usuario con ese DNI del mapa
        _, ok := users[int64(dni)]

        if !ok {
            // lo agregamos solo si no fue encontrado
            users[int64(dni)] = User{Name: name, Surname: surname, Dni: int64(dni)}
            log.Println("User created", users[int64(dni)])
        } else {
            // el usuario con ese DNI ya existe, por lo que no se agrega
            log.Println("User already exists", dni)
        }

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
    t.Execute(res, users)
}

// DeleteHandler sirve para eliminar un usuario
func DeleteHandler(res http.ResponseWriter, req *http.Request)  {
    // eliminacion del usuario
    userDniStr := req.URL.Path[len("/delete/"):]
    userDni, _ := strconv.Atoi(userDniStr)

    // eliminamos el usuario del mapa
    delete(users, int64(userDni))

    // redirigimos a root para ver la lista
    http.Redirect(res, req, "/", http.StatusFound)
}

// EditHandler sirve para editar un usuario existente. GET para obtener el formulario y POST para modificarlo
func EditHandler(res http.ResponseWriter, req *http.Request)  {
    if req.Method == "GET" {
        // renderizamos el form para edicion
        userDniStr := req.URL.Path[len("/edit/"):]
        userDni, _ := strconv.Atoi(userDniStr)

        // recuperamos el usuario del mapa de usuarios
        user, ok := users[int64(userDni)]

        if !ok {
            // si el usuario no fue encontrado, volvemos a la lista
            log.Print("User not exists")
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

        user, ok := users[int64(dni)]

        // si encontre el usuario, lo modifico
        if ok {
            user.Name = newName
            user.Surname = newSurname

            // no se modifica el elemento del mapa, sino una copia
            // luego, se pisa el valor
            // Info sobre el issue: https://code.google.com/p/go/issues/detail?id=3117
            users[int64(dni)] = user
        }

        // redirijo a la lista de usuarios
        http.Redirect(res, req, "/", http.StatusFound)
    }
}
