package generator

type ModelInfo struct {
	ModelName string
}

type GenConfig struct {
	Models          []ModelInfo
	ModePackagePath string
	DaoPackagePath  string
	OutputPath      string
}
