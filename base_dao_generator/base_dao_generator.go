package base_dao_generator

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/booyangcc/gorm-gen/generator"
)

func init() {
	generator.RegisterGenerator(new(BaseDaoGenerator))
}

type BaseDaoGenerator struct{}

type baseDaoConfig struct {
	DaoPackage string
}

func (g *BaseDaoGenerator) Gen(config generator.GenConfig) error {
	outputPath := filepath.Join(config.OutputPath, "base_dao.go")
	err := generator.GenerateFile(
		"base_dao_generator/base_dao.tmpl",
		outputPath,
		baseDaoConfig{DaoPackage: path.Base(config.DaoPackagePath)},
	)

	if err != nil {
		fmt.Println("Error generating BaseDAO:", err)
		return err
	}

	fmt.Println(outputPath + " gen success")
	return nil
}

func (g *BaseDaoGenerator) Name() string {
	return "base_dao_generator"
}
