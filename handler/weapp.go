package handler

import "github.com/kataras/iris"

//
func IndexHandler(c *iris.Context) {
	c.MustRender("index.html", struct{ Path string }{Path: ""})
}
