package generation

import "strings"

func GenerateTwitchAuthUrl(clientId string, redirect string, scopes []string) string {
	return string("https://id.twitch.tv/oauth2/authorize?response_type=code&client_id=" + clientId + "&redirect_uri=" + redirect + "&scope=" + strings.Join(scopes, "+"))
}
