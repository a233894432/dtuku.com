/*  project: app
    author:diogo
    time: 2016/11/22-14:25
*/
package main

import (
	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	//"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
)

func main() {
	// set the recovery
	iris.Use(recovery.New())

	// set static folder(s)
	iris.Static("/public", "./src/public", 1)

	// set the global middlewares
	iris.Use(logger.New())

	// register Routes
	registerRoutes()

	// start the server
	iris.Listen(":8080")

}

func registerRoutes() {
	// register index using a 'Handler'
	//iris.Handle("GET", "/", routes.Index())

	// this is other way to declare a route
	// using a 'HandlerFunc'
	iris.Get("/", func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, "You should see the favicon now at the side of your browser, if not please refresh or clear the browser's cache.")
	})

	// Dynamic route

	//iris.Get("/profile/:username", routes.Profile)("user-profile")
	//// user-profile is the custom,optional, route's Name: with this we can use the {{ url "user-profile" $username}} inside userlist.html

	//iris.Get("/all", routes.UserList)
}
