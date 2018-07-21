package main

import (
	_ "AiJia/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
	"net/http"
	_"AiJia/models"
)

func main() {
	ignoreStaticPath()
	beego.Run()

}

func ignoreStaticPath(){
	//
     //请求url中没有api字段。
	beego.InsertFilter("/",beego.BeforeRouter,TransparentStatic)
	beego.InsertFilter("/*",beego.BeforeRouter,TransparentStatic)
}

func TransparentStatic(ctx *context.Context){
	orpath:= ctx.Request.URL.Path
	beego.Debug("request url:",orpath)

	//如果请求url还有api字段，说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath,"api")>=0{
		return
	}
	http.ServeFile(ctx.ResponseWriter,ctx.Request,"static/html/"+ctx.Request.URL.Path)
}