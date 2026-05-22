package twitch

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zip.jespersen.chatsper/internal/config/file"
	Log "zip.jespersen.chatsper/internal/utils"
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
	return nil
}
