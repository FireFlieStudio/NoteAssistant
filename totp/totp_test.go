package totp

import "testing"

func TestTotpCommon(t *testing.T) {
	key := Generate("example@example.com")
	imageB64, _ := GenerateQRCodeB64(key)
	DisplayQRCodeOnTerminal(key.String())
	println(imageB64)
	passCode, _ := GeneratePassCode(key.Secret())
	println(ValidatePassCode(key.Secret(), passCode))
}
