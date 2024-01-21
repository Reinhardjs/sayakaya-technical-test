### Available endpoints :
- `GET` localhost:9090/users
- `GET` localhost:9090/users/birthday?month=:month&day=:day
- `GET` localhost:9090/users/:id
- `PUT` localhost:9090/users/:id
- `DELETE` localhost:9090/users/:id
- `POST` localhost:9090/users

<br>

### Example Response Body Payload for `GET` /users :
```
[
    {
        "id": 1,
        "email": "reinhardjsilalahi@gmail.com",
        "verifiedStatus": true,
        "birthday": "1998-01-22"
    },
    {
        "id": 2,
        "email": "reinhardjonathansilalahi@gmail.com",
        "verifiedStatus": true,
        "birthday": "1999-01-22"
    }
]
```

### Example Response Body Payload for `GET` /users/birthday?month=01&day=22 :
```
[
    {
        "id": 1,
        "email": "reinhardjsilalahi@gmail.com",
        "verifiedStatus": true,
        "birthday": "1998-01-22"
    },
    {
        "id": 2,
        "email": "reinhardjonathansilalahi@gmail.com",
        "verifiedStatus": true,
        "birthday": "1999-01-22"
    }
]
```

<br>

### Example Request Body Payload for `POST` /users :
```
{
    "email": "user1@email.com",
    "verifiedStatus": true,
    "birthday": "1999-01-22"
}
```

### Example Response Body Payload for `POST` /users :
```
{
    "id": 1,
    "email": "user1@email.com",
    "verifiedStatus": true,
    "birthday": "1999-01-22T00:00:00Z"
}
```

<br>

### Example Request Body Payload for `PUT` /users/1 :
```
{
    "email": "user1@email.com",
    "verifiedStatus": true,
    "birthday": "1999-01-22"
}
```

### Example Response Body Payload for `PUT` /users/1 :
```
{
    "id": 1,
    "email": "user1@email.com",
    "verifiedStatus": true,
    "birthday": "1999-01-22T00:00:00Z"
}
```

<br>

### Table Schema
---
![image](https://github.com/Reinhardjs/sayakaya-technical-test/assets/7758970/057b52c1-ee07-4f9f-afcc-1f2ea0368809)
![image](https://github.com/Reinhardjs/sayakaya-technical-test/assets/7758970/0ad0c8a5-11af-4249-8a27-c2869d04d547)
