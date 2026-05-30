package game

import (
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
)

type Game struct {
	*app.Application
	Scene 		*core.Node 
	Cam 		*camera.Camera
	Player 		*Character
	Opponent 	*Character
}

func NewGame(player string, opponent string, stage string) *Game {
	/* 
		Instancia o aplicativo, a cena e a câmera.
		Faz todo o boilerplate inicial necessário.
	*/

	// Create application and scene
	a := app.App()
	scene := core.NewNode()

	// Set the scene to be managed by the gui manager
	gui.Manager().Set(scene)

	// Create perspective camera
	cam := camera.New(1)
	cam.SetPosition(0, 1, 6)
	scene.Add(cam)

	// Set up orbit control for the camera
	camera.NewOrbitControl(cam)

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		cam.SetAspect(float32(width) / float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	// Set up callback to close the app when Escape is pressed
	a.Subscribe(window.OnKeyDown, func(evname string, ev interface{}) {
		kev := ev.(*window.KeyEvent)
		if kev.Key == window.KeyEscape || kev.Key == window.KeyQ {
			a.Exit()
		}
	})

	// Instancia o player, o oponente e o stage usando os arquivos YAML
	playerChar := LoadCharacter(player)
	playerChar.Sprite.SetPosition(-1, 0.5, 0)
	scene.Add(playerChar.Sprite)

	opponentChar := LoadCharacter(opponent)
	opponentChar.Sprite.SetPosition(1, 0.5, 0)
	scene.Add(opponentChar.Sprite)

	stageObj := LoadStage(stage)
	scene.Add(stageObj)

	// Adiciona eixos para referência
	scene.Add(helper.NewAxes(0.5))

	// Instancia a UI
	
	return &Game{
		Application: a,
		Scene: scene,
		Cam: cam,
		Player: playerChar,
		Opponent: opponentChar,
	}
}

func (*Game) Update(deltaTime time.Duration) {

}

func (*Game) Draw(renderer *renderer.Renderer, deltaTime time.Duration) {

}


