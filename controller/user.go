package controller

import (
	"NoteAssistant/common"
	"NoteAssistant/common/request"
	"NoteAssistant/model"
	"NoteAssistant/resp"
	"NoteAssistant/totp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterWithTotp(ctx *gin.Context) {
	var form request.GetQRCode
	if err := ctx.Bind(&form); err != nil {
		resp.ValidateError(ctx, form, err)
		return
	}

	key := totp.Generate(form.Email)
	qrCodeImageB64, err := totp.GenerateQRCodeB64(key)
	if err != nil {
		resp.InternalError(ctx, err)
		return
	}
	resp.Send(ctx, http.StatusOK, gin.H{"email": form.Email, "QRCodeImageB64": qrCodeImageB64, "secret": key.Secret()})
}

func LoginWithTotp(ctx *gin.Context) {
	var form request.Login
	if err := ctx.Bind(&form); err != nil {
		resp.ValidateError(ctx, form, err)
		return
	}

	user := model.User{Email: form.Email}
	common.GetDB().First(&user)
	if user.ID == 0 || len(user.Secret) == 0 {
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
		resp.ValidateError(ctx, form, err)
		return
	}

	user := &model.User{
		Name:   "ChangeMePlz",
		Email:  form.Email,
		Secret: form.Secret,
	}
	if userExists(user) {
		resp.Failed(ctx, gin.H{"Error": "注册失败，用户已存在!"})
		return
	}

	if !totp.ValidatePassCode(form.Secret, form.PassCode) {
		resp.Forbidden(ctx)
		return
	}

	common.GetDB().Create(user)
	resp.Success(ctx, gin.H{"name": user.Name, "email": user.Email})
}

func Info(ctx *gin.Context) {
	UID, ok := ctx.Get("UID")
	if !ok {
		resp.Forbidden(ctx)
		return
	}
	resp.Success(ctx, gin.H{"UID": UID})
}

func userExists(user *model.User) bool {
	common.GetDB().First(user)
	return user.ID != 0
}
