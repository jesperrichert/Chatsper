package main

import (
	"fmt"
	"log"

	"github.com/pterm/pterm"
	"zip.jespersen.chatsper/internal/config/application"
	"zip.jespersen.chatsper/internal/config/file"
	"zip.jespersen.chatsper/shared/generation"
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

	fmt.Println(pterm.Blue("\nChatsper Application is up and running!"))
	fmt.Println(pterm.Gray("---------------------------------------"))
	fmt.Println(pterm.Green("API Endpoint: http://localhost:3000/api"))
	fmt.Println(pterm.Green("API Dashboard: http://localhost:3000"))

	if configuration.TwitchPlatform.BotAccountUserId == 0 {
		fmt.Println(pterm.Red("You need to set a Bot User Id to tell Chatsper what account is the Bot to use."))
		fmt.Println(pterm.Gray("Use this url to get the user id: " + generation.GenerateTwitchAuthUrl(
			configuration.TwitchPlatform.ApplicationId,
			configuration.TwitchPlatform.RedirectUrl,
			[]string{
				"user:read:chat",
				"user:write:chat",
				"user:bot",
				"channel:bot",
				"channel:read:subscriptions",
			},
			"botAuth",
		)))
	}

	fmt.Println(pterm.Gray("---------------------------------------\n"))

	err = api.Run(":3000")
	if err != nil {
		return
	}
}
