package tests

var userBody = `{
  "first_name": "Juan",
  "last_name" : "Perez",
  "email" : "juanp@test.testy",
  "password": "juanito",
  "public": true
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

var answers_update = `{
  "answers" : {"city" : "GDL", "money" : "none", "CO2" : "lots"}
}`

var location_set = `{
  "location" : {"city" : "GDL", "money" : "none", "CO2" : "lots"}
}`
