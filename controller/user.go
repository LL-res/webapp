package controller

import (
	"net/http"
	"webapp/logic"
	"webapp/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUp(c *gin.Context) {
	var t models.SignUpUser
	if err := c.ShouldBind(&t); err != nil {
		zap.L().Error("can not bind", zap.Error(err))
		c.JSON(200, gin.H{
			"msg": "can not bind",
		})
	}
	if err := logic.SignUp(&t); err != nil {
		zap.L().Error("fail to sign up", zap.Error(err))
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
