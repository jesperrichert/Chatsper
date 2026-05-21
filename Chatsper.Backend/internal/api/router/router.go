package router

import (
	"github.com/gin-gonic/gin"
	shared2 "zip.jespersen.chatsper/Chatsper.Backend/shared/api/controller"
	shared "zip.jespersen.chatsper/Chatsper.Backend/shared/application"
)

type RouterConfig struct {
	shared.Chatsper
	ExampleController *shared2.ExampleController
}

func (c *RouterConfig) Setup() {
	if c.Api == nil {
		c.Api = gin.Default()
	}

	api := c.Api.Group("/api")
	{
		// v1
		v1 := api.Group("/v1")
		{
			v1.GET("/", c.ExampleController.Get)
		}
	}
}
