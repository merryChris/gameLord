package main

import (
	"fmt"
	"time"

	redis "gopkg.in/redis.v5"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/merryChris/gameLord/core"
	"github.com/merryChris/gameLord/hangman"
	"github.com/merryChris/gameLord/types"
	"github.com/merryChris/gameLord/utils"
)

func newMysqlOrm() orm.Ormer {
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:yamiedie@/game_lord?charset=utf8&loc=Local")
	orm.RunSyncdb("default", true, true)

	return orm.NewOrm()
}

func newRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "foobared",
		DB:       0,
	})

	return client
}

func main() {
	o := newMysqlOrm()

	u1 := new(types.User)
	u1.Name = utils.EncodeString("Zanwen")
	u1.Phone = "13082819196"
	u1.Salt = utils.GenerateUserSalt(u1.Name)
	u1.Password = utils.GenerateUserPassword(u1.Salt, "123456")
	fmt.Println(o.Insert(u1))
	u1.CurrentDevice = "chrome"

	u2 := new(types.User)
	u2.Name = utils.EncodeString("Gangga")
	u2.Phone = "13082819197"
	u2.Salt = utils.GenerateUserSalt(u2.Name)
	u2.Password = utils.GenerateUserPassword(u2.Salt, "123456")
	fmt.Println(o.Insert(u2))

	game := new(types.Game)
	game.Name = "Hangman"
	game.Description = "Hangman is a guessing game for players. The player tries to guess it by suggesting letters, within a certain number of guesses."
	fmt.Println(o.Insert(game))

	c1 := new(hangman.Category)
	c1.Name = "Animal"
	fmt.Println(o.Insert(c1))

	d1 := new(hangman.Dict)
	d1.Category = c1
	d1.Word = "lion"
	fmt.Println(o.Insert(d1))

	d2 := new(hangman.Dict)
	d2.Category = c1
	d2.Word = "panda"
	fmt.Println(o.Insert(d2))

	d3 := new(hangman.Dict)
	d3.Category = c1
	d3.Word = "tiger"
	fmt.Println(o.Insert(d3))
}
