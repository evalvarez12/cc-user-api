package tests

var userBody = `{
  "first_name": "Juan",
  "last_name" : "Perez",
  "email" : "juanp@test.testy",
  "password": "juanito",
  "city": "Austin",
  "county": "Hays",
  "state": "Texas",
  "country": "us",
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

var answers_update = `{"answers":{"result_food_total": "5", "result_housing_total": "6", "result_services_total": "3", "result_goods_total": "4", "result_transport_total": "8", "result_grand_total": "26"}}`

var location_set = `{"city":"Brooklyn","county":"Kings","state":"New York","country":"us"}`
