package handler

import (
	"github.com/alexxuyao/witravel/common"
	"github.com/dchest/captcha"
	"github.com/kataras/iris"
)

//
func CaptchaIdHandler(c *iris.Context) {
	catpchaId := captcha.NewLen(4)
	common.AjaxRespSuccess(c, catpchaId)
}
