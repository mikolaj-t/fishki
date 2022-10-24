package core

import "time"

/*
	type Date struct {
		Day   uint8  `json:"day" bson:"day,omitempty"`
		Month uint8  `json:"month" bson:"month,omitempty"`
		Year  uint16 `json:"year" bson:"year,omitempty"`
	}
*/

type Date uint32

func Today() Date {
	now := time.Now()
	zeroDay := time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC)

	date := Date(now.Sub(zeroDay).Hours() / 24)
	return date
}
