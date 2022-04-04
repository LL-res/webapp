package main

import (
	"fmt"
	"webapp/dao/mysql"
	"webapp/logger"
	"webapp/pkg/snowflake"
	"webapp/routers"
	"webapp/settings"

	"go.uber.org/zap"
)

func main() {
	//加载配置
	err := settings.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//初始化日志
	if err := logger.Init(settings.Conf); err != nil {
		fmt.Println(err.Error())
		return
	}
	defer zap.L().Sync()
	//初始化mysql
	if err := mysql.Init(settings.Conf); err != nil {
		zap.L().Error("mysql error", zap.Error(err))
		return
	}
	defer mysql.Db.Close()
	snowflake.Init("2022-04-04", 1)
	//初始化redis
	//注册路由
	r := routers.Init()
	//启动服务
	r.Run()

}
