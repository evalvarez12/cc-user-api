package tests

import (
	"encoding/json"
	"log"
	"github.com/revel/revel/testing"
	"io"
	"net/http"
	"strings"
)

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


func myVERB(verb, path string, contentType string, reader io.Reader, token string, t *AppTest) *http.Request {
	var err error
	var req *http.Request
	switch verb {
	case "POST" :
		req, err = http.NewRequest("POST", t.BaseUrl()+path, reader)
		req.Header.Set("Content-Type", contentType)
	case "GET" :
		req, err = http.NewRequest("GET", t.BaseUrl()+path, nil)
	case "DELETE" :
		req, err = http.NewRequest("DELETE", t.BaseUrl()+path, nil)
	}
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", token)
	return req
}


func testSucces(t *AppTest) {
	var result apiResult
	err := json.Unmarshal(t.ResponseBody, &result)
	t.AssertEqual(err, nil)
	// t.AssertEqual(result.Success, true)
}

// --------------- TEST FUNCTIONS -------------

func (t *AppTest) TestAdd() {
	t.Post("/user", "application/json; charset=utf-8", strings.NewReader(userBody))
	t.AssertOk()
	testSucces(t)
	t.AssertContentType("application/json; charset=utf-8")
	log.Println(string(t.ResponseBody))
}

func (t *AppTest) TestLogin() {
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

func (t *AppTest) TestDelete() {
	req := myVERB("DELETE", "/user", "", nil, token, t)
	t.NewTestRequest(req).Send()
	t.AssertOk()
	testSucces(t)
	t.AssertContentType("application/json; charset=utf-8")
	log.Println(string(t.ResponseBody))
}

func (t *AppTest) Before() {
	log.Println("-------->")
}

func (t *AppTest) After() {
	log.Println("<--------")
}
