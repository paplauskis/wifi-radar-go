package wifi

import "time"

type WifiReviewDTO struct {
	WifiReviewID int32     `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	WifiID       int64     `json:"wifi_id"`
	Text         string    `json:"text"`
	Rating       int8      `json:"rating"`
	Username     string    `json:"username"`
}
