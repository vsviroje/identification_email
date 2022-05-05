package dbAccess

import (
	"database/sql"
	"identification_email/models"
	"identification_email/utility/logger"
)

// GetConfigFromDB ...
// function will get all config from database if any error occur it will return
func GetConfigFromDB() (configArray []models.Config, err error) {
	logger.I("GetConfigFromDB invoked")
	defer logger.I("GetConfigFromDB returned")

	sqlSmt := "select type,`key`,value from config"
	// getting sql rows from data bases
	var sqlRow *sql.Rows
	if sqlRow, err = DB.Query(sqlSmt); err != nil {
		logger.E("DB.Query Failed for ", sqlSmt, err)
		return
	}
	defer sqlRow.Close()
	// iterating through rows
	for sqlRow.Next() {
		var tempData models.Config
		sqlRow.Scan(&tempData.Type, &tempData.Key, &tempData.Value)
		configArray = append(configArray, tempData)
	}
	return
}
