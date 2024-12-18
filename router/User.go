package router

import (
	"NoteAssistant/controller"
	"NoteAssistant/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) *gin.RouterGroup {
	userGroup := r.Group("/user")

	userGroup.GET("register", controller.RegisterWithTotp)
	userGroup.POST("login", controller.LoginWithTotp)
	userGroup.POST("bind", controller.BindSecret)
	userGroup.POST("info", middleware.AuthMiddleware(), controller.Info)

	return userGroup
}
