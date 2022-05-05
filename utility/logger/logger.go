package logger

import (
	"fmt"
	"identification_email/constants"
	"time"
)

func I(inputs ...interface{}) {
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"), " | ", constants.INFO, " | ")
	fmt.Println(inputs...)
}

func D(inputs ...interface{}) {
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"), " | ", constants.DEBUG, " | ")
	fmt.Println(inputs...)
}

func E(inputs ...interface{}) {
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"), " | ", constants.ERROR, " | ")
	fmt.Println(inputs...)
}
