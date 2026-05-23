package twitch_chat_routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func messages(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{})
	}
}
