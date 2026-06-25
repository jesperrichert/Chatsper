package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserAPI struct {
	gorm.Model
	SessionID uuid.UUID
	UserID    uint
	User      *UserEntity
}
