package handler

import (
	"github.com/gin-gonic/gin"

	"wangqingang/domain/common"
	"wangqingang/domain/middleware"
)

func ServerEngine() *gin.Engine {
	if common.Config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	if router == nil {
		panic("create server failed")
	}

	router.Use(middleware.CrossMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	regStaticRouter(router)
	regRegRouter(router)

	return router
}

func regStaticRouter(router *gin.Engine) {
	group := router.Group("/")
	group.Static("/static", "./static")
}

func regRegRouter(router *gin.Engine) {
	group := router.Group("/api/reg")
	group.POST("/start", RegStartHandler)
	group.POST("/stop", RegStopHandler)
	group.GET("/", RegGetHandler)
}
