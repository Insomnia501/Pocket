package routes

import (
	"pocket-serv/middleware"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// router.Use(cors.Default())
	router.Use(middleware.LogHandler())
	router.Use(middleware.ErrHandler())
	pprof.Register(router)
	loadRouter(router)
	return router
}

func loadRouter(router *gin.Engine) {
	InitHolderRouter(router)
	InitPublisherRouter(router)
}
