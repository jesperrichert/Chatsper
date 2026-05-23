package twitch_auth_routes

import (
	"github.com/gin-gonic/gin"
)

type Callback TwitchAPIRouteConfig

func (config *Callback) authCallback(ctx *gin.Context) {
	code := ctx.Query("code")
	ctx.JSON(200, gin.H{"code": code})
}
