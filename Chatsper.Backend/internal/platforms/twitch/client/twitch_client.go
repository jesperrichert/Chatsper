package twitch_client

import (
	"resty.dev/v3"
	"zip.jespersen.chatsper/internal/config/file"
	models "zip.jespersen.chatsper/internal/models/http/twitch"
	"zip.jespersen.chatsper/shared/static"
)

type TwitchClient struct {
	ChatsperConfig *file.Config
	TwitchAPI      resty.Client
}

func NewTwitchClient(
	config *file.Config,
) *TwitchClient {
	return &TwitchClient{
		ChatsperConfig: config,
		TwitchAPI:      *resty.New(),
	}
}

func (client *TwitchClient) User(bearer string, username string) (*models.TwtichAPIUserDataBaseResponse, error) {
	api := client.TwitchAPI.R()
	api.SetResult(&models.TwtichAPIUserDataBaseResponse{})

	api.Header.Add("Authorization", "Bearer "+bearer)
	api.Header.Add("Client-Id", client.ChatsperConfig.TwitchPlatform.ApplicationId)
	api.Header.Add("Content-Type", "application/json")

	req, err := api.Get(static.TwitchBaseAPIUrl + "/users?login=" + username)
	if err != nil {
		return nil, err
	}

	return req.Result().(*models.TwtichAPIUserDataBaseResponse), nil
}

func (client *TwitchClient) Validate(accessToken string) (int, error) {
	api := client.TwitchAPI.R()
	api.SetResult(&models.TwtichAPIUserAuthResponse{})

	api.Header.Add("Authorization", "OAuth "+accessToken)
	req, err := api.Get("https://id.twitch.tv/oauth2/validate")
	if err != nil {
		return 500, err
	}

	return req.StatusCode(), nil
}

func (client *TwitchClient) Token(code string, redirect string) (*models.TwtichAPIUserAuthResponse, error) {
	api := client.TwitchAPI.R()
	api.SetResult(&models.TwtichAPIUserAuthResponse{})

	req, err := api.Post("https://id.twitch.tv/oauth2/token?client_id=" + client.ChatsperConfig.TwitchPlatform.ApplicationId + "&client_secret=" + client.ChatsperConfig.TwitchPlatform.ApplicationSecret + "&code=" + code + "&grant_type=authorization_code&redirect_uri=" + redirect)
	if err != nil {
		return nil, err
	}

	return req.Result().(*models.TwtichAPIUserAuthResponse), nil
}

func (client *TwitchClient) Refresh(refreshToken string, redirect string) (*models.TwtichAPIUserAuthResponse, error) {
	api := client.TwitchAPI.R()
	api.SetResult(&models.TwtichAPIUserAuthResponse{})

	req, err := api.Post("https://id.twitch.tv/oauth2/token?client_id=" + client.ChatsperConfig.TwitchPlatform.ApplicationId + "&client_secret=" + client.ChatsperConfig.TwitchPlatform.ApplicationSecret + "&refresh_token=" + refreshToken + "&grant_type=refresh_token&redirect_uri=" + redirect)
	if err != nil {
		return nil, err
	}

	return req.Result().(*models.TwtichAPIUserAuthResponse), nil
}
