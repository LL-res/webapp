package logic

import (
	"errors"
	"webapp/dao/mysql"
	"webapp/models"
	"webapp/pkg/snowflake"

	"go.uber.org/zap"
)

var (
	UserNoExist     error = errors.New("No such user")
	InvalidPassword error = errors.New("invalid user or password")
	UserExist       error = errors.New("the user has already been existed")
)

func SignUp(t *models.SignUpUser) error {
	//查询是否存在
	exist, err := mysql.QueryExistenceByName(t.Name)
	if err != nil {
		return err
	}
	if exist == true {
		return UserExist
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
	}
	return nil
	//不存在则给一个id
	//存在则返回错误
}
func LogIn(t *models.LogInUser) error {
	exist, err := mysql.QueryExistenceByName(t.Name)
	if err != nil {
		return err
	} else if exist == false {
		return UserNoExist
	} else {
		result, err := mysql.QueryPasswordByName(t.Name)
		if err != nil {
			return err
		} //向dao层要一下加密密码，看看是否与本次的输入一致
		if mysql.Encrypt(t.Password) == result {
			return nil
		} else {
			return InvalidPassword
		}
	}
}
