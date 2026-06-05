package game

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
	Actions []string `yaml:"actions"`
}

type Character struct {
	CharacterData
	Life int
	Actions []*Action
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
	c := &Character{
		CharacterData: *charData,
		Life: charData.MaxLife,
	}

	// Carrega as ações do Personagem
	for _, action_id := range charData.Actions {
		log.Printf("Carregando ação %s para personagem %s\n", action_id, charData.Name)
		ar := GetActionRegistry()
		action, ok := ar.Get(action_id)
		if !ok {
			log.Printf("Ação %s não encontrada para personagem %s\n", action_id, charData.Name)
			continue
		}
		c.Actions = append(c.Actions, action)
	}

	return c
}

func (c *Character) Attack(target *Character) {
	damage := c.Strength
	target.Life -= damage
	if target.Life < 0 {
		target.Life = 0
	}
	log.Printf("%s atacou %s causando %d de dano! Vida restante de %s: %d\n", c.Name, target.Name, damage, target.Name, target.Life)
}
