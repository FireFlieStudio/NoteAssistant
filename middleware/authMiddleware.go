package middleware

import (
	"NoteAssistant/common"
	"NoteAssistant/model"
	"NoteAssistant/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			resp.Send(ctx, http.StatusUnauthorized, gin.H{"Error": "权限不足"})
			ctx.Abort()
			return
		}

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			resp.Send(ctx, http.StatusUnauthorized, gin.H{"Error": "权限不足"})
			ctx.Abort()
			return
		}

		DB := common.GetDB()
		user := model.User{ID: claims.UserID}
		DB.First(&user)
		if user.ID == 0 {
			resp.Send(ctx, http.StatusUnauthorized, gin.H{"Error": "权限不足"})
			ctx.Abort()
			return
		}

		ctx.Set("UID", user.ID)
		ctx.Next()
	}
}
