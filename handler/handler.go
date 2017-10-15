package handler

import (
	"github.com/gin-gonic/gin"

	"wangqingang/cunxun/common"
	"wangqingang/cunxun/middleware"
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

	regRegRouter(router)

	return router
}

func regRegRouter(router *gin.Engine) {
	group := router.Group("/api/reg")
	group.GET("/start", RegStartHandler)
	group.GET("/stop", RegStopHandler)
	group.GET("/kill", RegKillHandler)
}
