package tests

var userBody = `{
  "first_name": "Juan",
  "last_name" : "Perez",
  "email" : "juanp@test.testy",
  "password": "juanito",
  "answers" : {"city" : "CDMX", "money" : "lots"}
}`

var userBody_update = `{
  "first_name": "Juanito",
  "last_name" : "Perezin"
 }`

var loginBody = `{
  "email": "juanp@test.testy",
  "password": "juanito"
}`

var loginBody_badEmail = `{
  "email": "juanpp@test.testy",
  "password": "juanito"
}`

var loginBody_badPassword = `{
  "email": "juanp@test.testy",
  "password": "juanito2"
}`
