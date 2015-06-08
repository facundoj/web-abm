<!doctype html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Web ABM - Create</title>
</head>
<body>
    <h1>Create new User</h1>
    <form method="POST" action="/create/">
        <input type="text" placeholder="Name" name="name">
        <input type="text" placeholder="Surname" name="surname">
        <input type="number" placeholder="DNI" name="dni">
        <input type="submit" value="Create">
    </form>
</body>
</html>