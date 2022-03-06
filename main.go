package main

import (
	"log"
	"tests/buisness/controller"
	"tests/buisness/router"
	"tests/entities/aws"
	"tests/entities/repository"

	"github.com/labstack/echo/v4"
)

func main() {

	server := echo.New()

	awsConnection := aws.Dynamodb()
	dnaRepository := repository.NewDnaRepository(awsConnection)
	dnaRepository.CreateTables()

	mutantController := controller.NewMutantController(dnaRepository)

	router := router.NewRouter(*server, mutantController)
	log.Fatal(router.Start())

}
