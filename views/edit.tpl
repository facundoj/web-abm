<!doctype html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Web ABM - Create</title>
    </head>
    <body>
        <h1>Edit User</h1>
        <form method="POST" action="/edit/">
            <input type="text" disabled value="{{.Dni}}">
            <input type="hidden" name="dni" value="{{.Dni}}">

            <input type="text" name="name" value="{{.Name}}">
            <input type="text" name="surname" value="{{.Surname}}">

            <input type="submit" value="Aceptar">
        </form>
    </body>
</html>
