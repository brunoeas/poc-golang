package domain

import (
	"log"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User - Modelo do usuário
type User struct {
	Base
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255);unique_index"`
	Password string `json:"-" gorm:"type:varchar(255)"`
	Token    string `json:"token" gorm:"type:varchar(255);unique_index"`
}

// NewUser - Retorna uma nova intância de um usuário
func NewUser() *User {
	return &User{}
}

// Prepare - Prepara o objeto usuário e faz as validações necessárias
func (user *User) Prepare() error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Ocorreu um erro ao gerar senha: %v", err)
		return err
	}

	user.Password = string(password)
	user.Token = uuid.NewV4().String()

	err = user.validate()

	if err != nil {
		log.Fatalf("Erro na validação")
		return err
	}

	return nil
}

// validate - Realiza as validações do usuário
func (user *User) validate() error {
	return nil
}
