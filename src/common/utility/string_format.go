package utility

import (
	"pemuda-peduli/src/common/constants"
	"strings"
	"time"
)

func FormatPhoneNumber(phoneNumber string) (data string) {
	prefix := phoneNumber[0:3]
	if prefix != "+62" {
		phoneNumber = strings.Replace(phoneNumber, "0", "+62", 1)
	}
	data = phoneNumber
	return
}

func UsernameIsPhoneNumber(username string) (data string) {
	prefix := username[0:2]
	if prefix == "08" {
		prefix = username[0:3]
		if prefix != "+62" {
			username = strings.Replace(username, "0", "+62", 1)
		}
		data = username
		return
	} else {
		data = username
		return
	}
}

func StringDateFormat(stringDate string) (response time.Time, err error) {
	stringDate += "T00:00:00.000Z"
	response, err = time.Parse(constants.DefaultFormatFullDateTime, stringDate)
	return
}

func DateFormat(dateTime time.Time) (response string) {
	response = dateTime.Format(constants.DefaultFormatDate)
	return
}
