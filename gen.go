package gormgen

import (
	"fmt"
	"os/exec"
	"strings"

	_ "github.com/booyangcc/gorm-gen/base_dao_generator"
	_ "github.com/booyangcc/gorm-gen/dao_manager_generator"
	"github.com/booyangcc/gorm-gen/generator"
	_ "github.com/booyangcc/gorm-gen/model_dao_generator"
)

func Gen(config generator.GenConfig) {
	pkgPath, err := getCurrentPackagePath()
	if err != nil {
		panic(err)
	}

	if len(config.Models) == 0 {
		modelPath := strings.TrimPrefix(config.ModePackagePath, pkgPath+"/")
		modelStrs, err := generator.FindStructs(modelPath)
		if err != nil {
			fmt.Println("Error find struct:", err)
			panic(err)
		}
		config.Models = make([]generator.ModelInfo, 0, len(modelStrs))
		for _, modelName := range modelStrs {
			config.Models = append(config.Models, generator.ModelInfo{ModelName: modelName})
		}
	}

	config.OutputPath = strings.TrimPrefix(config.DaoPackagePath, pkgPath+"/")

	generators := generator.GetGenerator()
	for _, gen := range generators {
		err := gen.Gen(config)
		if err != nil {
			panic(err)
		}
	}
}

func getCurrentPackagePath() (string, error) {
	cmd := exec.Command("go", "list", ".")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
