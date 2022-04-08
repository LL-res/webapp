package routers

import (
	"time"
	"webapp/controller"
	"webapp/middleware"

	gzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init() (r *gin.Engine) {
	r = gin.New()
	r.Use(gzap.Ginzap(zap.L(), time.RFC850, true))
	r.Use(gzap.RecoveryWithZap(zap.L(), true))
	r.POST("/signup", controller.SignUp)
	r.POST("login", controller.LogIn)
	c := r.Group("/community", middleware.JWTAuthorization())
	{
		c.GET("/", controller.Community)
		c.GET("/:id", controller.CommunityDetail)
	}
	return
}
