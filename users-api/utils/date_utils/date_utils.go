package date_utils

import "time"

const apiDateLayout = "2006-01-02T15:04:05.000Z"

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

//GetNow utility was written as a separate function
//to provide an easy access point if one would like to
//customize format of displayed data.
func GetNow() time.Time {
	return time.Now().UTC()
}
