package application

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pterm/pterm"
	"gorm.io/gorm"
	"zip.jespersen.chatsper/internal/api/controller"
	"zip.jespersen.chatsper/internal/api/router"
	"zip.jespersen.chatsper/internal/config/file"
	"zip.jespersen.chatsper/internal/platforms/twitch"
	Log "zip.jespersen.chatsper/internal/utils"
)

type Chatsper struct {
	Config   *file.Config
	Database *gorm.DB
	Api      *gin.Engine
}

func (app *Chatsper) Build() {

	Log.Info("Adding Controller to API Server...")

	// Platforms
	twtich := twitch.NewTwitch(app.Config, app.Api, app.Database)
	err := twtich.Init()
	if err != nil {
		fmt.Println(pterm.Red("Error initializing Twitch"))
		os.Exit(0)
	}

	// Controller
	exampleController := controller.NewExampleController(app.Config, app.Database, app.Api)
	frontendController := controller.NewFrontendController(app.Config, app.Database, app.Api)
	authController := controller.NewAuthController(app.Config, app.Database, app.Api)
	routeConfig := router.RouterConfig{
		Api:                app.Api,
		IndexController:    exampleController,
		FrontendController: frontendController,
		AuthController:     authController,
	}

	Log.Info("Running Setup for API Router...")
	routeConfig.Setup()
	Log.Info("Loaded API Router for Chatsper")
}
