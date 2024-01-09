package routes

import (
	controller "pocket-serv/controllers"

	"github.com/gin-gonic/gin"
)

func InitHolderRouter(router *gin.Engine) {
	holder := router.Group("/holder")
	// holder router
	holder.POST("/showpocket", controller.ShowPocket)
}

func InitPublisherRouter(router *gin.Engine) {
	publisher := router.Group("/publisher")
	// publisher router
	publisher.POST("/publishvipcard", controller.PublishVIPCard)
	publisher.POST("/showvipmembers", controller.ShowVIPMembers)
	publisher.POST("/mintvipcard", controller.MintVIPCard)
	publisher.POST("/setvipcount", controller.SetVIPCount)
}
