package controller

import (
	"errors"
	"webapp/logic"
	"webapp/models"
	"webapp/pkg/jwt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const CurrentName = "username"

var ErrorUserNotLogin = errors.New("user not log in")

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
	token, _ := jwt.GenToken(t.Name) //这里最好的处理是把这个东西扔到logic层，让logic层的Login函数返回一个string与error
	ResponseSuccess(c, token)
}
func GetCurrentUser(c *gin.Context) (string, error) {
	username, ok := c.Get(CurrentName)
	if !ok {
		ResponseErrorWithMsg(c, CodeAuthFail, "please log in")
		return "", ErrorUserNotLogin
	}
	return username.(string), nil
}
