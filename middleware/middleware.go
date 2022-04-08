package middleware

import (
	"strings"
	"webapp/controller"
	"webapp/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func JWTAuthorization() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeAuthFail)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeAuthFail)
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeAuthFail)
			c.Abort()
			return
		}
		c.Set(controller.CurrentName, mc.Username)
		c.Next()
	}
}
