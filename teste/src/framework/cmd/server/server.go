package main

import (
	"fmt"
	"log"

	"github.com/brunoeas/poc-golang/teste/src/application/repositories"
	"github.com/brunoeas/poc-golang/teste/src/domain"
	"github.com/brunoeas/poc-golang/teste/src/framework/utils"
)

func main() {

	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Bruno",
		Email:    "brunin@email.com",
		Password: "adm123",
	}

	userRepo := repositories.UserRepositoryDb{Db: db}
	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(result)
}
