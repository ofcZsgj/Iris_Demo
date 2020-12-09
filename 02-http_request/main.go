package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	//GET
	app.Handle("GET", "/userpath", func(context iris.Context) {
		//获取Path
		path := context.Path()
		//日志输出
		app.Logger().Info(path) //[INFO] 2020/12/09 10:40 /userpath
		//写入返回数据：string类型
		context.WriteString("请求路径：" + path) //浏览器返回 请求路径：/userpath
	})
	//处理Get请求并接收参数
	//http://localhost:8080/userinfo?username=zsgj&pwd=helloworld
	app.Handle("GET", "/userinfo", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info(path)
		//获取get请求所携带的参数
		userName := ctx.URLParam("username")
		app.Logger().Info(userName)
		pwd := ctx.URLParam("pwd")
		app.Logger().Info(pwd)
		//返回html数据格式
		ctx.HTML("<h1>" + userName + "," + pwd + "</h1>")
	})
	//通过Get请求返回Json数据
	app.Get("/getJson", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "hello world", "requestCode": 200})
	})
	//POST
	app.Handle("POST", "/login", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info(path)
		//获取请求字段PostValue方法来获取post请求所提交的form表单
		name := ctx.PostValue("name")
		pwd, err := ctx.PostValueInt("pwd")
		if err != nil {
			panic(err.Error())
		}
		app.Logger().Info(name, " ", pwd)
		//返回html数据格式
		ctx.HTML(name)
	})
	//处理Post请求Json格式数据
	// Postman工具选择[{"key":"Content-Type",
	//"value":"application/json","description":""}]
	// 请求内容：{"name": "davie","age": 28}
	app.Post("/postlogin", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info("请求URL：", path)
		//Json数据解析
		var person Person
		if err := ctx.ReadJSON(&person); err != nil {
			panic(err.Error())
		}
		//输出 %#+v
		ctx.Writef("Received:%#+v\n", person) //Received:main.Person{Name:"zsgj", Age:28}
	})

	//PUT
	app.Put("/putinfo", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info("请求url: ", path)
	})
	//DELETE
	app.Delete("/", handler)
	//OPTIONS
	app.Options("/", handler)
	//Trace
	app.Trace("/", handler)
	//CONNECT
	app.Connect("/", handler)
	//HEAD
	app.Head("/", handler)
	//PATCH
	app.Patch("/", handler)
	//任意的http请求方法
	app.Any("/", handler)
	//监听
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

type Person struct {
	Name string `json:"name"` //小写警告：结构字段 'name' 具有 'json' 标记，但未被导出
	Age  int    `json:"age"`
}

func handler(ctx iris.Context) {
	ctx.Writef("method:%s, path:%s", ctx.Method(), ctx.Path())
}
