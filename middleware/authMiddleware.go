package middleware

import (
	"NoteAssistant/common"
	"NoteAssistant/model"
	"NoteAssistant/resp"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			resp.Forbidden(ctx)
			ctx.Abort()
			return
		}

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			resp.Forbidden(ctx)
			ctx.Abort()
			return
		}

		DB := common.GetDB()
		user := model.User{ID: claims.UserID}
		DB.First(&user)
		if user.ID == 0 {
			resp.Forbidden(ctx)
			ctx.Abort()
			return
		}

		ctx.Set("UID", user.ID)
		ctx.Next()
	}
}
