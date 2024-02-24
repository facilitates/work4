package utils

import (
	"time"
	"path/filepath"
	//"github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt"
	"os"
	"fmt"
)

var JWTsecret = []byte("ABAB")

type Claims struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		ID:       id,
		UserName: username,
		PassWord: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "work4",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JWTsecret)
	return token, err
}

func ParseToken(token string)(*Claims, error){
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token)(interface{}, error){
		return JWTsecret, nil
	})
	if tokenClaims != nil{
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid{
			return claims, nil
		}
	}
	return nil, err
}

func ParseAvatarExt(filename string) bool {
	fileExt := filepath.Ext(filename)
	if fileExt != ".jpeg" && fileExt != ".jpg" && fileExt != ".png" {
        return true
    }
	return false
}

func ParseVideoExt(filename string) bool {
	fileExt := filepath.Ext(filename)
	if fileExt != ".mp4" {
        return true
    }
	return false
}

func CreateFolder(username string) error {
	filepath := "./upload/avatar/"+username
	err := os.Mkdir(filepath, 0755)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
