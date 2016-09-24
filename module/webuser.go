package module

import "github.com/kataras/iris"

type WebUser struct {
	Context *iris.Context
}

func (u *WebUser) IsAuthorized() bool {
	return false
}
