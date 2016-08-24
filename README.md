# CC-USERS-API
## Setup

Copy and configure your environment variables:

```sh
cp .env.example .env
export $(cat .env | xargs)
```

Run in this order:
```
make images

make create-database

make create-api
```

*Address is 127.0.0.1:8082*

Enviromental variables:
```
CC_DBNAME - Name of the postgres DB
CC_DBUSER - Name of the postgres user
CC_DBPASS - Password of the user
CC_DBADDRESS - Address of the postgres service

CC_JWTSIGN - A secret string to sing JWT

CC_SPARKPOSTKEY - The key of the sparkpost app
```

## End-points
### Success:
```
{
  success: true,
  data: ...
}
```

### Error:
```
{
  success: false,
  message: “Error message”
}
```

## Routes
```
POST    /user              // Add a new user
POST    /user/login        // User login
GET     /user/logout       // Logout from current session
GET     /user/logoutall    // Logout user from all sessions
DELETE  /user              // Delete user
PUT     /user              // Update user (name or email)
PUT     /user/answers      // Update user answers
```

### Adding a user
Use:
```
POST     /user
```
Body:
```
{
  "first_name": "John",
  "last_name" : "Doe",
  "email" : "jdoe@test.testy",
  "password": "johnyboy",
  "answers" : {"city" : "CDMX", "money" : "lots"}
}
```
Validation on the fields:
```
"first_name" : required and min size 4
"last_name" : required and min size 4
"password" : required and min size 4
"email" : required and must math format regex
```

### Logging in
Use:
```
POST    /user/login
```
Body:
```
{
  "email": "jdoe@test.testy",
  "password": "johnyboy"
}
```
If successful returns:
```
{
  "success": true,
  "data": {
    "answers": "eyJDTzIiOiAibG90cyIsICJjaXR5IjogIkdETCIsICJtb25leSI6ICJub25lIn0=",
    "email": "jdoe@test.testy",
    "name": "John",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0Njc2Nzc4OTEsImlhdCI6MTQ2NjQ2ODI5MSwiaWQiOjY4LCJqdGkiOiJWZE44MyJ9.u-QfbyuieTRyiuqYIbxb01F0I1qdNUamQY4yMItrMhU",
    "user_id": 68
  }
}
```

### Update user answers
Use:
```
PUT     /user/answers
```
Body:
```
{
  "answers" : {"city" : "GDL", "money" : "none", "CO2" : "lots"}
}
```
