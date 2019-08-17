package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context){
		ctx.HTML("<h1>Welcome</h1>")
	})
	app.Run(iris.Addr(":8080"))
}