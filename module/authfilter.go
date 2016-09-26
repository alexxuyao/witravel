package module

import (
	"encoding/json"
	"errors"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/alexxuyao/witravel/common"
	"github.com/alexxuyao/witravel/dao"
	"github.com/alexxuyao/witravel/model"
	"github.com/astaxie/beego/orm"

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
	Sex        int      `json:"sex"`
	Language   string   `json:"language"`
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
			log.Debugln("do auth filter error:", err.Error())
			c.Render("error.html", struct{ ErrorMsg string }{ErrorMsg: err.Error()})
		}
	}()

	webuser := c.Get("webuser").(*WebUser)
	config := c.Get("container").(*WebContainer).Config
	state := c.URLParam("state")

	if !webuser.IsAuthorized() {

		log.Debugln("not authorized!")

		// 如果有wechat_redirect_type参数，表示是从微信授权跳转回来的
		if rType, ok := c.URLParams()["wechat_redirect_type"]; ok {

			log.Debugln("access from wechat, wechat_redirect_type is:", rType)

			code := c.URLParam("code")

			// 如果用户不授权，没有code
			if code = strings.TrimSpace(code); code != "" {

				log.Infoln("access from wechat, code is:", code)

				// 2.
				// 通过code换取网页授权access_token
				// https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code

				res, err := common.DoGet("https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + config.Wechat.AppId + "&secret=" + config.Wechat.AppSecret + "&code=" + code + "&grant_type=authorization_code")

				common.CheckError(err)

				restr := string(res)
				var sresp *AuthGetAccesTokenResp

				log.Infoln("get access_token:", restr)

				if strings.Contains(restr, "access_token") {
					sresp = &AuthGetAccesTokenResp{}
					err = json.Unmarshal(res, sresp)

					common.CheckError(err)
				} else {
					common.PanicWechatError(res)
				}

				// 3. 检查此用户是否已经注册
				if !isUserRegister(sresp.OpenId) {

					log.Infoln("the user is not register, openId is:", sresp.OpenId)

					if rType == "base" {
						//
						c.Render("redirect.html", RedirectVo{AppId: config.Wechat.AppId, State: state, Scope: "snsapi_userinfo", WechatRedirectType: "userinfo"})
						return
					} else if rType == "userinfo" {
						// 4. 拉取用户信息
						wechatuser, err := getWechatUserinfo(sresp.AccessToken, sresp.OpenId)
						common.CheckError(err)

						// 5. 注册用户
						log.Infoln("start do register, openId is:", sresp.OpenId)
						userRegister(wechatuser)
					}

				} else {
					log.Infoln("user has register, will auto login now, openId is:", sresp.OpenId)
				}

				// 6. 用户登录
				webuser.DoLogin(sresp.OpenId)
			} else {
				// 用户没有同意授权
				log.Infoln("access deny from wechat, rType is:", rType)
			}
		}
	}

	if webuser.IsAuthorized() {
		c.Next()
	} else {
		log.Infoln("go to redirect.html")

		// c.MustRender("redirect.html", RedirectVo{AppId: config.Wechat.AppId, State: state, Scope: "snsapi_base", WechatRedirectType: "base"})
		c.MustRender("redirect.html", RedirectVo{AppId: config.Wechat.AppId, State: state, Scope: "snsapi_userinfo", WechatRedirectType: "userinfo"})
		//rurl := url.QueryEscape(string(c.URI().FullURI()))
		//log.Debugln(rurl)
		//c.Redirect("https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + config.Wechat.AppId + "&redirect_uri=" + rurl + "&response_type=code&scope=snsapi_userinfo&state=0#wechat_redirect")
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

	log.Debugln("get userinfo from wechat, openId is ", openId, ", and result is :", restr)

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

// 注册用户
func userRegister(wechatuser *WechatUserInfo) error {

	return dao.DoTransaction(func(o orm.Ormer) error {

		u := model.User{OpenId: wechatuser.OpenId, City: wechatuser.City, Country: wechatuser.Country, HeadImgUrl: wechatuser.HeadImgUrl, Nickname: wechatuser.Nickname, Privilege: strings.Join(wechatuser.Privilege, ","), Province: wechatuser.Province, Sex: wechatuser.Sex, UnionId: wechatuser.UnionId}
		_, err := o.Insert(&u)

		if nil != err {
			return err
		}

		return nil
	})
}

// 用户是否已经登录
func isUserRegister(openId string) bool {

	registered := true

	if err := dao.DoTransaction(func(o orm.Ormer) error {

		u := model.User{OpenId: openId}
		err := o.Read(&u, "OpenId")

		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		registered = false
	}

	return registered
}
