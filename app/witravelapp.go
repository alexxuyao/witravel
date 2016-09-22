package main

import (
	"github.com/alexxuyao/witravel/handler"
	"github.com/alexxuyao/witravel/module"
	"github.com/kataras/iris"
)

type WebApp struct {
}

func (app *WebApp) Start() {

	mconfig := &module.ConfigModule{}
	mtoken := &module.AccessTokenModule{AppId: mconfig.GetConfig().Wechat.AppId, AppSecret: mconfig.GetConfig().Wechat.AppSecret}
	mcontainer := &module.WebContainer{Token: mtoken, Config: mconfig.GetConfig()}

	if mconfig.GetConfig().IsDebug {
		// this will reload the templates on each request, defaults to false
		iris.Config.IsDevelopment = true
	}

	// server the static file
	iris.Static("/static", "./static/", 1)

	wiTravel := iris.Party("/weapp")

	{
		// add a silly middleware
		wiTravel.UseFunc(func(c *iris.Context) {

			c.Set("container", mcontainer)

			//your authentication logic here...
			println("from ", c.PathString())
			authorized := true
			if authorized {
				c.Next()
			} else {
				c.Text(401, c.PathString()+" is not authorized for you")
			}

		})

		wiTravel.Get("/", handler.IndexHandler)
		wiTravel.Get("/wechat", handler.MiscHandler)
		wiTravel.Get("/travellist", handler.TravelListHandler)

		//		wiTravel.Get("/", func(c *iris.Context) {
		//			c.Write("from /wiTravel/ or /wiTravel if you pathcorrection on")
		//		})

		//		wiTravel.Get("/dashboard", func(c *iris.Context) {
		//			c.Write("/wiTravel/dashboard")
		//		})

		//		wiTravel.Delete("/delete/:userId", func(c *iris.Context) {
		//			c.Write("wiTravel/delete/%s", c.Param("userId"))
		//		})
	}

	// 管理后台
	admin := wiTravel.Party("/manage")

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

	iris.Listen(":8090")
}

func main() {

	app := &WebApp{}
	app.Start()
}
