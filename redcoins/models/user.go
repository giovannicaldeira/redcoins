package models

import (
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	u "github.com/giovanni-rc/redcoins/utils"
	"github.com/jinzhu/gorm"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Name     string `gorm:"not null" json:"name"`
	Birthday string `gorm:"not null" json:"birthday"`
	Token    string `json:"token";sql:"-"`
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

	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	response := u.Message(true, "Usuario criado com sucesso")
	response["user"] = user

	return response
}

func Login(email, password string) map[string]interface{} {

	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email nao encontrado")
		}
		return u.Message(false, "Erro de conexao. Tente novamente")
	}

	if password != user.Password {
		return u.Message(false, "Credenciais invalidas. Tente novamente")
	}

	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))
	user.Token = tokenString

	resp := u.Message(true, "Login realizado com sucesso")
	resp["user"] = user
	return resp
}

func GetUser(email string) *User {

	user := &User{}

	GetDB().Table("users").Where("email = ?", email).First(user)
	if user.Email == "" {
		return nil
	}

	return user
}
