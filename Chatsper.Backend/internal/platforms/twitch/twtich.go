package twitch

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zip.jespersen.chatsper/internal/config/file"
	twitch_auth "zip.jespersen.chatsper/internal/platforms/twitch/auth"
	twitch_chat "zip.jespersen.chatsper/internal/platforms/twitch/chat"
	twitch_client "zip.jespersen.chatsper/internal/platforms/twitch/client"
	Log "zip.jespersen.chatsper/internal/utils/log"
)

type Twitch struct {
	Config   *file.Config
	API      *gin.Engine
	Database *gorm.DB
}

func NewTwitch(
	config *file.Config,
	api *gin.Engine,
	database *gorm.DB,
) *Twitch {
	Log.Info("Loading Twitch Platform Config")
	return &Twitch{
		Config:   config,
		API:      api,
		Database: database,
	}
}

func (t *Twitch) Init() error {

	client := twitch_client.NewTwitchClient(t.Config)
	auth := twitch_auth.NewTwitchAuth(t.Config, t.API, t.Database, client)
	chat := twitch_chat.NewTwitchChat(t.Config, t.API, t.Database, client)

	auth.Init()
	chat.Init()

	return nil
}
