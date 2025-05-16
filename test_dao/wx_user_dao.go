package test_dao

import (
	"gorm.io/gorm"
    "github.com/booyangcc/gorm-gen/test_model"
)

type WxUserDao struct {
	*BaseDao[test_model.WxUser]
}

func NewWxUserDao(db *gorm.DB) *WxUserDao {
	return &WxUserDao{
		BaseDao: NewBaseDao[test_model.WxUser](db),
	}
}

