package scenes

import (
	"lvlassis/moon-world-turns/internal/events"

	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
)

type BaseScene struct {
	*core.Node
	Cam *camera.Camera
	Bus *events.Bus

	// GUI
	GuiRoot *gui.Panel
	GuiCam *camera.Camera
}

func NewBaseScene(bus *events.Bus, width, height int) Scene {
	scene := &BaseScene{
		Node: core.NewNode(),
		Bus: bus,
	}

	// Setup GUI
	scene.setupGUI(width, height)

	return scene
}

func (bs *BaseScene) setupGUI(width, height int) {
	// Cria câmera ortográfica para GUI 2D
	bs.GuiCam = camera.New(1)
	bs.GuiCam.SetPositionVec(&math32.Vector3{0, 0, 1})
	bs.GuiCam.UpdateMatrix()

	// Cria painel root para GUI
	bs.GuiRoot = gui.NewPanel(float32(width), float32(height))
	bs.GuiRoot.SetColor4(&math32.Color4{0, 0, 0, 0}) // Transparente

	// Adiciona o root da GUI à cena
	bs.Add(bs.GuiRoot)
}

func (bs *BaseScene) GetNode() *core.Node {
	return bs.Node
}

func (bs *BaseScene) Update(deltaTime float64) { }

func (bs *BaseScene) OnResize(width, height int) {
	// Ajusta câmera 3D
	if bs.Cam != nil {
		bs.Cam.SetAspect(float32(width) / float32(height))
	}

	// Ajusta câmera GUI
	bs.GuiCam.SetAspect(float32(width) / float32(height))

	// Ajusta tamanho do painel GUI
	bs.GuiRoot.SetSize(float32(width), float32(height))
}

func (bs *BaseScene) Render(renderer *renderer.Renderer) {
	// Renderiza o mundo 3D
	if bs.Cam != nil {
		renderer.Render(bs.Node, bs.Cam)
	}

	// Renderiza a GUI 2D
	renderer.Render(bs.GuiRoot, bs.GuiCam)
}

