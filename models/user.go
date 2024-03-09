package models

import (
	// "github.com/henrylee2cn/goutil/password"
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	gorm.Model
	UserName string `gorm:"unique"`
	PasswordDigest string
	AvatarFilePath string `gorm:"default:''"`
}

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	fmt.Println(err)
	return err == nil
}
