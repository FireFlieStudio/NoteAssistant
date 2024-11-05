package totp

import (
	"NoteAssistant/logger"
	"bytes"
	"crypto/md5"
	"encoding/base32"
	"encoding/base64"
	"errors"
	"github.com/mdp/qrterminal/v3"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/spf13/viper"
	"image/png"
	"os"
	"time"
)

func setDefaultJwtConfig() {
	viper.SetDefault("jwt.issuer", "example.com")
}

var (
	issuer string
)

func init() {
	setDefaultJwtConfig()
	issuer = viper.GetString("jwt.issuer")
}

type Totp struct{}

func Generate(email string) *otp.Key {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: email,
		Period:      30,
	})
	if err != nil {
		logger.Error("Totp Generate Error:", err)
		return nil
	}
	return key
}

func GeneratePassCode(secret string) (string, error) {
	if len(secret) == 0 {
		logger.Error("Totp Generate PassCode Error, Secret is empty")
		return "", errors.New("totp Generate PassCode Error, Secret is empty")
	}
	passCode, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		logger.Error("Totp Generate PassCode Error:", err)
		return "", err
	}
	return passCode, nil
}

func ValidatePassCode(secret, passCode string) bool {
	return totp.Validate(passCode, secret)
}

func DisplayQRCodeOnTerminal(qrCodeUrl string) {
	qrterminal.Generate(qrCodeUrl, qrterminal.L, os.Stdout)
}

func GenerateQRCodeB64(key *otp.Key) (string, error) {
	var buf bytes.Buffer
	img, err := key.Image(256, 256)
	if err != nil {
		logger.Error("Generate QRCode Error:", err)
		return "", err
	}
	png.Encode(&buf, img)
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func getSecretByEmail(email string) string {
	mdtVal := md5.Sum([]byte(email))
	b32NoPadding := base32.StdEncoding.WithPadding(base32.NoPadding)
	return b32NoPadding.EncodeToString(mdtVal[:])
}

func ValidatePassCodeByEmail(email, passCode string) bool {
	secret := getSecretByEmail(email)
	return ValidatePassCode(secret, passCode)
}
