package controller

import (
	"errors"
	"webapp/logic"
	"webapp/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUp(c *gin.Context) {
	var t models.SignUpUser
	if err := c.ShouldBind(&t); err != nil {
		zap.L().Error("can not bind for sign up", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := logic.SignUp(&t); err != nil {
		zap.L().Error("fail to sign up", zap.Error(err))
		if errors.Is(err, logic.UserExist) {
			ResponseError(c, CodeUserExist)
		} else {
			ResponseError(c, CodeBusyServer)
		}
		return
	}
	ResponseSuccess(c, nil)
}
func LogIn(c *gin.Context) {
	var t models.LogInUser
	if err := c.ShouldBind(&t); err != nil {
		zap.L().Error("can not bind for log in", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//logic层接手
	if err := logic.LogIn(&t); err != nil {
		zap.L().Error("fail to log in", zap.Error(err))
		if errors.Is(err, logic.UserNoExist) {
			ResponseError(c, CodeUserNoExist)
		} else if errors.Is(err, logic.InvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
		} else {
			ResponseError(c, CodeBusyServer)
		}
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}
