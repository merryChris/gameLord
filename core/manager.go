package core

import (
	"github.com/astaxie/beego/orm"
	"github.com/merryChris/gameLord/types"
	"github.com/merryChris/gameLord/utils"
	"github.com/spf13/viper"
	redis "gopkg.in/redis.v5"
)

type Manager struct {
	Initialized bool
	MysqlOrm    orm.Ormer
	RedisClient *redis.Client
}

func init() {
	orm.RegisterModel(new(types.User), new(types.Game), new(types.GameStatus))
}

func NewManager(redisConf *viper.Viper) (*Manager, error) {
	c, err := utils.NewRedisClient(redisConf)
	if err != nil {
		return nil, err
	}
	return &Manager{
		Initialized: false,
		MysqlOrm:    orm.NewOrm(),
		RedisClient: c,
	}, nil
}
