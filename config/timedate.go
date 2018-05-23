package config

import (
	"fmt"
	"strings"
	"time"
)

// Format is the format the date time variable in messages must follow
// when converting from Timestamp to a readable date and time format.
// It must follow this format to comply with the Go specifications and
// requirements to work properly.
const Format = "2006-01-02 15:04:05"

// GetTimeNowTimeStamp function that returns a timestamp of the current time.
// It picks up the time as time type, parses it to string, splits the contents to fit
// the format we want to convert, and calls ConvertDateToTimeStamp to convert it
// into a timestamp.
func GetTimeNowTimeStamp() int64 {
	timenow := time.Now().String()
	timenowSubstring := strings.Split(timenow, ".")

	value := convertDateToTimestamp(timenowSubstring[0])

	return value
}

// convertDateToTimestamp function that receives a date and parses it's format
// to return a unix timestamp of the date.
func convertDateToTimestamp(date string) int64 {

	t, err := time.Parse(Format, date)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t.Unix())
		return t.Unix()
	}
	return 0
}
