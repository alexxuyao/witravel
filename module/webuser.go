package module

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/alexxuyao/witravel/dao"
	"github.com/alexxuyao/witravel/model"
	"github.com/astaxie/beego/orm"
	"github.com/kataras/iris"
)

type WebUser struct {
	Context *iris.Context
}

// 是否已经登录
func (u *WebUser) IsAuthorized() bool {

	ret := true

	if v := u.Context.Session().Get("webuser"); nil == v {
		ret = false
	}

	if !ret {

		// 在开发环境下用谷哥浏览器访问
		if iris.Config.IsDevelopment {
			if !strings.Contains(u.Context.RequestHeader("User-Agent"), "QQ") {
				u.DoLogin("oCV2Lv4EPu766DWXa5uXTmtDLD00")
				ret = true

				log.Infoln("do debug login. user agent is:", u.Context.RequestHeader("User-Agent"))
			}
		}
	}

	return ret
}

// 登录操作
func (u *WebUser) DoLogin(openId string) {

	user := &model.User{OpenId: openId}

	if err := dao.DoTransaction(func(o orm.Ormer) error {

		err := o.Read(user, "OpenId")

		if nil != err {
			return err
		}

		return nil
	}); nil == err {
		u.Context.Session().Set("webuser", user)
		log.Infoln("user login:", user.Nickname)
	} else {
		// dologin fail
		log.Infoln("login fail, openId:", openId)
	}

}

// 取得当前用户
func (u *WebUser) Get() *model.User {
	return u.Context.Session().Get("webuser").(*model.User)
}
