package module

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/alexxuyao/witravel/common"

	"github.com/kataras/iris"
)

type AuthGetAccesTokenResp struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
}

type WechatUserInfo struct {
	OpenId     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        string   `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	UnionId    string   `json:"unionid"`
}

type RedirectVo struct {
	AppId              string
	Scope              string
	State              string
	WechatRedirectType string // base, userinfo
}

func DoAuthFilter(c *iris.Context) {

	defer func() {
		if rerr := recover(); rerr != nil {
			err := rerr.(error)
			log.Println("do auth filter error:", err.Error())
			c.Render("error.html", struct{ ErrorMsg string }{ErrorMsg: err.Error()})
		}
	}()

	webuser := c.Get("webuser").(*WebUser)
	config := c.Get("container").(*WebContainer).Config
	state := c.URLParam("state")

	if !webuser.IsAuthorized() {

		// 如果有wechat_redirect_type参数，表示是从微信授权跳转回来的
		if rType, ok := c.URLParams()["wechat_redirect_type"]; ok {

			code := c.URLParam("code")

			// 如果用户不授权，没有code
			if strings.TrimSpace(code) != "" {
				// 2.
				// 通过code换取网页授权access_token
				// https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code

				res, err := common.DoGet("https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + config.Wechat.AppId + "&secret=" + config.Wechat.AppSecret + "&code=" + code + "&grant_type=authorization_code")

				common.CheckError(err)

				restr := string(res)
				var sresp *AuthGetAccesTokenResp

				if strings.Contains(restr, "access_token") {
					sresp = &AuthGetAccesTokenResp{}
					err = json.Unmarshal(res, sresp)

					common.CheckError(err)
				} else {
					common.PanicWechatError(res)
				}

				// 3. 检查此用户是否已经注册
				if !isUserRegister(sresp.OpenId) {

					if rType == "base" {
						//
						c.Render("redirect.html", RedirectVo{AppId: config.Wechat.AppId, State: state, Scope: "snsapi_userinfo", WechatRedirectType: "userinfo"})
						return
					} else if rType == "userinfo" {
						// 4. 拉取用户信息
						wechatuser, err := getWechatUserinfo(sresp.AccessToken, sresp.OpenId)
						common.CheckError(err)

						// 5. 注册用户
						userRegister(wechatuser)
					}

				}

				// 6. 用户登录
				webuser.DoLogin(sresp.OpenId)
			} else {
				// 用户没有同意授权

			}
		}
	}

	if webuser.IsAuthorized() {
		c.Next()
	} else {
		c.Render("redirect.html", RedirectVo{AppId: config.Wechat.AppId, State: state, Scope: "snsapi_base", WechatRedirectType: "base"})
	}
}

func getWechatUserinfo(accessToken, openId string) (*WechatUserInfo, error) {
	// 拉取用户信息
	// https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN

	res, err := common.DoGet("https://api.weixin.qq.com/sns/userinfo?access_token=" + accessToken + "&openid=" + openId + "&lang=zh_CN")

	if err != nil {
		return nil, err
	}

	restr := string(res)

	if !strings.Contains(restr, "openid") {

		err = common.GetWechatError(res)
		if nil != err {
			return nil, err
		}

		return nil, errors.New("unknow error:" + restr)
	}

	wechatuser := &WechatUserInfo{}
	err = json.Unmarshal(res, wechatuser)

	if err != nil {
		return nil, err
	}

	return wechatuser, nil
}

func userRegister(wechatuser *WechatUserInfo) error {
	return nil
}

func isUserRegister(openId string) bool {
	return false
}
