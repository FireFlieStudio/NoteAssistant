package resp

import (
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
