package gormgen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

type ModelInfo struct {
	ModelName string
}

type GenConfig struct {
	Models          []ModelInfo
	ModelPath       string
	ModePackagePath string
	DaoPackage      string
	OutputPath      string
}

type TmplData struct {
	ModelName        string
	ModelPackagePath string
	DaoPackagePath   string
	DaoPackage       string
	ModelPackage     string
	Daos             []string
}

func Gen(config GenConfig) {
	// 生成 BaseDAO
	err := generateFile("templates/base_dao.tmpl", filepath.Join(config.DaoPackage, "base_dao.go"), TmplData{DaoPackage: config.DaoPackage})
	if err != nil {
		fmt.Println("Error generating BaseDAO:", err)
		return
	}

	daos := make([]string, 0)
	if len(config.Models) == 0 {
		modelStrs, err := findStructs(config.ModelPath)
		if err != nil {
			fmt.Println("Error find struct:", err)
			panic(err)
		}
		config.Models = make([]ModelInfo, 0, len(modelStrs))
		for _, modelName := range modelStrs {
			config.Models = append(config.Models, ModelInfo{ModelName: modelName})
			daos = append(daos, modelName)
		}
	}

	err = generateFile("templates/dao.tmpl", filepath.Join(config.DaoPackage, "dao.go"), TmplData{Daos: daos, DaoPackage: config.DaoPackage})
	if err != nil {
		fmt.Println("Error generating BaseDAO:", err)
		return
	}

	// 生成每个模型的 Dao
	for _, model := range config.Models {
		outputPath := filepath.Join(config.DaoPackage, fmt.Sprintf("%s_dao.go", toSnakeCase(model.ModelName)))
		tmplData := TmplData{
			ModelName:        model.ModelName,
			ModelPackagePath: config.ModePackagePath,
			ModelPackage:     path.Base(config.ModePackagePath),
			DaoPackage:       config.DaoPackage,
			Daos:             daos,
		}
		err := generateFile("templates/model_dao.tmpl", outputPath, tmplData)
		if err != nil {
			fmt.Printf("Error generating Dao for %s: %v\n", model.ModelName, err)
		}
	}
}

func generateFile(templatePath, outputPath string, data interface{}) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return err
	}

	// 创建输出目录
	err = os.MkdirAll(filepath.Dir(outputPath), os.ModePerm)
	if err != nil {
		return err
	}

	// 写入文件
	return os.WriteFile(outputPath, buf.Bytes(), 0644)
}

// 提取单个 Go 文件中的所有结构体名
func findStructsFromFile(filePath string) ([]string, error) {
	var structNames []string

	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec := spec.(*ast.TypeSpec)
			if _, ok := typeSpec.Type.(*ast.StructType); ok {
				structNames = append(structNames, typeSpec.Name.Name)
			}
		}
	}

	return structNames, nil
}

// 递归处理目录或处理单个 Go 文件
func findStructs(path string) ([]string, error) {
	var allStructs []string

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		err := filepath.Walk(path, func(fp string, fi os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !fi.IsDir() && strings.HasSuffix(fp, ".go") && !strings.HasSuffix(fp, "_test.go") {
				structs, err := findStructsFromFile(fp)
				if err != nil {
					return err
				}
				allStructs = append(allStructs, structs...)
			}
			return nil
		})
		return allStructs, err
	}

	// 单个文件
	if strings.HasSuffix(path, ".go") {
		return findStructsFromFile(path)
	}

	return nil, fmt.Errorf("不是有效的 Go 文件或目录: %s", path)
}

// 将驼峰命名转换为蛇形命名
func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return string(bytes.ToLower([]byte(string(result))))
}
