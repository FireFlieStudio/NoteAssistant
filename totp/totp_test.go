package totp

import "testing"

func TestTotpCommon(t *testing.T) {
	totp := NewTotp()
	if err := totp.Generate("admin@example.com"); err != nil {
		panic(err)
		return
	}
	passcode, err := totp.GeneratePassCode()
	if err != nil {
		panic(err)
		return
	}
	println(passcode)
	println(totp.ValidatePassCode(passcode))
	println(totp.QRCodeImageB64)
	totp.DisplayQRCodeOnTerminal()
}
