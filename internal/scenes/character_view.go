package scenes

import (
	"log"
	"lvlassis/moon-world-turns/internal/game"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
)

type CharacterView struct {
	CharacterID string
	Sprite *graphic.Sprite
}

func NewCharacterView(charData *game.CharacterData) *CharacterView {

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

	return &CharacterView{
		CharacterID: charData.ID,
		Sprite: char_sprite,
	}
}
