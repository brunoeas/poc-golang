package domain

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" sql:"index"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedAt", time.Now())
	if err != nil {
		log.Fatalf("Erro ao settar a propriedade CreatedAt: %v", err)
	}

	err = scope.SetColumn("ID", uuid.NewV4().String())
	if err != nil {
		log.Fatalf("Erro ao settar a propriedade ID: %v", err)
	}

	return nil
}
