package main

import (
	"github.com/alexxuyao/witravel/handler"
	"github.com/kataras/iris"
)

func main() {

	wiTravel := iris.Party("/wiTravel")

	{
		// add a silly middleware
		wiTravel.UseFunc(func(c *iris.Context) {
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

	iris.Listen(":8080")
}
