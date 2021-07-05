package routes

import (
	"github.com/gin-gonic/gin"
	"yp/controller"
)

func Route(r *gin.Engine){
	main :=  r.Group("/main")
	{
		main.GET("/index", controller.MainView)
		main.GET("/download", controller.Download)
		main.GET("/preview", controller.Preview)
		main.GET("/multiDownload", controller.MultiDownload)

	}

	r.GET("/", controller.MainView)


}

