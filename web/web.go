package web

import (
	"github.com/gin-gonic/gin"
	"mblog/util"
	"mblog/web/router"
)

var engine *gin.Engine

func Run() {
	if util.BlogConfig.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine = gin.New()
	engine.Use(gin.Recovery())
	engine.HTMLRender = util.GenRenderer()
	router.RegisterRoute(engine)
	engine.Run(util.BlogConfig.Web.Listen)
}
