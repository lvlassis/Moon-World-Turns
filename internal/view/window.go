package view

import (
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"
)

type Window struct {
	app *app.Application
	scene Scene
}

func NewWindow(scene Scene, app *app.Application) *Window {

	// Callback para redimensionar a viewport
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := app.GetSize()
		app.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		scene.OnResize(width, height)
	}
	app.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	// Callback para cliques de mouse
	onMouseDown := func(evname string, ev interface{}) {
		mouseEv := ev.(*window.MouseEvent)

		// Converte para float32
		x := float32(mouseEv.Xpos)
		y := float32(mouseEv.Ypos)

		// Se a cena for BattleScene, propaga o clique para a HUD
		if bs, ok := scene.(*BattleScene); ok {
			bs.GetHUD().HandleClick(x, y)
		}
	}
	app.Subscribe(window.OnMouseDown, onMouseDown)

	// Set background color to gray
	app.Gls().ClearColor(0.8, 0.8, 0.8, 1.0)

	return &Window{
		app: app,
		scene: scene,
	}
}

func (w *Window) Run() {
	w.app.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		// Limpa a Tela
		w.app.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)

		// Atualzia a cena
		w.scene.Update(deltaTime.Seconds())

		// Renderiza a cena (3D + GUI)
		w.scene.Render(renderer)
	})
}
