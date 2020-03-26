package models

import "time"

// Week model
type Week struct {
	Year int
	Week int
}

// Day model
type Day struct {
	Year int
	Week int
	Day  time.Weekday
}
