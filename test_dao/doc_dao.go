package test_dao

import (
	"gorm.io/gorm"
    "github.com/booyangcc/gorm-gen/test_model"
)

type DocDao struct {
	*BaseDao[test_model.Doc]
}

func NewDocDao(db *gorm.DB) *DocDao {
	return &DocDao{
		BaseDao: NewBaseDao[test_model.Doc](db),
	}
}

