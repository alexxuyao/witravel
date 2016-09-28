package handler

import (
	"github.com/alexxuyao/captcha"
	"github.com/alexxuyao/witravel/common"
	"github.com/kataras/iris"
)

type ValidateCaptchaResp struct {
	Result    bool   `json:"result"`
	CaptchaId string `json:"captchaId"`
}

// 生成一个新的验证码ID
func CaptchaIdHandler(c *iris.Context) {
	catpchaId := captcha.NewLen(4)
	common.AjaxRespSuccess(c, catpchaId)
}

// 校验验证码
func ValidateCaptchaHandler(c *iris.Context) {
	value := c.Param("value")
	captchaId := c.Param("captchaId")

	if captcha.VerifyStringSimple(captchaId, value) {
		common.AjaxRespSuccess(c, ValidateCaptchaResp{Result: true})
	} else {
		// clear the old one
		captcha.VerifyString(captchaId, value)

		// make a new one
		captchaId = captcha.NewLen(4)
		common.AjaxRespSuccess(c, ValidateCaptchaResp{Result: false, CaptchaId: captchaId})
	}

}
