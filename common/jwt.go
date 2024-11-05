package common

import (
	"NoteAssistant/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var tokenExpire time.Duration
var secret []byte
var issuer string

func setDefaultJwtConfig() {
	viper.SetDefault("jwt.tokenExpire", 30*time.Minute)
	viper.SetDefault("jwt.secret", "a secret")
	viper.SetDefault("jwt.issuer", "tomato")
}

func init() {
	setDefaultJwtConfig()
	tokenExpire = viper.GetDuration("jwt.tokenExpire")
	secret = []byte(viper.GetString("jwt.secret"))
	issuer = viper.GetString("jwt.issuer")
}

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(tokenExpire)
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString[7:], claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	return token, claims, err
}

func Encrypt(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

func Verify(inputPass, DBPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(DBPass), []byte(inputPass))
}
