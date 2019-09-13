package main

import (
	"github.com/kataras/iris"
)

func homeHandler(ctx iris.Context) {
	ctx.HTML("<h1>Chao mung den voi Techmaster Viet Nam</h1>")
}

func courseHandler(ctx iris.Context) {
	ctx.HTML("<h1>Chao mung den voi khoa hoc cua Techmaster Viet Nam</h1>")
}

func main() {
	app := iris.New()
	app.Handle("GET", "/", homeHandler)
	app.Handle("GET", "/khoa-hoc", courseHandler)
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
