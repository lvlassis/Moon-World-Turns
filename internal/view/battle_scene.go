package view

import (
	"lvlassis/moon-world-turns/internal/engine"
	"lvlassis/moon-world-turns/internal/events"
	"lvlassis/moon-world-turns/internal/view/hud"

	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
)

type BattleScene struct {
	*core.Node
	character1View *CharacterView
	character2View *CharacterView
	stage          *StageView
	battle         *engine.Battle
	bus            *events.Bus
	cam            *camera.Camera
	guiCam         *camera.Camera // Câmera ortográfica para GUI

	// GUI
	guiRoot    *gui.Panel
	healthBar1 *hud.HealthBar
	healthBar2 *hud.HealthBar

	actionButton  *gui.Button
	attackButton  *gui.Button
}

func NewBattleScene(char1, char2 *CharacterView, stage *StageView, battle *engine.Battle, bus *events.Bus, width, height int) Scene {
	scene := &BattleScene{
		Node:           core.NewNode(),
		character1View: char1,
		character2View: char2,
		battle:         battle,
		bus:            bus,
	}

	// Setup mundo 3D
	scene.setupWorld(stage)

	// Setup GUI
	scene.setupGUI(width, height)

	return scene
}

func (bs *BattleScene) setupWorld(stage *StageView) {
	// Posiciona os personagens e o estágio na cena
	bs.character1View.Sprite.SetPosition(-1, 0, 0)
	bs.Add(bs.character1View.Sprite)
	bs.character2View.Sprite.SetPosition(1, 0, 0)
	bs.Add(bs.character2View.Sprite)
	stage.SetPosition(0, -0.4, 0)
	bs.Add(stage)

	// Cria uma câmera para o mundo 3D
	bs.cam = camera.New(1)
	bs.cam.SetPosition(0, 1, 6)
	camera.NewOrbitControl(bs.cam)
	bs.Add(bs.cam)
}

func (bs *BattleScene) setupGUI(width, height int) {
	// Cria câmera ortográfica para GUI 2D
	bs.guiCam = camera.New(1)
	bs.guiCam.SetPositionVec(&math32.Vector3{0, 0, 1})
	bs.guiCam.UpdateMatrix()

	// Cria painel root para GUI
	bs.guiRoot = gui.NewPanel(float32(width), float32(height))
	bs.guiRoot.SetColor4(&math32.Color4{0, 0, 0, 0}) // Transparente

	character_1 := bs.battle.GetCharacter1()
	character_2 := bs.battle.GetCharacter2()

	// Cria barras de vida
	bs.healthBar1 = hud.NewHealthBar(20, 10, 200, character_1, bs.guiRoot)
	bs.healthBar2 = hud.NewHealthBar(float32(width)-220, 10, 200, character_2, bs.guiRoot)

	// // Cria botões de ataque
	bs.attackButton = hud.NewActionButton("Attack", character_1, engine.Action{Name: "Attack"})
	bs.guiRoot.Add(bs.attackButton)

	// Adiciona o root da GUI à cena
	bs.Add(bs.guiRoot)
}

func (bs *BattleScene) Render(renderer *renderer.Renderer) {
	// Renderiza o mundo 3D
	renderer.Render(bs.Node, bs.cam)

	// Renderiza a GUI 2D
	renderer.Render(bs.guiRoot, bs.guiCam)
}

func (bs *BattleScene) Update(deltaTime float64) {
	// Atualiza barras de vida com dados atuais da batalha
	char1 := bs.battle.GetCharacter1()
	char2 := bs.battle.GetCharacter2()

	bs.healthBar1.Update(char1.Name, char1.Life, char1.MaxLife)
	bs.healthBar2.Update(char2.Name, char2.Life, char2.MaxLife)
}

func (bs *BattleScene) OnResize(width, height int) {
	// Ajusta câmera 3D
	bs.cam.SetAspect(float32(width) / float32(height))

	// Ajusta câmera GUI
	bs.guiCam.SetAspect(float32(width) / float32(height))

	// Ajusta tamanho do painel GUI
	bs.guiRoot.SetSize(float32(width), float32(height))

	// Reposiciona barra 2 (direita)
	bs.healthBar2.Reposition(float32(width)-220, 10)
}

func (bs *BattleScene) GetNode() *core.Node {
	return bs.Node
}
