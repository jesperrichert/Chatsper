package shared

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zip.jespersen.chatsper/Chatsper.Backend/internal/config/file"
)

type Chatsper struct {
	Config   *file.Config
	Database *gorm.DB
	Api      *gin.Engine
}
