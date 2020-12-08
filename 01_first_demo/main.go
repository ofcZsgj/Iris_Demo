package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	//创建Application结构体对象指针
	app := iris.New()
	//端口监听
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
