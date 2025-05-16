package test_model

import (
	"time"
)

type WxUser struct {
	OpenID        string `gorm:"uniqueIndex"`
	UnionID       string `gorm:"index"`
	Nickname      string
	AvatarURL     string
	Gender        int
	Country       string
	Province      string
	IsActive      bool
	City          string
	LastLoginDate time.Time
}

func (d *WxUser) TableName() string {
	return "wx_user"
}
