package controller

import (
	"NoteAssistant/common"
	"NoteAssistant/model"
	"NoteAssistant/resp"
	"NoteAssistant/totp"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func RegisterWithTotp(ctx *gin.Context) {
	email := ctx.PostForm("email")
	if err := validator.New().Var(email, "required,email"); err != nil {
		resp.Send(ctx, http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	key := totp.Generate(email)
	qrCodeImageB64, err := totp.GenerateQRCodeB64(key)
	if err != nil {
		resp.Send(ctx, http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	resp.Send(ctx, http.StatusOK, gin.H{"email": email, "QRCodeImageB64": qrCodeImageB64, "secret": key.Secret()})
}

func LoginWithTotp(ctx *gin.Context) {
	passCode := ctx.PostForm("passCode")
	email := ctx.PostForm("email")
	if err := validator.New().Var(passCode, "required,len=6"); err != nil {
		resp.Send(ctx, http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	if err := validator.New().Var(email, "required,email"); err != nil {
		resp.Send(ctx, http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	DB := common.GetDB()
	user := model.User{Email: email}
	fmt.Println(user)
	DB.First(&user)
	if len(user.Secret) == 0 {
		resp.Send(ctx, http.StatusBadRequest, gin.H{"Error": "用户不存在或密钥未绑定"})
		return
	}

	if !totp.ValidatePassCode(user.Secret, passCode) {
		ctx.JSON(http.StatusForbidden, gin.H{"msg": "权限不足"})
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		resp.Send(ctx, http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	resp.Send(ctx, http.StatusOK, gin.H{"token": token})
}

func BindSecret(ctx *gin.Context) {
	passCode := ctx.PostForm("passCode")
	email := ctx.PostForm("email")
	secret := ctx.PostForm("secret")
	if err := validator.New().Var(passCode, "required,len=6"); err != nil {
		resp.Send(ctx, http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := validator.New().Var(email, "required,email"); err != nil {
		resp.Send(ctx, http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := validator.New().Var(secret, "required"); err != nil {
		resp.Send(ctx, http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if !totp.ValidatePassCode(secret, passCode) {
		resp.Send(ctx, http.StatusBadRequest, gin.H{"Error": "Totp验证不通过"})
		return
	}

	DB := common.GetDB()
	DB.Create(&model.User{
		Name:   "ChangeMePlz",
		Email:  email,
		Secret: secret,
	})
}

func Info(ctx *gin.Context) {
	UID, ok := ctx.Get("UID")
	if !ok {
		resp.Failed(ctx, nil)
		return
	}
	resp.Success(ctx, gin.H{"UID": UID})
}