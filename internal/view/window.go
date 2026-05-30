package view

import (
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"
)

type Window struct {
	app *app.Application
	scene *core.Node
	cam *camera.Camera
}

func NewWindow(scene *core.Node) *Window {
	// Cria um app para gerenciar a janela e o loop de renderização
	app := app.App()

	// Cria uma câmera
	cam := camera.New(1)
	cam.SetPosition(0, 1, 6)
	camera.NewOrbitControl(cam)
	scene.Add(cam)

	// Callback para redimensionar a viewport
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := app.GetSize()
		app.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		cam.SetAspect(float32(width) / float32(height))
	}
	app.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	// Set background color to gray
	app.Gls().ClearColor(0.8, 0.8, 0.8, 1.0)

	return &Window{
		app: app,
		scene: scene,
		cam: cam,
	}
}

func (w *Window) Run() {
	w.app.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		// Limpa a Tela
		w.app.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)

		// Renderiza a cena
		renderer.Render(w.scene, w.cam)
	})
}
