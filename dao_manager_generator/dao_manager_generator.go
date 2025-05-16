package dao_manager_generator

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/booyangcc/gorm-gen/generator"
)

func init() {
	generator.RegisterGenerator(new(DaoManagerGenerator))
}

type DaoManagerGenerator struct{}

type daoManagerConfig struct {
	DaoPackage       string
	ModelNames       []string
	ModelPackage     string
	ModelPackagePath string
}

func (g *DaoManagerGenerator) Gen(config generator.GenConfig) error {
	modelNames := make([]string, 0)
	for _, v := range config.Models {
		modelNames = append(modelNames, v.ModelName)
	}

	cfg := daoManagerConfig{
		DaoPackage:       path.Base(config.DaoPackagePath),
		ModelNames:       modelNames,
		ModelPackage:     path.Base(config.ModePackagePath),
		ModelPackagePath: config.ModePackagePath,
	}
	outputPath := filepath.Join(config.OutputPath, "dao_manager.go")
	err := generator.GenerateFile(
		"dao_manager_generator/dao_manager.tmpl",
		outputPath,
		cfg,
	)

	if err != nil {
		fmt.Println("Error generating BaseDAO:", err)
		return err
	}
	
	fmt.Println(outputPath + " gen success")
	return nil
}

func (g *DaoManagerGenerator) Name() string {
	return "dao_manager_generator"
}
