package gormgen

import (
	"testing"

	"github.com/booyangcc/gorm-gen/generator"

)

func TestGen(t *testing.T) {
	config := generator.GenConfig{
		ModePackagePath: "github.com/booyangcc/gorm-gen/test/test_model",
		DaoPackagePath:  "github.com/booyangcc/gorm-gen/test/test_dao",
	}
	Gen(config)
}
