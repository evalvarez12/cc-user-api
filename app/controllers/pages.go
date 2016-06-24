package controllers

import (
	"github.com/revel/revel"
	// "net/http"
)

type Page struct {
	App
}

func (c Page) PasswordReset(id uint, token string) revel.Result {
	// req, err = http.NewRequest("POST", c.BaseUrl()mt+"/user/reset", c.RenderJson(reset))
	// req.Header.Set("Content-Type", contentType)
	// fmt.Println("EL pass: ", reset.Password)
	return c.Render(id, token)
}
