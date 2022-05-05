package main

import (
	"identification_email/config"
	"identification_email/constants"
	"identification_email/dbAccess"
	"identification_email/utility/logger"
	"identification_email/webHandler"
	"log"
	"net/http"
)

// main ...
func main() {
	logger.I("main invoked")
	defer logger.I("main returned")

	initMain()
	defer deinitMain()

	// server will listen on 8080 port
	log.Println("listen on", constants.ConstPort)
	log.Fatal(http.ListenAndServe(constants.ConstPort, nil))
}

func initMain() {
	logger.I("initMain invoked")
	defer logger.I("initMain returned")

	webHandler.InitWebHandler()
	dbAccess.ConnectDB()
	config.GetGlobalConfigMap()
}

func deinitMain() {
	logger.I("deinitMain invoked")
	defer logger.I("deinitMain returned")

	dbAccess.DisconnectDB()
}
