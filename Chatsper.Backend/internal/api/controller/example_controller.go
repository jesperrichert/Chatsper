package controller

import (
	"github.com/gin-gonic/gin"
	"zip.jespersen.chatsper/Chatsper.Backend/shared/api/controller"
)

func NewExampleController(config shared.ExampleController) *shared.ExampleController {
	return &shared.ExampleController{Chatsper: config.Chatsper}
}

func (controller *shared.ExampleController) Get(ctx *gin.Context) {

}
