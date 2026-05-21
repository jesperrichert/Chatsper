package application

import "github.com/gin-gonic/gin"

func NewGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	return engine
}
