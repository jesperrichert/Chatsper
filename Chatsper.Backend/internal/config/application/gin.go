package application

import (
	"io"

	"github.com/gin-gonic/gin"
)

func NewGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	engine := gin.Default()
	return engine
}
