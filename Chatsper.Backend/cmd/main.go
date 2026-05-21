package main

import (
	"fmt"
	"log"

	"github.com/pterm/pterm"
	"zip.jespersen.chatsper/Chatsper.Backend/internal/config/application"
	"zip.jespersen.chatsper/Chatsper.Backend/internal/config/file"
)

func main() {
	configuration, err := file.LoadFromFile("./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	api := application.NewGin()
	database := application.NewDatabase(configuration.Database.Type, configuration.Database.ConnectionUrl)

	chatsper := &application.Chatsper{
		Config:   configuration,
		Database: database,
		Api:      api,
	}

	chatsper.Build()

	err = api.Run(":3000")

	fmt.Println(pterm.Green("Chatsper Application is up and running!"))
	fmt.Println(pterm.Green("---------------------------------------"))
	fmt.Println(pterm.Green("API Endpoint: http://localhost:3000/api"))
	fmt.Println(pterm.Green("API Dashboard: http://localhost:3000"))
	fmt.Println(pterm.Green("---------------------------------------"))

	if err != nil {
		return
	}
}
