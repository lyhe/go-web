package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.tmpl", gin.H{
		"doubi": "db",
	})
}
