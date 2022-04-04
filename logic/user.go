package logic

import (
	"errors"
	"fmt"
	"webapp/dao/mysql"
	"webapp/models"
	"webapp/pkg/snowflake"

	"go.uber.org/zap"
)

func SignUp(t *models.SignUpUser) error {
	//查询是否存在
	exist, err := mysql.QueryExistenceByName(t.Name)
	if err != nil {
		return err
	}
	if exist == true {
		return errors.New("the user has already existed")
	} else {
		id := snowflake.GenID()
		//插入数据库
		user := models.User{
			Name:     t.Name,
			PassWord: t.Password,
			ID:       id,
		}
		if err := mysql.InsertUser(&user); err != nil {
			zap.L().Error("fail to insert", zap.Error(err))
			return err
		}
		fmt.Println("llll")
	}
	return nil
	//不存在则给一个id
	//存在则返回错误
}
