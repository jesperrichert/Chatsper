package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pterm/pterm"
	"zip.jespersen.chatsper/internal/config/file"
	Log "zip.jespersen.chatsper/internal/utils/log"
	"zip.jespersen.chatsper/shared/static"
)

func Setup() {
	_, fileCheck := os.ReadFile("./config.yml")
	if fileCheck != nil {
		config := file.ConfigTemplate

		fmt.Println(pterm.LightCyan("\n> Setup Chatsper"))
		pterm.Underscore.Println("You are running Chatsper v." + static.Version)
		fmt.Println("\nLet's start with the first Setup")
		fmt.Println()

		fmt.Println("Select your Favourite Platform (TWITCH)")
		fmt.Print("> ")
		text := readString()

		if text == "TWITCH" {
			fmt.Println(pterm.LightGreen("Please enter your Twitch-Application-Client-ID"))
			fmt.Println(pterm.Gray("How to: https://doc.jespersen.zip/s/chatspers/doc/get-twtich-application-client-id-B2zM9vCnHb"))
			fmt.Print("> ")
			text = readString()

			if len(text) == 0 {
				setupError("Please enter your Clinet-ID")
				os.Exit(0)
			} else {
				config.TwitchPlatform.ApplicationId = text
			}

			fmt.Println(pterm.LightGreen("Please enter your Twitch-Application-Client-Secret"))
			fmt.Println(pterm.Gray("How to: https://doc.jespersen.zip/s/chatspers/doc/get-twtich-application-client-secret-lwmc7GJaKE"))
			fmt.Print("> ")
			text = readString()

			if len(text) == 0 {
				setupError("Please enter your Clinet-Secret")
				os.Exit(0)
			} else {
				config.TwitchPlatform.ApplicationSecret = text
			}

			fmt.Println(pterm.Gray("Done. Next: Database"))
			fmt.Println(pterm.LightGreen("Select your Database Type (SQLITE, POSTGRES)"))
			fmt.Print("> ")
			text = readString()

			switch text {
			case "SQLITE":
				config.Database.Type = "SQLITE"
			case "POSTGRES":
				config.Database.Type = "POSTGRES"
				fmt.Println(pterm.Gray("POSTGRES Databse needs a Connection Url"))
				fmt.Println(pterm.LightGreen("Enter your Connection URl (e.g. postgres://postgres:postgres@localhost:5432/postgres)"))
				fmt.Print("> ")
				text = readString()

				if len(text) == 0 {
					setupError("No url provided!")
					os.Exit(0)
				} else {
					config.Database.ConnectionUrl = text
				}
			default:
				setupError("No database type selected!")
			}

			fmt.Println()
			pterm.Bold.Printfln("%s", pterm.Green("Setup successfully executed! Chatsper now will follow the default Startup and init with your setup... Please wait."))
			fmt.Println()

		} else {
			setupError("You need to select a Platform Type!!!")
		}
	} else {
		Log.Info("Chatsper setup is up-to-date.")
	}

}

func readString() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.ReplaceAll(text, "\n", "")
	return text
}

func setupError(message string) {
	fmt.Println(pterm.Red(message))
	os.Exit(500)
}
