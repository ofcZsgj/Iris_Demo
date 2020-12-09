package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	//创建Application结构体对象指针
	app := iris.New()
	app.Logger().SetLevel("debug")
	//添加两个内置处理程序
	//可以从任何与http相关的panic中回复，并将请求记录到终端
	app.Use(recover.New())
	app.Use(logger.New())
	//请求方法：GET
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})
	//请求方法：GET
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})
	//请求方法：GET
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})
	//端口监听 loacalhost:8080 loacalhost:8080/ping loacalhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
