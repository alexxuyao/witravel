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

	log.Debugln("this user agent is:", u.Context.RequestHeader("User-Agent"))
	if iris.Config.IsDevelopment {
		if !strings.Contains(u.Context.RequestHeader("User-Agent"), "QQ") {
			return true
		}
	}

	if v := u.Context.Session().Get("webuser"); nil == v {
		return false
	}

	return true
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
	} else {
		// dologin fail
		log.Infoln("login fail, openId:", openId)
	}

}

// 取得当前用户
func (u *WebUser) Get() *model.User {
	return u.Context.Session().Get("webuser").(*model.User)
}
