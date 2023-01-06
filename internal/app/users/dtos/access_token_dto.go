package dtos

import "time"

type AccessTokenDto struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}
