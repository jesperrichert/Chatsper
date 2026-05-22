package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zip.jespersen.chatsper/internal/config/file"
)

type IndexController struct {
	Config   *file.Config
	Database *gorm.DB
	Api      *gin.Engine
}

func NewExampleController(
	config *file.Config,
	database *gorm.DB,
	api *gin.Engine,
) *IndexController {
	return &IndexController{
		Config:   config,
		Database: database,
		Api:      api,
	}
}

func (controller *IndexController) Get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}
