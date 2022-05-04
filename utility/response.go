package utility

import (
	"encoding/json"
	"identification_email/constants"
	"log"
	"net/http"

	"github.com/spf13/cast"
)

// GenrateResponse ...
// function will generate final response
func GenrateResponse(w http.ResponseWriter, data interface{}, err error) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set(constants.ConstContentType, constants.ConstApplicationSlashJSON)
	resp := make(map[string]interface{})
	resp[constants.ConstMessage] = constants.ConstSuccess
	resp[constants.ConstStatusCode] = constants.Const200
	if err != nil {
		resp[constants.ConstMessage] = cast.ToString(err)
		resp[constants.ConstStatusCode] = constants.Const451
	}
	resp[constants.ConstData] = data
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
