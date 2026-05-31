package engine

import (
	"log"
	"os"
	"gopkg.in/yaml.v3"
)

type StageData struct {
	StageName  string     `yaml:"name"`
	Ground     groundData `yaml:"ground"`
	Lights     []lightData `yaml:"lights"`
}

type Stage struct {
	StageData
}

type groundData struct {
	Image string `yaml:"image"`
	Color string `yaml:"color"`
}

type lightData struct {
	Type  string  `yaml:"type"`
	Color struct {
		Red   float32 `yaml:"red"`
		Green float32 `yaml:"green"`
		Blue  float32 `yaml:"blue"`
	}  
	Intensity float32 `yaml:"intensity"`
	Position struct {
		X float32 `yaml:"x"`
		Y float32 `yaml:"y"`
		Z float32 `yaml:"z"`
	}
}

func LoadStageData(path string) *StageData {
	// Lê o arquivo
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Erro ao ler Stage %s: %s", path, err)
	}

	var stageData StageData
	err = yaml.Unmarshal(data, &stageData)
	if err != nil {
		log.Fatalf("Erro ao parsear YAML de %s: %s", path, err)
	}

	return &stageData
}

func NewStage(stageData *StageData) *Stage {
	return &Stage{
		StageData: *stageData,
	}
}
