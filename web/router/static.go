package router

import "github.com/gin-gonic/gin"

func static(engine *gin.Engine) {
	engine.Static("/static", "./static")
	engine.StaticFile("/favicon.ico", "./static/favicon.ico")
}