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

	wiTravel := iris.Party("/wiTravel")

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

		wiTravel.Get("/wechat", handler.MiscHandler)

		wiTravel.Get("/weapp/travellist", handler.TravelListHandler)

		wiTravel.Get("/", func(c *iris.Context) {
			c.Write("from /wiTravel/ or /wiTravel if you pathcorrection on")
		})

		wiTravel.Get("/dashboard", func(c *iris.Context) {
			c.Write("/wiTravel/dashboard")
		})

		wiTravel.Delete("/delete/:userId", func(c *iris.Context) {
			c.Write("wiTravel/delete/%s", c.Param("userId"))
		})
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

		admin.Get("/initMenu", handler.InitMenuHandler)
	}

	iris.Listen(":8090")
}

func main() {

	app := &WebApp{}
	app.Start()
}
