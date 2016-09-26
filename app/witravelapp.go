package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/alexxuyao/witravel/handler"
	"github.com/alexxuyao/witravel/module"
	"github.com/kataras/iris"
)

type WebApp struct {
}

func (app *WebApp) Start() {

	log.SetLevel(log.DebugLevel)

	mconfig := &module.ConfigModule{}
	mtoken := &module.AccessTokenModule{AppId: mconfig.GetConfig().Wechat.AppId, AppSecret: mconfig.GetConfig().Wechat.AppSecret}
	mcontainer := &module.WebContainer{Token: mtoken, Config: mconfig.GetConfig()}

	if mconfig.GetConfig().IsDebug {
		// this will reload the templates on each request, defaults to false
		iris.Config.IsDevelopment = true
	}

	// server the static file
	iris.Static("/static", "./static/", 1)

	// 公共的，给微信做回调用的，不用验证权限，但有的要做签名校验
	pub := iris.Party("/pub")

	{
		pub.Get("/wechat", handler.MiscHandler)
	}

	// 用户访问的，要用微信登录的
	wiTravel := iris.Party("/weapp")

	{
		// add a silly middleware
		wiTravel.UseFunc(func(c *iris.Context) {

			webuser := &module.WebUser{Context: c}

			c.Set("container", mcontainer)
			c.Set("webuser", webuser)

			//your authentication logic here...
			//			log.Debugln("from ", c.PathString())
			//			log.Debugln("host ", c.HostString())
			//			log.Debugln("requestURI ", string(c.RequestURI()))
			//			log.Debugln("QueryString ", string(c.URI().QueryString()))
			//			log.Debugln("String ", string(c.URI().String()))
			log.Infoln("FullURI ", string(c.URI().FullURI()))

			module.DoAuthFilter(c)

		})

		wiTravel.Get("/", handler.IndexHandler)
		wiTravel.Post("/travellist", handler.TravelListHandler)

	}

	// 管理后台，管理员用微信登录的
	admin := iris.Party("/manage")

	{
		// add a silly middleware
		admin.UseFunc(func(c *iris.Context) {
			//your authentication logic here...
			println("from ", c.PathString())

			c.Set("container", mcontainer)

			authorized := true
			if authorized {
				c.Next()
			} else {
				c.Text(401, c.PathString()+" is not authorized for you")
			}

		})

		admin.Get("/initmenu", handler.InitMenuHandler)
	}

	iris.Listen(":80")
}

func main() {

	app := &WebApp{}
	app.Start()
}
