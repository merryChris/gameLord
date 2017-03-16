package core

import (
	"encoding/json"
	"time"

	"github.com/merryChris/gameLord/types"
	"github.com/merryChris/gameLord/utils"
	"github.com/spf13/viper"
)

type UserManager struct {
	*Manager
	loginTokenExpiration time.Duration
	gameTokenExpiration  time.Duration
}

func NewUserManager(redisConf *viper.Viper) (*UserManager, error) {
	m, err := NewManager(redisConf)
	if err != nil {
		return nil, err
	}

	um := &UserManager{
		Manager:              m,
		loginTokenExpiration: redisConf.GetDuration("login_token_expiration"),
		gameTokenExpiration:  redisConf.GetDuration("game_token_expiration"),
	}
	um.Initialized = true
	return um, nil
}

func (this *UserManager) Signup(req types.UserSignupJsonRequest) (string, bool) {
	user := types.User{Name: utils.EncodeString(req.Name)}
	if err := this.MysqlOrm.Read(&user, "Name"); err == nil {
		return types.Error111(), false
	}

	user.Salt = utils.GenerateUserSalt(user.Name)
	user.Password = utils.GenerateUserPassword(user.Salt, req.Password)
	userId, err := this.MysqlOrm.Insert(&user)
	if err != nil {
		return types.Error101(), false
	}
	user.CurrentDevice = req.DeviceName

	token, ok := user.SetLoginToken(this.RedisClient, this.loginTokenExpiration)
	if !ok {
		return types.Error101(), false
	}

	respBytes, _ := json.Marshal(&types.UserSignupJsonResponse{
		BaseJsonResponse: types.BaseJsonResponse{
			Code:        110,
			MessageType: "user_success"},
		UserId:     userId,
		LoginToken: token})
	return string(respBytes), true
}

func (this *UserManager) Login(req types.UserLoginJsonRequest) (string, bool) {
	user := types.User{Name: utils.EncodeString(req.Name)}
	if err := this.MysqlOrm.Read(&user, "Name"); err != nil {
		return types.Error112(), false
	}

	password := utils.GenerateUserPassword(user.Salt, req.Password)
	if password != user.Password {
		return types.Error113(), false
	}
	user.CurrentDevice = req.DeviceName

	token, ok := user.SetLoginToken(this.RedisClient, this.loginTokenExpiration)
	if !ok {
		return types.Error101(), false
	}

	respBytes, _ := json.Marshal(&types.UserLoginJsonResponse{
		BaseJsonResponse: types.BaseJsonResponse{
			Code:        110,
			MessageType: "user_success"},
		UserId:     user.Id,
		LoginToken: token})
	return string(respBytes), true
}

func (this *UserManager) Logout(req types.UserLogoutJsonRequest) (string, bool) {
	user := types.User{Id: req.UserId}
	if err := this.MysqlOrm.Read(&user); err != nil {
		return types.Error112(), false
	}
	user.CurrentDevice = req.DeviceName

	if ok := user.CheckLoginToken(this.RedisClient, req.LoginToken); !ok {
		return types.Error114(), false
	}

	user.DelLoginToken(this.RedisClient)
	respBytes, _ := json.Marshal(&types.UserLogoutJsonResponse{
		Code:        110,
		MessageType: "user_success"})
	return string(respBytes), true
}

func (this *UserManager) LoadGame(req types.UserLoadGameJsonRequest) (string, bool) {
	user := types.User{Id: req.UserId}
	if err := this.MysqlOrm.Read(&user); err != nil {
		return types.Error112(), false
	}
	user.CurrentGameId = req.GameId
	user.CurrentDevice = req.DeviceName

	if ok := user.CheckLoginToken(this.RedisClient, req.LoginToken); !ok {
		return types.Error114(), false
	}

	game := types.Game{Id: req.GameId}
	if err := this.MysqlOrm.Read(&game); err != nil {
		return types.Error121(), false
	}

	gameStatus := types.GameStatus{UserId: req.UserId, GameId: req.GameId, Status: 0}
	if _, _, err := this.MysqlOrm.ReadOrCreate(&gameStatus, "UserId", "GameId"); err != nil {
		return types.Error101(), false
	}

	token, ok := user.SetGameToken(this.RedisClient, this.gameTokenExpiration)
	if !ok {
		return types.Error101(), false
	}

	respBytes, _ := json.Marshal(&types.UserLoadGameJsonResponse{
		BaseJsonResponse: types.BaseJsonResponse{
			Code:        110,
			MessageType: "user_success"},
		UserId:     gameStatus.UserId,
		GameId:     gameStatus.GameId,
		GameStatus: gameStatus.Status,
		GameToken:  token})
	return string(respBytes), true
}

func (this *UserManager) LeaveGame(req types.UserLeaveGameJsonRequest) (string, bool) {
	user := types.User{Id: req.UserId}
	if err := this.MysqlOrm.Read(&user); err != nil {
		return types.Error112(), false
	}
	user.CurrentGameId = req.GameId
	user.CurrentDevice = req.DeviceName

	if ok := user.CheckGameToken(this.RedisClient, req.GameToken); !ok {
		return types.Error115(), false
	}

	user.DelGameToken(this.RedisClient)
	respBytes, _ := json.Marshal(&types.UserLeaveGameJsonResponse{
		Code:        110,
		MessageType: "user_success"})
	return string(respBytes), true
}

func (this *UserManager) Close() {
	if this.Initialized {
		this.Initialized = false
	}
}
