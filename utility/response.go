package utility

import (
	"encoding/json"
	"identification_email/constants"
	"identification_email/utility/logger"
	"net/http"

	"github.com/spf13/cast"
)

// GenrateResponse ...
// function will generate final response
func GenrateResponse(w http.ResponseWriter, data interface{}, err error) {
	logger.I("GenrateResponse invoked")
	defer logger.I("GenrateResponse returned")

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
		logger.E("Marshal Failed for ", resp, err)
		return
	}
	w.Write(jsonResp)
}
