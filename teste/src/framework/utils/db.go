package utils

import (
	"log"
	"os"

	"github.com/brunoeas/poc-golang/teste/src/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Erro carregando o .env: %v", err)
	}

	dsn := os.Getenv("dsn")

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
		panic(err)
	}

	// defer db.Close()

	db.AutoMigrate(&domain.User{})

	return db
}
