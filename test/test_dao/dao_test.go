package test_dao

import (
	"os"
	"testing"

	"github.com/booyangcc/gorm-gen/test/test_model"
)

func Test_Dao(t *testing.T) {
	dsn := os.Getenv("siqi")
	d, err := NewDaoManager(dsn, "dev")
	if err != nil {
		panic(err)
	}

	u := &test_model.AdminUser{
		UserName: "test",
		Role:     "a",
		IsActive: false,
	}
	d.AdminUserDao.Create(u)
}
