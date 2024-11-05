package resp

import (
	"NoteAssistant/common/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Send(ctx *gin.Context, code int, data gin.H) {
	ctx.JSON(code, data)
}

func Success(ctx *gin.Context, data gin.H) {
	Send(ctx, http.StatusOK, data)
}

func Failed(ctx *gin.Context, data gin.H) {
	Send(ctx, http.StatusBadRequest, data)
}

func BadRequest(ctx *gin.Context, err error) {
	Send(ctx, http.StatusBadRequest, gin.H{"Error": err.Error()})
}

func InternalError(ctx *gin.Context, err error) {
	Send(ctx, http.StatusInternalServerError, gin.H{"Msg": "服务器内部问题", "Error": err.Error()})
}

func Forbidden(ctx *gin.Context) {
	Send(ctx, http.StatusForbidden, gin.H{"Error": "权限不足"})
}

func ValidateError(ctx *gin.Context, validator request.Validator) {
	Send(ctx, http.StatusBadRequest, gin.H{"Error": validator.GetMessages()})
}
