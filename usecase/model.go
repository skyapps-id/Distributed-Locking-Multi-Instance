package usecase

import "time"

type GetPromoResponse struct {
	UUID      string    `json:"uuid"`
	Code      string    `json:"code"`
	Qouta     int       `json:"qouta"`
	UpdatedAt time.Time `json:"updated_at"`
}
