package router

import "github.com/gin-gonic/gin"

func Collector(r *gin.Engine) *gin.Engine {
	UserRouter(r)
	return r
}
