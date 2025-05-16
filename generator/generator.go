package generator

import (
	"bytes"
	"embed"
	"fmt"

	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed template/*.tmpl
var templateFS embed.FS

type Generator interface {
	Gen(config GenConfig) error
	Name() string
}

var generators = make(map[string]Generator)

func RegisterGenerator(p Generator) {
	generators[p.Name()] = p
}

func GetGenerator() map[string]Generator {
	return generators
}

func GenerateFile(templatePath, outputPath string, data interface{}) error {
	tmpl, err := template.ParseFS(templateFS, templatePath)
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
func FindStructsFromFile(filePath string) ([]string, error) {
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
func FindStructs(path string) ([]string, error) {
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
				structs, err := FindStructsFromFile(fp)
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
		return FindStructsFromFile(path)
	}

	return nil, fmt.Errorf("不是有效的 Go 文件或目录: %s", path)
}

// 将驼峰命名转换为蛇形命名
func ToSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return string(bytes.ToLower([]byte(string(result))))
}
