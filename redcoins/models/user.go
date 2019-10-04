package models

import (
	"strings"

	u "github.com/giovanni-rc/redcoins/utils"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Name     string `gorm:"not null" json:"name"`
	Birthday string `gorm:"not null" json:"birthday"`
}

func (user *User) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email já cadastrado"), false
	}

	userTemp := &User{}

	err := GetDB().Table("users").Where("email = ?", user.Email).First(user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Erro de conexão. Tente Novamente"), false
	}

	if userTemp.Email != "" {
		return u.Message(false, "Email já cadastrado"), false
	}

	return u.Message(false, "Request Validado!"), true
}

func (user *User) Create() map[string]interface{} {

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	GetDB().Create(user)

	if user.ID <= 0 {
		return u.Message(false, "Falha ao criar usuario, erro de conexao.")
	}

	response := u.Message(true, "Usuario criado com sucesso")
	response["user"] = user

	return response
}

func GetUser(email string) *User {

	user := &User{}

	GetDB().Table("users").Where("email = ?", email).First(user)
	if user.Email == "" {
		return nil
	}

	return user
}
