package config

import (
	"identification_email/dbAccess"
	"identification_email/models"
)

// Configs ...
var Configs map[string]map[string]string

// GetGlobalConfigMap ...
// function will return global config map from config table
// if any error occur it will return error
func GetGlobalConfigMap() (err error) {
	Configs = make(map[string]map[string]string)
	for {
		// getting configs from database
		var configArray []models.Config
		if configArray, err = dbAccess.GetConfigFromDB(); err != nil {
			break
		}
		// ranging through config
		for _, data := range configArray {
			if Configs[data.Type] == nil {
				Configs[data.Type] = make(map[string]string)
			}
			Configs[data.Type][data.Key] = data.Value
		}
		break
	}
	return
}
