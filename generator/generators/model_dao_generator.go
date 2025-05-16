package generators

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/booyangcc/gorm-gen/generator"
)

func init() {
	generator.RegisterGenerator(new(ModelDaoGenerator))
}

type ModelDaoGenerator struct{}

type modelDaoConfig struct {
	DaoPackage       string
	ModelPackage     string
	ModelPackagePath string
	ModelName        string
}

func (g *ModelDaoGenerator) Gen(config generator.GenConfig) error {

	// 生成每个模型的 Dao
	for _, model := range config.Models {
		outputPath := filepath.Join(config.OutputPath, fmt.Sprintf("%s_dao.go", generator.ToSnakeCase(model.ModelName)))

		cfg := modelDaoConfig{
			ModelPackagePath: config.ModePackagePath,
			ModelPackage:     path.Base(config.ModePackagePath),
			DaoPackage:       path.Base(config.DaoPackagePath),
			ModelName:        model.ModelName,
		}

		err := generator.GenerateFile(
			"template/model_dao.tmpl",
			outputPath,
			cfg,
		)
		if err != nil {
			fmt.Printf("Error generating Dao for %s: %v\n", model.ModelName, err)
			return err
		}

		fmt.Println(outputPath + " gen success")
	}
	return nil
}

func (g *ModelDaoGenerator) Name() string {
	return "model_dao_generator"
}
