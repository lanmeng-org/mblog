package router

import (
	"github.com/gin-gonic/gin"
	"MBlog/web/controller"
	"MBlog/web/middleware"
)

func RegisterRoute(engine *gin.Engine) {
	static(engine)

	admin := engine.Group("admin")
	{
		admin.Use(middleware.CheckAuth())
	}


	engine.GET("/", controller.HomeIndex)
}

