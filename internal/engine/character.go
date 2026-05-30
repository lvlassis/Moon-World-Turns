package engine

import (
	"log"
	"os"
	"gopkg.in/yaml.v3"
)

type CharacterData struct {
	ID   string `yaml:"id"`
	Name  string `yaml:"name"`
	Image string `yaml:"image"`
	MaxLife int `yaml:"life"`
	Strength int `yaml:"strength"`
	Speed int `yaml:"speed"`
}

type Character struct {
	CharacterData
	Life int
}

func LoadCharacterData(path string) *CharacterData {

	// Lê o arquivo
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Erro ao ler Character %s: %s", path, err)
	}

	// Faz unmarshal do YAML
	var charData CharacterData
	err = yaml.Unmarshal(data, &charData)
	if err != nil {
		log.Fatalf("Erro ao parsear YAML de %s: %s", path, err)
	}
	
	return &charData
}

func NewCharacter(charData *CharacterData) *Character {
	return &Character{
		CharacterData: *charData,
		Life: charData.MaxLife,
	}
}
