package main

import (
	"fmt"
	"time"
)

func main() {

	const (
		layoutTime   = "15:04:05"
		startedTime1 = "10:00:00"
		startedTime2 = "23:59:59"
		endedTime1   = "00:00:00"
		endedTime2   = "15:31:00"
	)
	isNextDay := false
	timeLocation, _ := time.LoadLocation("Asia/Jakarta")

	// timeNow := time.Now().Format(layoutTime)
	timeNow := time.Now().In(timeLocation).Format(layoutTime)
	// timeNow := "22:30:00"

	if isNextDay == true {
		if timeNow > startedTime1 && timeNow < startedTime2 {

			fmt.Println("True", timeNow)

		} else if timeNow > endedTime1 && timeNow < endedTime2 {

			fmt.Println("True", timeNow)

		} else {

			fmt.Println("False", timeNow)

		}
	} else {
		if timeNow > startedTime1 && timeNow < endedTime2 {
			fmt.Println("True", timeNow)
		} else {
			fmt.Println("False", timeNow)
		}
	}

}
