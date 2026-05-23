package twitch_auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zip.jespersen.chatsper/internal/config/file"
	twitch_auth_routes "zip.jespersen.chatsper/internal/platforms/twitch/auth/routes"
	twitch_client "zip.jespersen.chatsper/internal/platforms/twitch/client"
)

type TwitchAuth struct {
	Config   *file.Config
	API      *gin.Engine
	Database *gorm.DB
	Client   *twitch_client.TwitchClient
}

func NewTwitchAuth(
	config *file.Config,
	api *gin.Engine,
	database *gorm.DB,
	client *twitch_client.TwitchClient,
) *TwitchAuth {
	return &TwitchAuth{
		Config:   config,
		API:      api,
		Database: database,
		Client:   client,
	}
}

func (auth *TwitchAuth) Init() {
	twitch_auth_routes.Configure(auth.API, auth.Database, auth.Client)
}
