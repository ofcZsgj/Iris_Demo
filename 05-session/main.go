package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/boltdb"
)

var (
	USERNAME = "userName"
	ISLOGIN = "isLogin"
)

func main() {
	app := iris.New()
	sessionID := "mySession"
	//创建session并使用
	sess := sessions.New(sessions.Config{
		Cookie: sessionID,
	})
	//用户登录功能
	app.Post("/login", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info("请求Path：", path)
		userName := ctx.PostValue("name")
		passwd := ctx.PostValue("pwd")
		if userName == "zsgj" && passwd == "123" {
			session := sess.Start(ctx)
			//session存储用户名
			session.Set(USERNAME, userName)
			//session存储登录状态
			session.Set(ISLOGIN, true)
			ctx.WriteString("login success!")
		} else {
			session := sess.Start(ctx)
			session.Set(ISLOGIN, false)
			ctx.WriteString("please sign up")
		}
	})
	//用户退出登录功能
	app.Get("/logout", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info("退出登录path: ", path)
		session := sess.Start(ctx)
		//删除session
		session.Delete(ISLOGIN)
		session.Delete(USERNAME)
		ctx.WriteString("login out success")
	})
	//根据session记录的用户是否登录信息来查询
	app.Handle("GET", "/query", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info(path)
		session := sess.Start(ctx)

		isLogin, err := session.GetBoolean(ISLOGIN)
		if err != nil {
			ctx.WriteString("please sign up")
			return
		}
		if isLogin {
			app.Logger().Info("sign success and query success")
			ctx.WriteString("query success")
		} else {
			app.Logger().Info("sign failed and quer failed")
			ctx.WriteString("query success")
		}
	})

	//session 绑定数据库
	db, err := boltdb.New("sessions.db", 0600)
	if err != nil {
		panic(err.Error())
	}
	//程序中断时将数据库关闭
	iris.RegisterOnInterrupt(func() {
		defer db.Close()
	})
	//session和db绑定
	sess.UseDatabase(db)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}