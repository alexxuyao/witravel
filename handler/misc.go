package handler

import (
	"fmt"

	"github.com/kataras/iris"
)

// 用于微信提交认证
func MiscHandler(c *iris.Context) {
	fmt.Println("signature is : " + c.URLParam("signature"))
	fmt.Println("timestamp is : " + c.URLParam("timestamp"))
	c.Write(c.URLParam("echostr"))
}
