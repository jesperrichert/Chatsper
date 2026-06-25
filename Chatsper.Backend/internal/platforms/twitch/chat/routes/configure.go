package twitch_chat_routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Configure(api *gin.Engine, database *gorm.DB) {
	platform := api.Group("/platform")
	{
		twitch := platform.Group("/twitch")
		{
			chat := twitch.Group("/chat")
			{
				chat.POST("/messages", messages(database))
			}
		}
	}
}
