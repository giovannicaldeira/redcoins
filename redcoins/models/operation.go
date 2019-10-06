package models

import (
	"time"

	u "redcoins/utils"

	"github.com/jinzhu/gorm"
)

type Operation struct {
	gorm.Model
	Qty    float64   `gorm:"not null" json:"qty"`
	Date   time.Time `gorm:"type:timestamp;not null" json:"date"`
	Value  float64   `gorm:"not null" json:"value"`
	Type   uint      `gorm:"not null" json:"type"` // 0 - compra ; 1 - venda
	UserID uint      `gorm:"not null" json:"user_id"`
}

func (operation *Operation) Create() map[string]interface{} {

	if resp, ok := operation.Validate(); !ok {
		return resp
	}

	operation.Value = operation.Qty * 34500

	GetDB().Create(operation)

	if operation.ID <= 0 {
		return u.Message(false, "Falha ao registrar operacao, erro de conexao.")
	}

	response := u.Message(true, "Operacao registrada com sucesso")
	response["operation"] = operation

	return response
}

func (operation *Operation) Validate() (map[string]interface{}, bool) {

	userTemp := &User{}

	err := GetDB().Table("users").Where("id = ?", operation.UserID).First(userTemp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Erro de conexão. Tente Novamente"), false
	}

	if operation.Qty <= 0 {
		return u.Message(false, "Quantidade invalida para operacao. Tente novamente"), false
	}

	if operation.Type != 0 && operation.Type != 1 {
		return u.Message(false, "Tipo de operacao invalido. Tente novamente"), false
	}

	if userTemp.Email == "" {
		return u.Message(false, "Usuario nao cadastrado"), false
	}

	return u.Message(false, "Request Validado!"), true
}

func GetOperationByUser(id uint64) (map[string]interface{}, []*Operation) {

	userTemp := &User{}

	err := GetDB().Table("users").Where("id = ?", id).First(userTemp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Erro de conexão. Tente Novamente"), nil
	} else if err == gorm.ErrRecordNotFound {
		return u.Message(false, "Usuario nao encontrado. Tente Novamente"), nil
	}

	operations := make([]*Operation, 0)
	err = GetDB().Table("operations").Where("user_id = ?", id).Find(&operations).Error
	if err != nil {
		return u.Message(false, "Erro ao processar sua pesquisa. Tente Novamente"), nil
	}

	response := u.Message(true, "Busca realizada com sucesso")
	response["operations"] = operations

	return response, operations

}

func GetOperationByDate(date string) (map[string]interface{}, []*Operation) {

	operations := make([]*Operation, 0)

	start := date + " 00:00:00"
	end := date + " 23:59:59"

	err := GetDB().Table("operations").Where("date BETWEEN ? and ?", start, end).Find(&operations).Error
	if err != nil {
		return u.Message(false, "Erro ao processar sua pesquisa. Tente Novamente"), nil
	}

	response := u.Message(true, "Busca realizada com sucesso")
	response["operations"] = operations

	return response, operations
}
