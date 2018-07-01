package controller

import "github.com/gin-gonic/gin"

func HomeIndex(c *gin.Context) {
	c.HTML(200, "member/login.html", gin.H{"title":"test home"})
}