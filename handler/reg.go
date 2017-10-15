package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	e "wangqingang/cunxun/error"
)

func RegStartHandler(c *gin.Context) {
	var req RegStartRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, e.IP(e.IArticleCreate, e.MParamsErr, e.ParamsBindErr, err))
		return
	}
	c.JSON(http.StatusOK, OKResponse{Code: e.OK})
}
