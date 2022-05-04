package user

import (
	"identification_email/constants"
	"identification_email/service"
	"net/http"
)

func RequestHandler() {

	signUp := http.HandlerFunc(service.SignUp)
	http.Handle(constants.ConstUserSignUp, handler(signUp))

	logIn := http.HandlerFunc(service.Login)
	http.Handle(constants.ConstUserLogin, handler(logIn))

	logOut := http.HandlerFunc(service.LogOut)
	http.Handle(constants.ConstUserLogout, handler(logOut))

}

// handler
func handler(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		originalHandler.ServeHTTP(w, r)
	})
}
