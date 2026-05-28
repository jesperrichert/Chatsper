package twitch_auth_routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pterm/pterm"
	"gorm.io/gorm"
	models "zip.jespersen.chatsper/internal/models/database"
)

type TwitchAuthController TwitchAPIRouteConfig

func (config *TwitchAuthController) authCallback(ctx *gin.Context) {
	code := ctx.Query("code")
	state := ctx.Query("state")
	token, err := config.Client.Token(code, config.Client.ChatsperConfig.TwitchPlatform.RedirectUrl)
	if err != nil {
		return
	}
	user, err := config.Client.User(token.AccessToken)
	if err != nil {
		return
	}

	userModel := models.UserEntity{
		Username:          user.Data[0].Login,
		Email:             user.Data[0].Email,
		Avatar:            user.Data[0].ProfileImageUrl,
		PrimaryConnection: "TWITCH",
		Permissions:       make([]string, 0),
		TwitchUser: &models.TwitchUserEntity{
			Username:     user.Data[0].Login,
			AccessToken:  token.AccessToken,
			Scopes:       token.Scope,
			TwitchUserId: user.Data[0].Id,
			RefreshToken: token.RefreshToken,
		},
	}

	ctxB := context.Background()
	twitchUser, err := gorm.G[models.TwitchUserEntity](config.Database).Where("twitch_user_id = ?", user.Data[0].Id).Preload("User", nil).First(ctxB)
	if twitchUser.TwitchUserId != user.Data[0].Id {
		config.Database.Create(&userModel)
		_, err = gorm.G[models.TwitchUserEntity](config.Database).Where("twitch_user_id = ?", userModel.TwitchUser.TwitchUserId).Update(ctxB, "user_id", userModel.ID)
	} else {
		config.Database.Model(&twitchUser).Updates(
			models.UserEntity{
				Username: user.Data[0].Login,
				Email:    user.Data[0].Email,
				Avatar:   user.Data[0].ProfileImageUrl,
				TwitchUser: &models.TwitchUserEntity{
					Username:     user.Data[0].Login,
					AccessToken:  token.AccessToken,
					Scopes:       token.Scope,
					RefreshToken: token.RefreshToken,
				},
			},
		)
		if err != nil {
			return
		}
	}
	if state == "botAuth" {
		fmt.Println()
		fmt.Println(pterm.Yellow("Bot Authentication Successful... You have authenticated a Bot with Chataper."))
		fmt.Println(pterm.Yellow("Is this your Bot Account. Then set the UserId in the config.yml to verify what account acts as bot account."))
		fmt.Println(pterm.Yellow("Username: " + user.Data[0].Login + ", Twitch User Id: " + user.Data[0].Id))
		fmt.Println()
	}

	ctx.Redirect(http.StatusPermanentRedirect, "/")
}
