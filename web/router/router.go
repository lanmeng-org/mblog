package router

import (
	"github.com/gin-gonic/gin"
	"mblog/web/controller"
	"mblog/web/middleware"
)

func RegisterRoute(engine *gin.Engine) {
	static(engine)

	admin := engine.Group("admin")
	{
		admin.Use(middleware.CheckAuth())
	}


	engine.GET("/", controller.HomeIndex)
}

