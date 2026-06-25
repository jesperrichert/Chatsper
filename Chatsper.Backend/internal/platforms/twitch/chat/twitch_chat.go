package twitch_chat

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"resty.dev/v3"
	"zip.jespersen.chatsper/internal/config/file"
	models "zip.jespersen.chatsper/internal/models/http/twitch"
	"zip.jespersen.chatsper/internal/platforms/twitch/chat/routes"
	twitch_client "zip.jespersen.chatsper/internal/platforms/twitch/client"
	"zip.jespersen.chatsper/shared/static"
)

type TwitchChat struct {
	Config   *file.Config
	API      *gin.Engine
	Database *gorm.DB
	Client   *twitch_client.TwitchClient
}

func NewTwitchChat(
	config *file.Config,
	api *gin.Engine,
	database *gorm.DB,
	client *twitch_client.TwitchClient,
) *TwitchChat {
	return &TwitchChat{
		Config:   config,
		API:      api,
		Database: database,
		Client:   client,
	}
}

func (chat *TwitchChat) Init() {
	twitch_chat_routes.Configure(chat.API, chat.Database)
}

func (chat *TwitchChat) SendMessage(broadcasterId int, senderId int, message string, bearer string) *error {
	client := resty.New()
	req := client.R()

	req.Header.Add("Authorization", "Bearer "+bearer)
	req.Header.Add("Client-Id", chat.Config.TwitchPlatform.ApplicationId)
	req.Header.Add("Content-Type", "application/json")

	req.SetBody(
		&models.TwitchChatMessageRequest{
			BroadcasterId: broadcasterId,
			SenderId:      senderId,
			Message:       message,
		},
	)
	_, err := req.Post(static.TwitchBaseAPIUrl + "/chat/messages")
	return &err
}
