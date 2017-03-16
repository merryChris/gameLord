package types

import (
	"encoding/json"
)

func Error101() string {
	respBytes, _ := json.Marshal(&BaseJsonResponse{
		Code:        101,
		MessageType: "err_internal"})
	return string(respBytes)
}

func Error102() string {
	respBytes, _ := json.Marshal(&BaseJsonResponse{
		Code:        102,
		MessageType: "err_missing_data"})
	return string(respBytes)
}

func Error103() string {
	respBytes, _ := json.Marshal(&BaseJsonResponse{
		Code:        103,
		MessageType: "err_invalid_data"})
	return string(respBytes)
}

func Error111() string {
	respBytes, _ := json.Marshal(&BaseJsonResponse{
		Code:        111,
		MessageType: "err_duplicate_uname"})
	return string(respBytes)
}

func Error112() string {
	respBytes, _ := json.Marshal(&BaseJsonResponse{
		Code:        112,
		MessageType: "err_user_not_exist"})
	return string(respBytes)
}

func Error113() string {
	respBytes, _ := json.Marshal(&BaseJsonResponse{
		Code:        113,
		MessageType: "err_password_mismatch"})
	return string(respBytes)
}

func Error114() string {
	respBytes, _ := json.Marshal(&BaseJsonResponse{
		Code:        114,
		MessageType: "err_login_token_mismatch"})
	return string(respBytes)
}

func Error115() string {
	respBytes, _ := json.Marshal(&BaseJsonResponse{
		Code:        115,
		MessageType: "err_game_token_mismatch"})
	return string(respBytes)
}

func Error121() string {
	respBytes, _ := json.Marshal(&BaseJsonResponse{
		Code:        121,
		MessageType: "err_game_not_exist"})
	return string(respBytes)
}

func Error131() string {
	respBytes, _ := json.Marshal(&BaseJsonResponse{
		Code:        131,
		MessageType: "err_user_record_not_exist"})
	return string(respBytes)
}
