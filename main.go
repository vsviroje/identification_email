package main

import (
	"identification_email/config"
	"identification_email/constants"
	"identification_email/dbAccess"
	"identification_email/webHandler"
	"log"
	"net/http"
)

// main ...
func main() {

	initMain()
	defer deinitMain()

	// server will listen on 8080 port
	log.Println("listen on", constants.ConstPort)
	log.Fatal(http.ListenAndServe(constants.ConstPort, nil))
}

func initMain() {
	webHandler.InitWebHandler()
	dbAccess.ConnectDB()
	config.GetGlobalConfigMap()
}

func deinitMain() {
	dbAccess.DisconnectDB()
}
