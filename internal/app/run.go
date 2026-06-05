package app

import (
	"fmt"
	"lvlassis/moon-world-turns/internal/events"
	"lvlassis/moon-world-turns/internal/game"
	"lvlassis/moon-world-turns/internal/scenes"

	"github.com/g3n/engine/app"
)

func Run() {
	character1_path := "./data/charmander.yaml"
	character2_path := "./data/squirtle.yaml"
	stage_path := "./data/stage.yaml"

	ar := game.GetActionRegistry()
	fmt.Println(ar)

	// Cria um app para gerenciar a janela e o loop de renderização
	a := app.App()

	// Cria o event bus
	bus := events.NewBus()

	// Carrega os dados dos personagens
	fmt.Println("Loading Charmander")
	character1_data := game.LoadCharacterData(character1_path)
	fmt.Println("Loading Squirtle")
	character2_data := game.LoadCharacterData(character2_path)

	// Carrega os dados do estágio
	fmt.Println("Loading Stage")
	stage_data := game.LoadStageData(stage_path)

	// Cria os personagens
	character1 := game.NewCharacter(character1_data)
	fmt.Println("Character 1:", character1)
	character2 := game.NewCharacter(character2_data)
	fmt.Println("Character 2:", character2)

	// Cria o estágio
	// stage := game.NewStage(stage_data)

	// Cria a batalha
	battle := game.NewBattle(character1, character2, bus)

	// Cria as views dos personagens
	character1_view := scenes.NewCharacterView(character1_data)
	character2_view := scenes.NewCharacterView(character2_data)

	// Cria a view do estágio
	stage_view := scenes.NewStageView(stage_data)

	// Obtém o tamanho inicial da janela
	width, height := a.GetSize()

	// Monta uma cena com os personagens
	battleScene := scenes.NewBattleScene(character1_view, character2_view, stage_view, battle, bus, width, height)

	sceneManager := scenes.NewSceneManager()
	sceneManager.AddScene("battle", battleScene)
	sceneManager.SetCurrentScene("battle")

	// Cria a janela
	window := scenes.NewWindow(sceneManager, a)

	// Inicia a batalha
	battle.Start()

	// Renderiza
	window.Run()
}
