package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"webapp/models"
)

func QueryExistenceByName(name string) (bool, error) {
	sqlStr := "select count(username) from user where username = ?"
	var count int
	if err := Db.Get(&count, sqlStr, name); err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
func InsertUser(user *models.User) error {
	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	fmt.Println("ijij")
	if _, err := Db.Exec(sqlStr, user.ID, user.Name, Encrypt(user.PassWord)); err != nil {
		return err
	}
	return nil
}
func Encrypt(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
func QueryPasswordByName(name string) (string, error) {
	sqlStr := "select password from user where username = ?"
	var ans string
	if err := Db.Get(&ans, sqlStr, name); err != nil {
		return "", err
	}
	return ans, nil
}
