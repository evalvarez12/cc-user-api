package controllers

import (
	"financy/api/app/ds"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type message struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Ver conf/routes
func (c App) CORS() revel.Result {
	c.headers()
	return c.RenderJson(nil)
}

func (c App) headers() {
	if origin := c.Request.Header.Get("Origin"); origin != "" {
		c.Response.Out.Header().Add("Access-Control-Allow-Origin", origin)
		c.Response.Out.Header().Add("Access-Control-Allow-Headers", c.Request.Header.Get("Access-Control-Request-Headers"))
		// c.Response.Out.Header().Add("Access-Control-Allow-Methods", c.Request.Header.Get("Access-Control-Request-Methods"))
		c.Response.Out.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	}
}

func (c App) OK() revel.Result {
	c.headers()
	return c.RenderJson(message{
		Success: true,
	})
}

func (c App) Error(err error) revel.Result {
	c.headers()
	return c.RenderJson(message{
		Success: false,
		Error:   err.Error(),
	})
}

// TODO check this error type for validation
func (c App) ErrorData(data interface{}) revel.Result {
	c.headers()
	return c.RenderJson(message{
		Success: false,
		Data:    data,
	})
}

func (c App) Data(data interface{}) revel.Result {
	c.headers()
	return c.RenderJson(message{
		Success: true,
		Data:    data,
	})
}

func (c App) GetSession() (uint, error) {
	sToken := c.Request.Header.Get("X-Auth-Token")
	token, err := jwt.Parse(sToken, "ccsignature")
	if err == nil && token.Valid {
		return token.Claims["user_id"], nil
	}
	return 0, jwt.ErrNoTokenInRequest
}
