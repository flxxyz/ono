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
	c.JSON(http.StatusOK, NewBody(0, message))
}

func (c *Context) Fail(message string) {
	c.JSON(http.StatusOK, NewBody(1, message))
}

func (c *Context) Response(code int, message string, data interface{}) {
	obj := NewBody(code, message)
	obj["data"] = data
	c.JSON(http.StatusOK, obj)
}

func NewBody(code int, message string) H {
	return H{
		"code":    code,
		"message": message,
	}
}
