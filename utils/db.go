package utils

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	redis "gopkg.in/redis.v5"
)

func InitOrm(conf *viper.Viper) {
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",
		conf.GetString("username"),
		conf.GetString("password"),
		conf.GetString("host"),
		conf.GetString("port"),
		conf.GetString("dbname")))
	orm.SetMaxIdleConns("default", conf.GetInt("max_idle_connections"))
	orm.SetMaxOpenConns("default", conf.GetInt("max_open_connections"))
}

func NewRedisClient(conf *viper.Viper) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.GetString("host"), conf.GetString("port")),
		Password: conf.GetString("password"),
		DB:       conf.GetInt("db"),
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}
	return client, nil
}
