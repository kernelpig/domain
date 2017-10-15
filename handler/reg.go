package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	e "wangqingang/domain/error"
	"wangqingang/domain/model"
)

func RegStartHandler(c *gin.Context) {
	var req RegStartRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, e.IP(e.IRegStart, e.MParaErr, e.ParaBindErr, err))
		return
	}
	model.RegList.Start(req.AppKeyId, req.AppKeySecret, req.TemplateId, req.DomainList)
	c.JSON(http.StatusOK, RegResponse{Code: e.OK, Status: model.RegList.Status})
}

func RegStopHandler(c *gin.Context) {
	model.RegList.Stop()
	c.JSON(http.StatusOK, RegResponse{Code: e.OK, Status: model.RegList.Status})
}

func RegGetHandler(c *gin.Context) {
	resItem := &RegGetItem{
		AppKeyId:     model.RegList.AppIdKey,
		AppKeySecret: model.RegList.AppSecret,
		TemplateId:   model.RegList.TemplateID,
		DomainList:   model.RegList.Join(),
	}
	c.JSON(http.StatusOK, RegGetResponse{Code: e.OK, List: []*RegGetItem{resItem}})
}
