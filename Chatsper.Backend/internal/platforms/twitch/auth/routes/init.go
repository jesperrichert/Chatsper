package twitch_auth_routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	twitch_client "zip.jespersen.chatsper/internal/platforms/twitch/client"
)

type TwitchAPIRouteConfig struct {
	Client   *twitch_client.TwitchClient
	Database *gorm.DB
}

type Routes struct {
	Callback *TwitchAuthController
}

func Configure(api *gin.Engine, database *gorm.DB, client *twitch_client.TwitchClient) {

	config := &TwitchAPIRouteConfig{
		Client:   client,
		Database: database,
	}

	routes := &Routes{
		Callback: (*TwitchAuthController)(config),
	}

	platform := api.Group("/platform")
	{
		twitch := platform.Group("/twitch")
		{
			auth := twitch.Group("/auth")
			{
				auth.GET("/callback", routes.Callback.authCallback)
			}
		}
	}
}
