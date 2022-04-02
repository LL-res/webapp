package routers

import (
	"time"

	gzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init() (r *gin.Engine) {
	r = gin.New()
	r.Use(gzap.Ginzap(zap.L(), time.RFC850, true))
	r.Use(gzap.RecoveryWithZap(zap.L(), true))
	//r.GET()
	return
}
