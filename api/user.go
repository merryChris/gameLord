package api

import (
	"io"
	"net/http"

	"github.com/merryChris/gameLord/core"
	"github.com/merryChris/gameLord/types"
	"github.com/spf13/viper"
)

type UserHandler struct {
	Handler
	userManager *core.UserManager
}

var (
	RoutingUserHandler = &UserHandler{}
)

func (this *UserHandler) Init(redisConf *viper.Viper) error {
	um, err := core.NewUserManager(redisConf)
	if err != nil {
		return err
	}
	this.userManager = um
	this.Initialized = true
	return nil
}

func (this *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	if resp, ok := this.SelfCheck(); !ok {
		io.WriteString(w, resp)
		return
	}

	var requestObject types.UserSignupJsonRequest
	if resp, ok := ParseRequestJsonData(r.Body, &requestObject); !ok {
		io.WriteString(w, resp)
		return
	}
	if resp, ok := CheckRequestJsonData(requestObject); !ok {
		io.WriteString(w, resp)
		return
	}
	resp, _ := this.userManager.Signup(requestObject)
	io.WriteString(w, resp)
	return
}

func (this *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if resp, ok := this.SelfCheck(); !ok {
		io.WriteString(w, resp)
		return
	}

	var requestObject types.UserLoginJsonRequest
	if resp, ok := ParseRequestJsonData(r.Body, &requestObject); !ok {
		io.WriteString(w, resp)
		return
	}
	if resp, ok := CheckRequestJsonData(requestObject); !ok {
		io.WriteString(w, resp)
		return
	}
	resp, _ := this.userManager.Login(requestObject)
	io.WriteString(w, resp)
}

func (this *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if resp, ok := this.SelfCheck(); !ok {
		io.WriteString(w, resp)
		return
	}

	var requestObject types.UserLogoutJsonRequest
	if resp, ok := ParseRequestJsonData(r.Body, &requestObject); !ok {
		io.WriteString(w, resp)
		return
	}
	if resp, ok := CheckRequestJsonData(requestObject); !ok {
		io.WriteString(w, resp)
		return
	}
	resp, _ := this.userManager.Logout(requestObject)
	io.WriteString(w, resp)
}

func (this *UserHandler) LoadGame(w http.ResponseWriter, r *http.Request) {
	if resp, ok := this.SelfCheck(); !ok {
		io.WriteString(w, resp)
		return
	}

	var requestObject types.UserLoadGameJsonRequest
	if resp, ok := ParseRequestJsonData(r.Body, &requestObject); !ok {
		io.WriteString(w, resp)
		return
	}
	if resp, ok := CheckRequestJsonData(requestObject); !ok {
		io.WriteString(w, resp)
		return
	}
	resp, _ := this.userManager.LoadGame(requestObject)
	io.WriteString(w, resp)
}

func (this *UserHandler) LeaveGame(w http.ResponseWriter, r *http.Request) {
	if resp, ok := this.SelfCheck(); !ok {
		io.WriteString(w, resp)
		return
	}

	var requestObject types.UserLeaveGameJsonRequest
	if resp, ok := ParseRequestJsonData(r.Body, &requestObject); !ok {
		io.WriteString(w, resp)
		return
	}
	if resp, ok := CheckRequestJsonData(requestObject); !ok {
		io.WriteString(w, resp)
		return
	}
	resp, _ := this.userManager.LeaveGame(requestObject)
	io.WriteString(w, resp)
}

func (this *UserHandler) Close() {
	if this.Initialized {
		this.userManager.Close()
		this.Initialized = false
	}
}
