package tests

var userBody = `{
  "first_name": "Juan",
  "last_name" : "Perez",
  "email" : "jb00@bad.seed",
  "password": "juanito",
  "answers" : {"city" : "CDMX", "money" : "lots"}
}`

var loginBody = `{
  "email": "jb00@bad.seed",
  "password": "juanito"
}`
