package webHandler

import (
	"identification_email/utility/logger"
	"identification_email/webHandler/user"
)

func InitWebHandler() {
	logger.I("InitWebHandler invoked")
	defer logger.I("InitWebHandler returned")

	user.RequestHandler()
}
