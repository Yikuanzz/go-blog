package database

import "github.com/yikuanzz/go-blog/global"

// JwtBlacklist JWT 黑名单表
type JwtBlacklist struct {
	global.MODEL
	Jwt string `json:"jwt" gorm:"type:text"` // Jwt
}
