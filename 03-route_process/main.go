package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	//      http://localhost:8080?date=20190310&city=beijing
	//GET： http://localhost:8080/weather/2019-03-10/beijing
	//      http://localhost:8080/weather/2019-03-11/beijing
	//      http://localhost:8080/weather/2019-03-11/tianjin
	app.Get("/weather/{date}/{city}", func(ctx iris.Context) {
		path := ctx.Path()
		date := ctx.Params().Get("date")
		city := ctx.Params().Get("city")
		ctx.WriteString(path + " , " + date + " , " + city)
	})

	//Get正则表达式路由
	// 使用：context.Params().Get("name") 获取正则表达式变量
	// 请求1：/hello/1  /hello/2  /hello/3 /hello/10000
	//正则表达式：{name}
	app.Get("/hello/{name}", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info(path)
		//获取正则表达式变量内容值
		name := ctx.Params().Get("name")
		ctx.HTML("<h1>" + name + "</h1>")
	})

	//自定义正则表达式变量路由请求 {val:uint64}进行变量类型限制
	app.Get("/api/users/{isLogin:bool}", func(ctx iris.Context) {
		isLogin, err := ctx.Params().GetBool("isLogin")
		if err != nil {
			ctx.StatusCode(iris.StatusNonAuthoritativeInfo)
		}
		if isLogin {
			ctx.WriteString("已登录")
		} else {
			ctx.WriteString("未登录")
		}
	})

	//用户模块users
	//  xxx/users/register 注册
	//  xxx/users/login  登录
	//  xxx/users/info   获取用户信息

	//路由组请求
	userParty := app.Party("/users", func(ctx iris.Context) {
		//处理下一级请求
		ctx.Next()
	})

	//路由组下面的下面一级请求
	//../users/register
	userParty.Get("/register", func(ctx iris.Context) {
		app.Logger().Info("register")
		ctx.HTML("<h1>register</h1>")
	})

	usersRouter := app.Party("/admin", userMiddleware)

	//Done
	usersRouter.Done(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("response sent to" + ctx.Path())
	})

	usersRouter.Get("/info", func(ctx iris.Context) {
		ctx.HTML("<h1>用户信息</h1>")
		ctx.Next()
	})

	usersRouter.Get("/query", func(ctx iris.Context) {
		ctx.HTML("<h1>查询信息</h1>")
	})

	app.Run(iris.Addr(":8081"), iris.WithoutServerError(iris.ErrServerClosed))
}

//用户路由中间件
func userMiddleware(ctx iris.Context) {
	ctx.Next()
}
