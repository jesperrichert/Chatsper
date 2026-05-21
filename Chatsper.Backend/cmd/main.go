package main

import (
	"fmt"
	"log"

	"zip.jespersen.chatsper/Chatsper.Backend/internal/config"
)

type Chatsper struct {
	Config *config.Config
}

func main() {
	configuration, err := config.LoadFromFile("./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	chatsper := &Chatsper{
		Config: configuration,
	}
	fmt.Println(chatsper.Config.Version)
}
