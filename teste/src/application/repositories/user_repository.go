package repositories

import (
	"log"

	"github.com/brunoeas/poc-golang/teste/src/domain"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {
	err := user.Prepare()

	if err != nil {
		log.Fatalf("Erro na validação do usuário: %v", err)
		return user, err
	}

	err = repo.Db.Create(user).Error

	if err != nil {
		log.Fatalf("Erro ao persistir usuário: %v", err)
		return user, err
	}

	return user, nil
}
