package types

type BaseJsonRequest struct {
	DeviceName string `json:"device_name"`
}

type BaseJsonResponse struct {
	Code        int    `json:"code"`
	MessageType string `json:"message_type"`
}

type UserSignupJsonRequest struct {
	BaseJsonRequest
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserSignupJsonResponse struct {
	BaseJsonResponse
	UserId     int64  `json:"user_id"`
	LoginToken string `json:"login_token"`
}

type UserLoginJsonRequest struct {
	BaseJsonRequest
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserLoginJsonResponse UserSignupJsonResponse

type UserLogoutJsonRequest struct {
	BaseJsonRequest
	UserId     int64  `json:"user_id"`
	LoginToken string `json:"login_token"`
}

type UserLogoutJsonResponse BaseJsonResponse

type UserLoadGameJsonRequest struct {
	BaseJsonRequest
	UserId     int64  `json:"user_id"`
	GameId     int64  `json:"game_id"`
	LoginToken string `json:"login_token"`
}

type UserLoadGameJsonResponse struct {
	BaseJsonResponse
	UserId     int64  `json:"user_id"`
	GameId     int64  `json:"game_id"`
	GameStatus int64  `json:"game_status"`
	GameToken  string `json:"game_token"`
}

type UserLeaveGameJsonRequest struct {
	BaseJsonRequest
	UserId    int64  `json:"user_id"`
	GameId    int64  `json:"game_id"`
	GameToken string `json:"game_token"`
}

type UserLeaveGameJsonResponse BaseJsonResponse
