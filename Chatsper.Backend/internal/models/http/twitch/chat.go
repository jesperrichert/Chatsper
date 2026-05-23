package models

type TwitchChatMessageRequest struct {
	BroadcasterId int    `json:"broadcaster_id"`
	SenderId      int    `json:"sender_id"`
	Message       string `json:"message"`
}
