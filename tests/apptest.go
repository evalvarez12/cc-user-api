package tests

import (
	"encoding/json"
	"log"
	"github.com/revel/revel/testing"
	"io"
	"net/http"
	"strings"
)

var chargeBody = `{
	"name": "nueva carga",
	"category_id" : 2,
	"description": "taquitos y tequila",
	"expected_date": "2015-01-05T17:14:12-06:00",
	"real_amount": {
	  "value": 123,
	  "scale": 1
	},
	"coin": "MXN",
	"kind": "expense",
	"source": "Source B",
	"destination": "Destination D"

}`

var categoryBody = `{
      "name": "New Category",
	  "kind": "income"
}`

var accountBody = `{
  "name": "New Account",
  "bank": "Cartera",
  "description": "Gastar"
}`

var userBody = `{
  "name": "JhonyBoy66",
  "complete_name": "Juan Perez Perez",
  "email": "jb00@bad.seed",
  "created_at": "2016-03-23T13:14:29.141081-06:00",
  "updated_at": "2016-03-23T13:14:29.141081-06:00",
  "password": "juanito"
}`

var loginBody = `{
  "name": "JhonyBoy66",
  "password": "juanito"
}`

var token string

type AppTest struct {
	testing.TestSuite
}

type apiResult struct {
	Success bool        `json:"succes"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type loginResult struct {
	Success bool                   `json:"succes"`
	Error   string                 `json:"error,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

func myPost(path string, contentType string, reader io.Reader, token string, t *AppTest) *http.Request {
	req, err := http.NewRequest("POST", t.BaseUrl()+path, reader)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("X-Auth-Token", token)
	return req
}

func myGet(path, token string, t *AppTest) *http.Request {
	req, err := http.NewRequest("GET", t.BaseUrl()+path, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Auth-Token", token)
	return req
}

func testSucces(t *AppTest) {
	var result apiResult
	err := json.Unmarshal(t.ResponseBody, &result)
	t.AssertEqual(err, nil)
	// t.AssertEqual(result.Success, true)
}

// --------------- TEST FUNCTIONS -------------

func (t *AppTest) TestChargesList() {
	req := myGet("/charges", token, t)
	t.NewTestRequest(req).Send()
	t.AssertOk()
	testSucces(t)
	t.AssertContentType("application/json; charset=utf-8")
	log.Println(string(t.ResponseBody))
}

func (t *AppTest) TestChargesAdd() {
	req := myPost("/charges", "application/json; charset=utf-8", strings.NewReader(chargeBody), token, t)
	t.NewTestRequest(req).Send()
	t.AssertOk()
	testSucces(t)
	t.AssertContentType("application/json; charset=utf-8")
	log.Println(string(t.ResponseBody))
}

func (t *AppTest) TestAccountsList() {
	req := myGet("/accounts", token, t)
	t.NewTestRequest(req).Send()
	t.AssertOk()
	testSucces(t)
	t.AssertContentType("application/json; charset=utf-8")
	log.Println(string(t.ResponseBody))
}

func (t *AppTest) TestAccountAdd() {
	req := myPost("/accounts", "application/json; charset=utf-8", strings.NewReader(accountBody), token, t)
	t.NewTestRequest(req).Send()
	t.AssertOk()
	testSucces(t)
	t.AssertContentType("application/json; charset=utf-8")
	log.Println(string(t.ResponseBody))
}

func (t *AppTest) TestCategoriesList() {
	req := myGet("/categories", token, t)
	t.NewTestRequest(req).Send()
	t.AssertOk()
	testSucces(t)
	t.AssertContentType("application/json; charset=utf-8")
	log.Println(string(t.ResponseBody))
}

func (t *AppTest) TestCategoriesPie() {
	req := myGet("/categories/pie", token, t)
	t.NewTestRequest(req).Send()
	t.AssertOk()
	testSucces(t)
	t.AssertContentType("application/json; charset=utf-8")
	log.Println(string(t.ResponseBody))
}

func (t *AppTest) TestCategoryAdd() {
	req := myPost("/categories", "application/json; charset=utf-8", strings.NewReader(categoryBody), token, t)
	t.NewTestRequest(req).Send()
	t.AssertOk()
	testSucces(t)
	t.AssertContentType("application/json; charset=utf-8")
	log.Println(string(t.ResponseBody))
}

func (t *AppTest) TestAUserAdd() {
	t.Post("/user", "application/json; charset=utf-8", strings.NewReader(userBody))
	t.AssertOk()
	testSucces(t)
	t.AssertContentType("application/json; charset=utf-8")
	log.Println(string(t.ResponseBody))
}

func (t *AppTest) TestAUserLogin() {
	t.Post("/user/login", "application/json; charset=utf-8", strings.NewReader(loginBody))
	buf := t.ResponseBody
	var logRes loginResult
	err := json.Unmarshal(buf, &logRes)
	t.AssertEqual(err, nil)
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
	token = logRes.Data["token"].(string)
	log.Println(string(t.ResponseBody))
	log.Println("Setting TOKEN to: " + token)
}

func (t *AppTest) Before() {
	log.Println("-------->")
}

func (t *AppTest) After() {
	log.Println("<--------")
}
