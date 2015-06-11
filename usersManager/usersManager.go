package usersManager

import (
    "log"
    "errors"
)

// User es una estructura que agrupa la informacion de los usuarios
type User struct {
    Dni int64
    Name string
    Surname string
}

// coleccion de usuarios - closure
var users map[int64]User

// Init crea una colleccion donde almacenara los usuarios
func Init()  {
    // usaremos el DNI como key
    users = make(map[int64]User)
}

// CreateUser crea un nuevo usuario y lo agrega a la colleccion
func CreateUser(name string, surname string, dni int64)  {
    // recuperamos el usuario con ese DNI del mapa
    _, ok := users[dni]

    if !ok {
        // lo agregamos solo si no fue encontrado
        users[dni] = User{Name: name, Surname: surname, Dni: dni}
        log.Println("User created", users[dni])
    } else {
        // el usuario con ese DNI ya existe, por lo que no se agrega
        log.Println("User already exists", dni)
    }
}

// DeleteUser elimina el usuario especificado de la coleccion de usuarios
func DeleteUser(dni int64)  {
    // eliminamos el usuario del mapa
    delete(users, dni)
}

// EditUser modifica el usuario con el DNI especificado
func EditUser(dni int64, newName string, newSurname string)  {
    user, ok := users[int64(dni)]

    // si encontre el usuario, lo modifico
    if ok {
        user.Name = newName
        user.Surname = newSurname

        // no se modifica el elemento del mapa, sino una copia
        // luego, se pisa el valor
        // Info sobre el issue: https://code.google.com/p/go/issues/detail?id=3117
        // http://play.golang.org/p/qRl3wG5yVU
        users[int64(dni)] = user
    }
}

// GetUsers devuelve el mapa de usuarios existentes en el momento
func GetUsers() []User {
    usersList := make([]User, len(users))
    i := 0
    for _, user := range users {
        usersList[i] = user
        i = i + 1
    }
    return usersList
}

// GetUser devuelve el usuario con el DNI especificado
func GetUser(dni int64) (User, error)  {
    user, ok := users[dni]

    if !ok {
        // si el usuario no fue encontrado, volvemos a la lista
        log.Print("User not exists")
        return user, errors.New("User doesn't exist")
    }

    return user, nil
}
