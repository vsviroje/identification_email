package config

import "github.com/spf13/cast"

var UserSession map[string]string

func init() {
	UserSession = make(map[string]string)
}

func AddUserInSession(userID int, token string) {
	UserSession[cast.ToString(userID)] = token
}

func DeleteUserFromSession(userID int) {
	delete(UserSession, cast.ToString(userID))
}

func IsUserInSession(userID int) bool {
	_, isOk := UserSession[cast.ToString(userID)]
	return isOk
}
