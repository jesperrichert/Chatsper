package application

import (
	"zip.jespersen.chatsper/Chatsper.Backend/internal/api/controller"
	"zip.jespersen.chatsper/Chatsper.Backend/internal/api/router"
	shared "zip.jespersen.chatsper/Chatsper.Backend/shared/application"
)

func (app *shared.Chatsper) Build() {

	//Register Controller
	exampleController := controller.NewExampleController(*app)

	routeConfig := router.RouterConfig{
		Chatsper:          *app,
		ExampleController: exampleController,
	}

	routeConfig.Setup()
}
