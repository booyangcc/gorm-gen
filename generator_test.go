package gormgen

import "testing"

func TestGen(t *testing.T) {
	config := GenConfig{
		ModelPath:       "./test_model/",
		ModePackagePath: "github.com/booyangcc/gorm-gen/test_model",
		DaoPackage:      "test_dao",
	}
	Gen(config)
}
