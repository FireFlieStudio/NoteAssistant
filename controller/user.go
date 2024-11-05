package controller

import (
	"NoteAssistant/common"
	"NoteAssistant/common/request"
	"NoteAssistant/model"
	"NoteAssistant/resp"
	"NoteAssistant/totp"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func RegisterWithTotp(ctx *gin.Context) {
	email := ctx.Query("email")
	if err := validator.New().Var(email, "required,email"); err != nil {
		resp.Send(ctx, http.StatusBadRequest, gin.H{"Error": "邮箱不能为空或邮箱格式有误"})
		return
	}

	key := totp.Generate(email)
	qrCodeImageB64, err := totp.GenerateQRCodeB64(key)
	if err != nil {
		resp.BadRequest(ctx, err)
		return
	}
	resp.Send(ctx, http.StatusOK, gin.H{"email": email, "QRCodeImageB64": qrCodeImageB64, "secret": key.Secret()})
}

func LoginWithTotp(ctx *gin.Context) {
	var form request.Login
	if err := ctx.Bind(&form); err != nil {
		resp.ValidateError(ctx, form)
		return
	}

	user := model.User{Email: form.Email}
	common.GetDB().First(&user)

	if user.ID == 0 {
		resp.Failed(ctx, gin.H{"Error": "用户不存在"})
		return
	}

	if len(user.Secret) == 0 {
		resp.Send(ctx, http.StatusBadRequest, gin.H{"Error": "用户不存在或密钥未绑定"})
		return
	}

	if !totp.ValidatePassCode(user.Secret, form.PassCode) {
		resp.Forbidden(ctx)
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		resp.InternalError(ctx, err)
		return
	}

	resp.Success(ctx, gin.H{"token": token})
}

func BindSecret(ctx *gin.Context) {
	var form request.Register
	if err := ctx.Bind(&form); err != nil {
		resp.ValidateError(ctx, form)
		return
	}

	if !totp.ValidatePassCode(form.Secret, form.PassCode) {
		resp.Forbidden(ctx)
		return
	}

	common.GetDB().Create(&model.User{
		Name:   "ChangeMePlz",
		Email:  form.Email,
		Secret: form.Secret,
	})
	resp.Success(ctx, gin.H{"msg": "注册成功"})
}

func Info(ctx *gin.Context) {
	UID, ok := ctx.Get("UID")
	if !ok {
		resp.Forbidden(ctx)
		return
	}
	resp.Success(ctx, gin.H{"UID": UID})
}
