package dbAccess

import (
	"database/sql"
	"fmt"
	"identification_email/models"

	"github.com/spf13/cast"
)

// GetOtpBasedonMobileAndOtp ...
// function will get otp row from table based on mobile and otp combination and if any error occur
// it will return error
func GetUserByEmail(email string) (*models.User, error) {
	sqlSmt := "select id,userName,name,email,phoneNum,password,createdAt,updatedAt from user where email = ?"

	data := models.User{}
	// getting sql rows from data bases
	var sqlRow *sql.Rows
	sqlRow, err := DB.Query(sqlSmt, email)
	if err != nil {
		fmt.Println("Error occured while getting sql rows db", err)
		return nil, err
	}
	defer sqlRow.Close()
	isDataPresent := false
	// iterating through rows
	for sqlRow.Next() {
		var createdOn, updatedOn string
		sqlRow.Scan(&data.ID, &data.UserName, &data.Name, &data.Email, &data.PhoneNum, &data.Password, &createdOn, &updatedOn)
		data.CreatedAt = cast.ToTime(createdOn)
		data.UpdatedAt = cast.ToTime(updatedOn)
		isDataPresent = true
	}
	if isDataPresent {
		return &data, nil
	}

	return nil, nil
}

// CreatOTPRow ...
// function will receive top data as input based on that it will create row on table
func CreatUser(data *models.Credentials) (err error) {

	sqlStr := "INSERT INTO user (userName,email,password) VALUES (?,?,?)"
	// prepare the statement
	var stmt *sql.Stmt
	if stmt, err = DB.Prepare(sqlStr); err != nil {
		fmt.Println("Error occured while prparing sql statement", err)
		return
	}
	// format all vals at once
	if _, err = stmt.Exec(data.Email, data.Email, data.Password); err != nil {
		fmt.Println("Error occured while excuting create sql statement", err)
		return
	}
	return
}

// // GetOtpCountBaseonCreatedOn ...
// // function will get otp count based on created on and if any error occur
// // it will return error
// func GetOtpCountBaseonCreatedOn(mobile, createdOn string, resend int) (count int, err error) {
// 	sqlSmt := "select mobile,otp,status,created_on from otp where mobile = ? and created_on >= ? and status = 0 and resend = ?"
// 	// getting sql rows from data bases
// 	var sqlRow *sql.Rows
// 	if sqlRow, err = DB.Query(sqlSmt, mobile, createdOn, resend); err != nil {
// 		fmt.Println("Error occured while getting sql rows db", err)
// 		return
// 	}
// 	defer sqlRow.Close()
// 	// iterating through rows
// 	for sqlRow.Next() {
// 		count++
// 	}
// 	return
// }

// // GetOtpBasedonMobileAndOtp ...
// // function will get otp row from table based on mobile and otp combination and if any error occur
// // it will return error
// func GetOtpBasedonMobileAndOtp(mobile, otp string) (data OTP, err error) {
// 	sqlSmt := "select mobile,otp,status,created_on from otp where mobile = ? and otp = ? order by created_on desc limit 1"
// 	// getting sql rows from data bases
// 	var sqlRow *sql.Rows
// 	if sqlRow, err = DB.Query(sqlSmt, mobile, otp); err != nil {
// 		fmt.Println("Error occured while getting sql rows db", err)
// 		return
// 	}
// 	defer sqlRow.Close()
// 	// iterating through rows
// 	for sqlRow.Next() {
// 		var createdOn string
// 		sqlRow.Scan(&data.Mobile, &data.OTP, &data.Status, &createdOn)
// 		data.CreatedOn = cast.ToTime(createdOn)
// 	}
// 	return
// }

// // UpdateOtpRowBasedOnMobileAndOtp ...
// // function will update otp row based otp and mobile combination and if any error occur it will return error
// func UpdateOtpRowBasedOnMobileAndOtp(mobile, otp string) (err error) {
// 	sqlSmt := "update otp set status = 1 where mobile = ? and otp = ?"
// 	// updating row on database
// 	if _, err = DB.Exec(sqlSmt, mobile, otp); err != nil {
// 		fmt.Println("Error occured while updating otp row status column", err)
// 	}
// 	return
// }
