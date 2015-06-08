<!doctype html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>Web ABM - List</title>
    </head>
    <body>
        <h1>Users list</h1>
        <ul>
            {{range .}}
                <li>{{.Surname}}, {{.Name}} ({{.Dni}}) - <a href="/delete/{{.Dni}}">Remove</a></li>
            {{end}}
        </ul>
        <a href="/create/">Add user</a>
    </body>
</html>