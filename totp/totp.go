package totp

import (
	"NoteAssistant/logger"
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/mdp/qrterminal/v3"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"image/png"
	"os"
	"time"
)

type Totp struct {
	Secret         string
	QRCodeUrl      string
	QRCodeImageB64 string
}

func NewTotp() *Totp {
	return &Totp{}
}

func (t *Totp) Generate(accountName string) error {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Tomato",
		AccountName: accountName,
		Period:      30,
	})
	if err != nil {
		logger.Error("Totp Generate Error:", err)
		return err
	}

	imageB64, err := t.generateQRCode(key)
	if err != nil {
		logger.Error("Totp GenerateQRCode Error:", err)
		return err
	}

	t.QRCodeImageB64 = imageB64
	t.Secret = key.Secret()
	t.QRCodeUrl = key.String()

	return nil
}

func (t *Totp) GeneratePassCode() (string, error) {
	if t.Secret == "" {
		logger.Error("Totp Generate PassCode Error, Secret is empty")
		return "", errors.New("totp Generate PassCode Error, Secret is empty")
	}
	passCode, err := totp.GenerateCode(t.Secret, time.Now())
	if err != nil {
		logger.Error("Totp Generate PassCode Error:", err)
		return "", err
	}
	return passCode, nil
}

func (t *Totp) ValidatePassCode(passCode string) bool {
	return totp.Validate(passCode, t.Secret)
}

func (t *Totp) DisplayQRCodeOnTerminal() {
	qrterminal.Generate(t.QRCodeUrl, qrterminal.L, os.Stdout)
}

func (t *Totp) generateQRCode(key *otp.Key) (string, error) {
	var buf bytes.Buffer
	img, err := key.Image(256, 256)
	if err != nil {
		logger.Error("Generate QRCode Error:", err)
		return "", err
	}
	png.Encode(&buf, img)
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
