package main

import (
	"github.com/gin-gonic/gin"
	"go-web/router"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	// 注册一个路径，gin 加载模板的时候会从该目录查找
	// 参数是一个匹配字符，如 views/*/* 的意思是 模板目录有两层
	// gin 在启动时会自动把该目录的文件编译一次缓存，不用担心效率问题
	r.LoadHTMLGlob("views/*/*")
	r.GET("/", router.Index)
	r.NoRoute(router.NotFound)

	_ = r.Run(":80")
}
