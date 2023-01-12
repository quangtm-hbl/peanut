package main

import (
	"fmt"
	"log"
	"peanut/config"
	"peanut/infra"
	"peanut/pkg/i18n"

	"gorm.io/gorm"
)

func main() {
	fmt.Println("---- Hello world! ----")

	config.Setup()

	i18n.SetupI18n()

	dbClient := dbConnect()
	infra.Migration(dbClient)
	server := infra.SetupServer(dbClient)

	server.Router.Run(":8081")
}

func dbConnect() *gorm.DB {
	db, err := infra.ConnectMySQL()
	if err != nil {
		log.Fatal("[main] DB connect error: ", err)
	}
	return db
}
