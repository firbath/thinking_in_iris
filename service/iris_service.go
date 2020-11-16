/*
@Time : 2020/11/15 18:32
@Author : Firbath
@File : iris_service.go
@Software: GoLand
*/
package service

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func IrisRun(host string, port int) {
	app := iris.Default()
	registerController(app)
	registerViews(app)
	//app.Listen(fmt.Sprintf("%s:%d", host, port))

	app.Run(iris.Addr(fmt.Sprintf("%s:%d", host, port)))
}

func registerController(app *iris.Application) {

	mvc.New(app.Party("/")).Handle(new(RootController))
	mvc.New(app.Party("/books")).Handle(new(BookController))

}

func registerViews(app *iris.Application) {
	// 注册静态文件路由
	app.HandleDir("/img", "./assets/img")
	app.HandleDir("/css", "./assets/css")
	app.HandleDir("/javascript", "./assets/javascript")

	tmpl := iris.HTML("./assets/templates", ".html")
	app.RegisterView(tmpl)

	app.Get("/", func(ctx iris.Context) {
		// 绑定： {{.message}}　为　"Hello world!"
		ctx.ViewData("message", "Hello world!")
		// 渲染模板文件： ./assets/templates/demo.html
		ctx.View("demo.html")
	})

}
