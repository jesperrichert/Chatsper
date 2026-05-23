package models

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Username    string
	Email       string
	CreatedAt   int
	Permissions []string
}

type TwitchUserEntity struct {
	gorm.Model
	Username     string
	UserId       string
	AccessToken  string
	RefreshToken string
	Scopes       []string
	User         UserEntity
}
