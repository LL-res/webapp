package controller

import (
	"strconv"
	"webapp/logic"

	"github.com/gin-gonic/gin"
)

func Community(c *gin.Context) {
	//logic层直接接手，获取社区id与名称
	s, err := logic.GetCommunityList()
	if err != nil {
		ResponseError(c, CodeBusyServer)
	} else {
		ResponseSuccess(c, s)
	}
}
func CommunityDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	} else {
		detail, err := logic.GetCommunityDetail(id)
		if err != nil {
			ResponseError(c, CodeBusyServer)
		} else {
			ResponseSuccess(c, detail)
		}
	}
}
