package types

import (
	"time"

	"github.com/merryChris/gameLord/utils"
	redis "gopkg.in/redis.v5"
)

type User struct {
	Id            int64  `orm:"column(id)"`
	Name          string `orm:"column(name);unique"`
	Phone         string `orm:"column(phone)"`
	Salt          string `orm:"column(salt)"`
	Password      string `orm:"column(password)"`
	CurrentGameId int64  `orm:"-"`
	CurrentDevice string `orm:"-"`
}

type Game struct {
	Id          int64  `orm:"column(id)"`
	Name        string `orm:"column(name)"`
	Description string `orm:"column(description)"`
}

type GameStatus struct {
	Id     int64 `orm:"column(id)"`
	UserId int64 `orm:"column(user_id)"`
	GameId int64 `orm:"column(game_id)"`
	Status int64 `orm:"column(status)"` // Game Archive Exists or Not
}

func (this *User) SetLoginToken(client *redis.Client, tokenExpiration time.Duration) (string, bool) {
	key := utils.GetSha1Hash(utils.GenerateLoginTokenKey(this.Salt, this.Id, this.CurrentDevice))
	value := utils.GetSha1Hash(utils.GenerateTokenValue(this.Salt))
	err := client.Set(key, value, tokenExpiration).Err()
	return value, err == nil
}

func (this *User) CheckLoginToken(client *redis.Client, token string) bool {
	key := utils.GetSha1Hash(utils.GenerateLoginTokenKey(this.Salt, this.Id, this.CurrentDevice))
	return client.Get(key).Val() == token
}

func (this *User) DelLoginToken(client *redis.Client) {
	key := utils.GetSha1Hash(utils.GenerateLoginTokenKey(this.Salt, this.Id, this.CurrentDevice))
	client.Del(key)
}

func (this *User) SetGameToken(client *redis.Client, tokenExpiration time.Duration) (string, bool) {
	key := utils.GetSha1Hash(utils.GenerateGameTokenKey(this.Salt, this.Id, this.CurrentGameId, this.CurrentDevice))
	value := utils.GetSha1Hash(utils.GenerateTokenValue(this.Salt))
	err := client.Set(key, value, tokenExpiration).Err()
	return value, err == nil
}

func (this *User) CheckGameToken(client *redis.Client, token string) bool {
	key := utils.GetSha1Hash(utils.GenerateGameTokenKey(this.Salt, this.Id, this.CurrentGameId, this.CurrentDevice))
	return client.Get(key).Val() == token
}

func (this *User) DelGameToken(client *redis.Client) {
	key := utils.GetSha1Hash(utils.GenerateGameTokenKey(this.Salt, this.Id, this.CurrentGameId, this.CurrentDevice))
	client.Del(key)
}
