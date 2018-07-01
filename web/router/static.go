package router

import (
	"github.com/gin-gonic/gin"
	"mblog/util"
)

func static(engine *gin.Engine) {
	for k,v:=range util.BlogConfig.Web.Static.Dirs {
		engine.Static(k, v)
	}

	for k,v:=range util.BlogConfig.Web.Static.Files {
		engine.StaticFile(k, v)
	}
}