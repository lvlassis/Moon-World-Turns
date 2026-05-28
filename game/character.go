package game

import (
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type characterData struct {
	Name  string `yaml:"name"`
	Image string `yaml:"image"`
	MaxLife int `yaml:"life"`
	Strength int `yaml:"strength"`
	Speed int `yaml:"speed"`
}

type Character struct {
	characterData
	Sprite 		*graphic.Sprite
	Life 		int
}



func LoadCharacter(path string) *Character {

	// Lê o arquivo
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Erro ao ler arquivo %s: %s", path, err)
	}

	// Faz unmarshal do YAML
	var charData characterData
	err = yaml.Unmarshal(data, &charData)
	if err != nil {
		log.Fatalf("Erro ao parsear YAML de %s: %s", path, err)
	}

	// Carrega a imagem em uma textura
	char_texture, err := texture.NewTexture2DFromImage(charData.Image)
	if err != nil {
		log.Fatalf("Falha ao carregar %s: %s", charData.Image, err)
	}
	
	// Ajusta os parâmetros da textura
	char_texture.SetMagFilter(gls.NEAREST)
	char_texture.SetMinFilter(gls.NEAREST)

	// Cria material com a textura
	char_material := material.NewStandard(math32.NewColor("white"))
	char_material.AddTexture(char_texture)
	char_material.SetTransparent(true)

	// Cria Sprite com Material
	var proportion float32 = float32(char_texture.Height()) / float32(char_texture.Width())
	char_sprite := graphic.NewSprite(1, proportion, char_material)

	// Cria e retorna o Character usando os dados lidos
	return &Character{
		characterData: charData, 
		Life: charData.MaxLife, 
		Sprite: char_sprite,
	}
}


