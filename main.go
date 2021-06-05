package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yp/common"
	"yp/config"
	"yp/routes"
	_ "yp/yunpan"
)

func main(){

	r := gin.Default()
	// 设置日志中间件，主要用于打印请求日志
	r.Use(gin.Logger())

	// 设置Recovery中间件，主要用于拦截paic错误，不至于导致进程崩掉
	r.Use(gin.Recovery())


	//读取配置文件端口
	var port string


	//自定义分隔符
	r.Delims(config.Conf.View.LetfDelim, config.Conf.View.RightDelim)

	//静态文件路径
	r.StaticFS("/public", http.Dir("./static"))

	//加载模板方法
	common.TempLateCommon(r)
	// 加载templates目录下面的所有模版文件，包括子目录
	r.LoadHTMLGlob(config.Conf.View.HtmlGlob)

	//路由注册
	routes.Route(r)


	if config.Conf.Server.Port == 0{
		port = "8080"
	}else{
		port = strconv.Itoa(config.Conf.Server.Port)
	}
	r.Run(":"+port)
}

