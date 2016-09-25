package module

import "github.com/kataras/iris"

type WebUser struct {
	Context *iris.Context
}

func (u *WebUser) IsAuthorized() bool {
	return false
}

func (u *WebUser) DoLogin(openId string) {

}
