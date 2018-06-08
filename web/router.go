package web

import (
	"github.com/gin-gonic/gin"
	"MBlog/util"
)

type WebHandler interface {
	Handle(c *gin.Context)
}

var r *gin.Engine

func init() {
	if util.BlogConfig.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r = gin.New()
	r.Use(gin.Recovery())

	r.SetFuncMap(util.GetFuncMap())
	r.LoadHTMLFiles(util.GetTemplateFiles()...)

	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

}

func Run() {
	r.Run(util.BlogConfig.Web.Listen)
}

