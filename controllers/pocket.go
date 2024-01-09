package controllers

import (
	"net/http"
	"pocket-serv/models"
	"pocket-serv/service"

	"github.com/gin-gonic/gin"
)

func ShowPocket(ctx *gin.Context) {
	//P1
}

func PublishVIPCard(ctx *gin.Context) {
	var (
		reqVIPCard *models.VIPCardReq
	)

	ctx.Bind(&reqVIPCard)

	name := reqVIPCard.Name
	descripe := reqVIPCard.Descripe
	address := reqVIPCard.Address

	result, err := service.DeployERC721Contract(name, address, descripe)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    400,
			"result":  result,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"result":  result,
		"message": "success",
	})
}

func ShowVIPMembers(ctx *gin.Context) {
	//P1
}

func MintVIPCard(ctx *gin.Context) {
}

func SetVIPCount(ctx *gin.Context) {
	//P1
}
