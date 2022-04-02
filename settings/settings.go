package settings

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf *App

type App struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Port int    `mapstructure:"port"`
	*Mysql
	*Log
	*Redis
}
type Mysql struct {
	Host     string `mapstructure:"host"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
	Port     int    `mapstructure: "port"`
}
type Log struct {
	File      string `mapstructure:"filename"`
	Level     string `mapstructure:"level"`
	MaxBackUp int    `mapstructure:"max_backups"`
	MaxAge    int    `mapstructure:"max_age"`
	MaxSize   int    `mapstructure:"max_size"`
}
type Redis struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Db   int    `mapstructure:"db"`
}

func Init() (err error) {
	Conf = new(App)
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		log.Println("1")
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("hai")
		fmt.Println("changed")
		if err = viper.Unmarshal(Conf); err != nil {
			log.Println("2")
			return
		}
	})
	if err = viper.Unmarshal(Conf); err != nil {
		log.Println("2")
		return
	}

	return

}
