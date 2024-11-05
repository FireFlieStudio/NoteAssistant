package request

type GetQRCode struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}

func (getQRCode GetQRCode) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Email.required": "邮箱不能为空",
	}
}

type Register struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	PassCode string `form:"passcode" json:"passcode" binding:"required"`
	Secret   string `form:"secret" json:"secret" binding:"required"`
}

func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Email.required":    "邮箱不能为空",
		"Secret.required":   "密钥不能为空",
		"PassCode.required": "密码不能为空",
	}
}

type Login struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	PassCode string `form:"passcode" json:"passcode" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Email.required":    "邮箱不能为空",
		"PassCode.required": "密码不能为空",
	}
}
