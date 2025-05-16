package test_dao

import (
	"os"
	"testing"
	"time"

	"github.com/booyangcc/gorm-gen/example/test_model"
)

func Test_Dao(t *testing.T) {
	dsn := os.Getenv("siqi_dev")
	d, err := NewDaoManager(dsn, "dev")
	if err != nil {
		panic(err)
	}

	u := &test_model.AdminUser{
		UserName:      "test",
		Role:          "a",
		IsActive:      false,
		LastLoginDate: time.Now(),
	}
	d.AdminUserDao.Create(u)
}
