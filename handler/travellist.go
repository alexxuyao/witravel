package handler

import (
	"time"

	"github.com/kataras/iris"
)

// 用于微信提交认证
func TravelListHandler(c *iris.Context) {

	time.Sleep(2 * time.Second)

	c.JSON(0, struct{}{})
}
