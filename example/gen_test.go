package gormgen

import (
	"testing"

	gormgen "github.com/booyangcc/gorm-gen"
	"github.com/booyangcc/gorm-gen/generator"
)

func TestGen(t *testing.T) {
	config := generator.GenConfig{
		ModePackagePath: "github.com/booyangcc/gorm-gen/example/test_model",
		DaoPackagePath:  "github.com/booyangcc/gorm-gen/example/test_dao",
	}
	gormgen.Gen(config)
}
