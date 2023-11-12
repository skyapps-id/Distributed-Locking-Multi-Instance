package entity

import "time"

type Promo struct {
	UUID      string
	Code      string
	Qouta     int
	UpdatedAt time.Time
}
