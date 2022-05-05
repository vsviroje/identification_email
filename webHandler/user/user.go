package user

import (
	"identification_email/constants"
	"identification_email/service"
	"identification_email/utility"
	"identification_email/utility/logger"
	"net/http"
)

func RequestHandler() {
	logger.I("RequestHandler invoked")
	defer logger.I("RequestHandler returned")

	signUp := http.HandlerFunc(service.SignUp)
	http.Handle(constants.ConstUserSignUp, utility.Handler(signUp))

	logIn := http.HandlerFunc(service.Login)
	http.Handle(constants.ConstUserLogin, utility.Handler(logIn))

	logOut := http.HandlerFunc(service.Logout)
	http.Handle(constants.ConstUserLogout, utility.AuthorizedHandler(logOut))

	test := http.HandlerFunc(service.Test)
	http.Handle(constants.ConstUserTest, utility.AuthorizedHandler(test))
}
