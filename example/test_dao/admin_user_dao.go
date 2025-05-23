package test_dao

import (
	"gorm.io/gorm"
    "github.com/booyangcc/gorm-gen/example/test_model"
)

type AdminUserDao struct {
	*BaseDao[test_model.AdminUser]
}

func NewAdminUserDao(db *gorm.DB) *AdminUserDao {
	return &AdminUserDao{
		BaseDao: NewBaseDao[test_model.AdminUser](db),
	}
}

