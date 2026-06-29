package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zip.jespersen.chatsper/internal/config/file"
	models "zip.jespersen.chatsper/internal/models/database"
)

type AuthController struct {
	Config   *file.Config
	Database *gorm.DB
	Api      *gin.Engine
}

type authRequest struct {
	SessionID string `json:"session_id"`
}

func NewAuthController(
	config *file.Config,
	database *gorm.DB,
	api *gin.Engine,
) *AuthController {
	return &AuthController{
		Config:   config,
		Database: database,
		Api:      api,
	}
}

func (controller *AuthController) Post(ctx *gin.Context) {
	sessionId := ctx.GetHeader("Authorization")
	if sessionId == "" {
		model := authRequest{}
		err := ctx.ShouldBindBodyWithJSON(&model)
		if err != nil {
			// Ignore
		}
		sessionId = model.SessionID
	}

	userAPI, err := gorm.G[models.UserAPI](controller.Database).Where("session_id = ?", sessionId).Preload("User", nil).First(context.Background())
	if err != nil {
		println(err.Error())
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"authenticated": false,
			"error":         "Database Error",
		})
		return
	}

	if userAPI.SessionID.String() == sessionId {
		ctx.JSON(200, gin.H{
			"authenticated": true,
		})
		return
	}

	ctx.JSON(http.StatusUnauthorized, gin.H{
		"authenticated": false,
		"error":         "Invalid Session ID",
	})
}
