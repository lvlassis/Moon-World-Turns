package view

import (
	"lvlassis/moon-world-turns/internal/engine"
	"lvlassis/moon-world-turns/internal/events"

	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
)

type BattleScene struct {
	*core.Node
	character1 	*CharacterView
	character2 	*CharacterView
	stage 			*StageView
	battle 			*engine.Battle
	bus 				*events.Bus
	cam 				*camera.Camera
	hud					*BattleHUD
	guiCam			*camera.Camera  // Câmera ortográfica para GUI
}

func NewBattleScene(char1, char2 *CharacterView, stage *StageView, battle *engine.Battle, bus *events.Bus, width, height int) Scene {
	scene := &BattleScene{
		Node: core.NewNode(),
		character1: char1,
		character2: char2,
		battle: battle,
		bus: bus,
	}

	// Posiciona os personagens e o estágio na cena
	scene.character1.Sprite.SetPosition(-1, 0, 0)
	scene.Add(scene.character1.Sprite)
	scene.character2.Sprite.SetPosition(1, 0, 0)
	scene.Add(scene.character2.Sprite)
	stage.SetPosition(0, -0.4, 0)
	scene.Add(stage)
	
	// Cria uma câmera
	scene.cam = camera.New(1)
	scene.cam.SetPosition(0, 1, 6)
	camera.NewOrbitControl(scene.cam)
	scene.Add(scene.cam)

	// Cria câmera ortográfica para GUI 2D
	scene.guiCam = camera.New(1)
	scene.guiCam.SetPositionVec(&math32.Vector3{0, 0, 1})
	scene.guiCam.UpdateMatrix()

	// Cria a HUD
	scene.hud = NewBattleHUD(width, height, battle)

	// IMPORTANTE: Adiciona o root da HUD como node filho para receber eventos
	// (o g3n propaga eventos através da hierarquia de nodes)
	scene.Add(scene.hud.root)

	return scene
}

func (bs *BattleScene) onMouseDown(evname string, ev interface{}) {

}

func (bs *BattleScene) Render(renderer *renderer.Renderer) {
		// Renderiza o mundo 3D
		renderer.Render(bs.Node, bs.cam)

		// Renderiza a GUI 2D
		renderer.Render(bs.hud.root, bs.guiCam)
}

func (bs *BattleScene) Update(deltaTime float64) {
	// Atualiza a HUD com os dados atuais da batalha
	char1 := bs.battle.GetCharacter1()
	char2 := bs.battle.GetCharacter2()

	bs.hud.UpdateHealthBars(
		char1.Name, char1.Life, char1.MaxLife,
		char2.Name, char2.Life, char2.MaxLife,
	)
}

func (bs *BattleScene) OnResize(width, height int) {
		// Ajusta câmera 3D
		bs.cam.SetAspect(float32(width) / float32(height))

		// Ajusta câmera GUI (ortográfica)
		aspect := float32(width) / float32(height)
		bs.guiCam.SetAspect(aspect)

		// Ajusta HUD
		bs.hud.OnResize(float32(width), float32(height))
}

// GetHUD retorna a HUD da batalha
func (bs *BattleScene) GetHUD() *BattleHUD {
	return bs.hud
}

