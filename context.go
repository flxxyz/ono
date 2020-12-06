package ono

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type HandleFunc func(ctx *Context)

type H gin.H

type Context struct {
	TimeoutCtx context.Context
	*gin.Context
}

func (c *Context) RequestBody(data interface{}) interface{} {
	body, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(body, data)
	return data
}

func (c *Context) Success(message string) {
	c.Response(http.StatusOK, H{
		"code":    0,
		"message": message,
	})
}

func (c *Context) Fail(message string) {
	c.Response(http.StatusOK, H{
		"code":    1,
		"message": message,
	})
}

func (c *Context) Response(code int, data interface{}) {
	c.JSON(code, data)
}
