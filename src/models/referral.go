package models

import "time"

type Referral struct {
	ID            int       `json:"id"`
	User          string    `json:"user"`
	Hash          string    `json:"hash"`
	PromoCode     string    `json:"promo_code"`
	Count         int       `json:"count"`
	DateGenerated time.Time `json:"date_generated"`
	DateRegistered time.Time `json:"date_registered"`
}
