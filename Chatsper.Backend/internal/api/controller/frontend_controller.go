package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zip.jespersen.chatsper/internal/config/file"
	"zip.jespersen.chatsper/shared/generation"
	"zip.jespersen.chatsper/shared/static"
)

type FrontendController struct {
	Config   *file.Config
	Database *gorm.DB
	Api      *gin.Engine
}

func NewFrontendController(
	config *file.Config,
	database *gorm.DB,
	api *gin.Engine,
) *FrontendController {
	return &FrontendController{
		Config:   config,
		Database: database,
		Api:      api,
	}
}

func (controller *FrontendController) Get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"version": static.Version,
		"auth": gin.H{
			"validate_url": "/api/auth/session",
		},
		"twitch": gin.H{
			"auth_url": generation.GenerateTwitchAuthUrl(controller.Config.TwitchPlatform.ApplicationId, controller.Config.TwitchPlatform.RedirectUrl, []string{"user:read:email"}, ""),
		},
	})
}
