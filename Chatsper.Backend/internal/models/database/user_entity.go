package models

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Username          string
	Email             string
	Avatar            string
	Permissions       []string `gorm:"serializer:json"`
	PrimaryConnection string   // TWITCH, ...
	TwitchUserID      uint
	TwitchUser        *TwitchUserEntity
}

type TwitchUserEntity struct {
	gorm.Model
	Username     string
	TwitchUserId string
	AccessToken  string
	RefreshToken string
	Scopes       []string `gorm:"serializer:json"`
	UserID       uint
	User         *UserEntity
}
