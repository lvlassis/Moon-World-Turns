package app

import (
	"fmt"
	"lvlassis/moon-world-turns/internal/engine"
	"lvlassis/moon-world-turns/internal/events"
	"lvlassis/moon-world-turns/internal/view"

	"github.com/g3n/engine/app"
)

func Run() {
	// Cria um app para gerenciar a janela e o loop de renderização
	a := app.App()

	// Cria o event bus
	bus := events.NewBus()

	// Carrega os dados dos personagens
	fmt.Println("Loading Charmander")
	charmander_data := engine.LoadCharacterData("./data/charmander.yaml")
	fmt.Println("Loading Squirtle")
	squirtle_data := engine.LoadCharacterData("./data/squirtle.yaml")

	// Carrega os dados do estágio
	fmt.Println("Loading Stage")
	stage_data := engine.LoadStageData("./data/stage.yaml")

	// Cria os personagens
	charmander := engine.NewCharacter(charmander_data)
	fmt.Println("Charmander:", charmander)
	squirtle := engine.NewCharacter(squirtle_data)
	fmt.Println("Squirtle:", squirtle)

	// Cria o estágio
	// stage := engine.NewStage(stage_data)

	// Cria a batalha
	battle := engine.NewBattle(charmander, squirtle, bus)

	// Cria as views dos personagens
	charmander_view := view.NewCharacterView(charmander_data)
	squirtle_view := view.NewCharacterView(squirtle_data)

	// Cria a view do estágio
	stage_view := view.NewStageView(stage_data)

	// Obtém o tamanho inicial da janela
	width, height := a.GetSize()

	// Monta uma cena com os personagens
	scene := view.NewBattleScene(charmander_view, squirtle_view, stage_view, battle, bus, width, height)

	// Cria a janela
	window := view.NewWindow(scene, a)

	// Inicia a batalha
	battle.Start()

	// Renderiza
	window.Run()
}
