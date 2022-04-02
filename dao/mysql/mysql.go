package mysql

import (
	"fmt"
	"webapp/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func Init(c *settings.App) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", c.Mysql.Name, c.Password, c.Mysql.Host, c.Mysql.Port, c.Dbname)
	Db, err = sqlx.Connect("mysql", dsn)
	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(10)
	return
}
