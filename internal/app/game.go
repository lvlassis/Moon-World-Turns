package app

import (
	"fmt"
	"lvlassis/moon-world-turns/internal/engine"
	"lvlassis/moon-world-turns/internal/view"

	"github.com/g3n/engine/core"
)

func Run() {
	// Carrega os dados dos personagens
	fmt.Println("Loading Charmander")
	charmander_data := engine.LoadCharacterData("./data/charmander.yaml")
	fmt.Println("Loading Squirtle")
	squirtle_data := engine.LoadCharacterData("./data/squirtle.yaml")

	// Cria os personagens
	charmander := engine.NewCharacter(charmander_data)
	fmt.Println("Charmander:", charmander)
	squirtle := engine.NewCharacter(squirtle_data)
	fmt.Println("Squirtle:", squirtle)

	// Cria as views dos personagens
	charmander_view := view.NewCharacterView(charmander_data)
	squirtle_view := view.NewCharacterView(squirtle_data)

	// Monta uma cena com os personagens
	scene := core.NewNode()
	scene.Add(charmander_view.Sprite)
	scene.Add(squirtle_view.Sprite)

	// Renderiza
	window := view.NewWindow(scene)
	window.Run()
}
