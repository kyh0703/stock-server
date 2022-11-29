package dto

import "time"

type AccessTokenDto struct {
	Token  string
	Expire time.Time
}
