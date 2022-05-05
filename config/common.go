package config

import (
	"identification_email/constants"
	"identification_email/dbAccess"
	"identification_email/utility/logger"
)

// Configs ...
var Configs map[string]map[string]string

// GetGlobalConfigMap ...
// function will return global config map from config table
// if any error occur it will return error
func GetGlobalConfigMap() (err error) {
	logger.I("GetGlobalConfigMap invoked")
	defer logger.I("GetGlobalConfigMap returned")

	Configs = make(map[string]map[string]string)
	// getting configs from database

	configArray, err := dbAccess.GetConfigFromDB()
	if err != nil {
		logger.E("GetConfigFromDB Failed for ", err)
		return
	}

	// ranging through config
	for _, data := range configArray {
		if Configs[data.Type] == nil {
			Configs[data.Type] = make(map[string]string)
		}
		Configs[data.Type][data.Key] = data.Value
	}

	return
}

func GetEnvMode() string {
	logger.I("GetEnvMode", constants.ENV, constants.MODE, Configs[constants.ENV][constants.MODE])
	return Configs[constants.ENV][constants.MODE]
}

func GetValueByEnvMode(Type string, mode string, key string) string {
	logger.I("GetValueByEnvMode invoked")
	defer logger.I("GetValueByEnvMode returned")

	logger.I("type:"+(Type+"/"+mode)+" key:"+key, " : ", Configs[Type+"/"+mode][key])

	return Configs[Type+"/"+mode][key]

}
