package handler

import (
	"math/rand"
	"time"

	"github.com/kataras/iris"
)

//
func IndexHandler(c *iris.Context) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	c.MustRender("index.html", struct {
		Path string
		Rand int
	}{Path: "", Rand: r.Int()})
}
