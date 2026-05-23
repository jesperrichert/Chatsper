package router

import (
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
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

	isDockerBuild := os.Getenv("IS_DOCKER_BUILD")

	if isDockerBuild == "YES" {
		c.Api.Static("/assets", os.Getenv("FRONTEND_BUILD")+"/client/assets")
		c.Api.StaticFile("/", os.Getenv("FRONTEND_BUILD")+"/client/index.html")
		c.Api.NoRoute(func(c *gin.Context) {
			c.File(os.Getenv("FRONTEND_BUILD>") + "/client/index.html")
		})
	} else {
		frontendBox, _ := rice.FindBox("../../../../Chatsper.Frontend/build/client")
		assetsBox, _ := rice.FindBox("../../../../Chatsper.Frontend/build/client/assets")
		assets := assetsBox.HTTPBox()
		indexPage, _ := frontendBox.String("index.html")

		// Dashboard
		c.Api.StaticFS("/assets", assets)
		c.Api.GET("/", func(ctx *gin.Context) {
			ctx.Writer.WriteHeader(http.StatusOK)
			ctx.Writer.Write([]byte(indexPage))
		})
		c.Api.NoRoute(func(c *gin.Context) {
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Write([]byte(indexPage))
		})

	}

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
