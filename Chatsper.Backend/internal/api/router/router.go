package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"zip.jespersen.chatsper/internal/api/controller"
)

type RouterConfig struct {
	Api             *gin.Engine
	IndexController *controller.IndexController
}

func (c *RouterConfig) Setup() {
	if c.Api == nil {
		c.Api = gin.Default()
	}

	// Dashboard
	c.Api.Static("/assets", os.Getenv("FRONTEND_BUILD")+"/client/assets")
	c.Api.StaticFile("/", os.Getenv("FRONTEND_BUILD")+"/client/index.html")
	c.Api.NoRoute(func(c *gin.Context) {
		c.File(os.Getenv("FRONTEND_BUILD") + "/client/index.html")
	})

	// API
	api := c.Api.Group("/api")
	{
		api.GET("/", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"status":  "ok",
				"version": "/api/v1",
				"platforms": gin.H{
					"twtich": "/platform/twtich",
				},
			})
		})
		// v1
		v1 := api.Group("/v1")
		{
			v1.GET("/", c.IndexController.Get)
		}
	}
}
