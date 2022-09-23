package entity

import "time"

type Generate struct {
	Email string
	Code  uint32
	Time  time.Time
}
