package test_model

import "time"



type AdminUser struct {
	UserName      string `gorm:"type:varchar(60);uniqueIndex"`
	Role          string `gorm:"type:varchar(60);"`
	IsActive      bool   `gorm:"type:int(2);"`
	Password      string `gorm:"type:varchar(256);"`
	AvatarURL     string `gorm:"type:varchar(256);"`
	LastLoginDate time.Time
}

func (a *AdminUser) TableName() string {
	return "admin_user"
}
