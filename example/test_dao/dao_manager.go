package test_dao

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

    "github.com/booyangcc/gorm-gen/example/test_model"

)

type DaoManager struct {
	DB                *gorm.DB
	AdminUserDao *AdminUserDao
}

func newDaoManager(db *gorm.DB) *DaoManager{
	return &DaoManager{
		DB:                db,
		AdminUserDao : NewAdminUserDao(db),
	}
}

func NewDaoManager(dsn string, mode string) (*DaoManager, error) {
	gormConfig := &gorm.Config{}
	if mode == "dev" {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      false,
				Colorful:                  true,
			},
		)
		gormConfig.Logger = newLogger
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), gormConfig)
	if err != nil {
		return nil, err
	}

    models := []interface{}{new(test_model.AdminUser),
    }
	err = db.AutoMigrate(models...)
	if err != nil {
		return nil, err
	}

    if mode == "dev" {
		db = db.Debug()
	}

	return NewDaoManagerWithDB(db)
}

func NewDaoManagerWithDB(db *gorm.DB) (*DaoManager, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	return newDaoManager(db), nil
}
