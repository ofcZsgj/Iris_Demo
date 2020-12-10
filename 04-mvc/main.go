package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

//自定义的控制器
type UserController struct {
}

func main() {
	app := iris.New()
	//注册自定义控制器处理请求 New()返回一个*Application
	mvc.New(app).Handle(new(UserController))
	//路由组的mvc处理 localhost:8081/user/...
	mvc.Configure(app.Party("/user"), func(ctx *mvc.Application) {
		ctx.Handle(new(UserController))
	})
	app.Run(iris.Addr(":8081"))
}

//自动处理基础的HTTP请求
//Get localhost:8080/user/
func (uc *UserController) Get() string {
	iris.New().Logger().Info("自动处理Get请求")
	return "hello world"
}

//Post
func (uc *UserController) Post() {
	iris.New().Logger().Info("自动处理Post请求")
}

//根据请求类型和请求URL自动匹配处理方法
//Get localhost:8080/user/info
//Get请求/info就需要将方法命名为GetInfo，注意大写
func (uc *UserController) GetInfo() mvc.Result {
	iris.New().Logger().Info("Get请求的路径为/user/info")
	return mvc.Response{
		Object: map[string]interface{}{
			"code":    1,
			"message": "Get Success",
		},
	}
}

//手动指定某个URL用特定方法来执行，不通过自动匹配处理方法执行
func (uc *UserController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET", "/query", "QueryInfo")
}
func (uc *UserController) QueryInfo() mvc.Result {
	iris.New().Logger().Info("Query User Info")
	return mvc.Response{}
}
