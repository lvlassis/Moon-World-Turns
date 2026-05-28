package main

import (
	"fmt"
	"log"
	"lvlassis/moon-world-turns/game"
	"time"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
)


func DesenhaStatus(char *game.Character) string {
	return fmt.Sprintf("%s HP: %d / %d", char.Name, char.Life, char.MaxLife)
}

func main() {
	g := game.NewGame()
	cam := g.Cam
	scene := g.Scene

	// Cria Char 1
	// Load Ness sprite texture
	char1 := game.LoadCharacter("./data/charmander.yaml")
	char1.Sprite.SetPosition(-1, 0.5, 0)
	scene.Add(char1.Sprite)
	g.Player = char1

	char2 := game.LoadCharacter("./data/squirtle.yaml")
	char2.Sprite.SetPosition(1, 0.5, 0)
	scene.Add(char2.Sprite)
	g.Opponent = char2

	stage := game.LoadStage("./data/stage.yaml")
	scene.Add(stage)

	scene.Add(helper.NewAxes(0.5))

	// ------- UI ----------

	// Indicador de vida
	lifeText := gui.NewLabel("")
	lifeText.SetPosition(10, 10)
	lifeText.SetColor(math32.NewColor("white"))
	text := fmt.Sprintf("%s HP: %d / %d | %s HP: %d / %d", char1.Name, char1.Life, char1.MaxLife, char2.Name, char2.Life, char2.MaxLife)
	lifeText.SetText(text)	
	scene.Add(lifeText)

	// Create button in the bottom-right corner
	button := gui.NewButton("Attack")
	button.SetSize(120, 40)
	button.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		log.Println("Attacking")
		char2.Life -= char1.Strength
	})
	scene.Add(button)

	// Position button in the bottom-right corner (updated on resize)
	positionButton := func() {
		_, height := g.GetSize()
		button.SetPosition(float32(10), float32(height-50))
	}
	positionButton()

	// Update button position when window is resized
	g.Subscribe(window.OnWindowSize, func(evname string, ev interface{}) {
		positionButton()
	})


	// Set background color to gray
	g.Gls().ClearColor(0.8, 0.8, 0.8, 1.0)

	// Run the application
	g.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		// Limpa a Tela
		g.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)

		// Atualiza o estado da UI
		text := fmt.Sprintf("%s HP: %d / %d | %s HP: %d / %d", char1.Name, char1.Life, char1.MaxLife, char2.Name, char2.Life, char2.MaxLife)
		lifeText.SetText(text)	

		// Renderiza a cena
		renderer.Render(scene, cam)
	})
}
