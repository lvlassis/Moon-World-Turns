package scenes

import (
	"lvlassis/moon-world-turns/internal/events"
	"lvlassis/moon-world-turns/internal/game"
	"lvlassis/moon-world-turns/internal/scenes/hud"

	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/gui"
)

type BattleScene struct {
	BaseScene

	character1View *CharacterView
	character2View *CharacterView
	stage          *StageView
	battle         *game.Battle

	// GUI
	healthBar1 *hud.HealthBar
	healthBar2 *hud.HealthBar
	actionsPanel  *gui.Panel
}

func NewBattleScene(char1, char2 *CharacterView, stage *StageView, battle *game.Battle, bus *events.Bus, width, height int) Scene {
	bs := &BattleScene{
		BaseScene:      *NewBaseScene(bus, width, height).(*BaseScene),
		character1View: char1,
		character2View: char2,
		battle:         battle,
	}

	// Setup mundo 3D
	bs.setupWorld(stage)

	// Setup Battle GUI
	bs.setupBattleGUI(width, height)

	return bs
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
	bs.Cam = camera.New(1)
	bs.Cam.SetPosition(0, 1, 6)
	camera.NewOrbitControl(bs.Cam)
	bs.Add(bs.Cam)
}

func (bs *BattleScene) setupBattleGUI(width, height int) {
	character_1 := bs.battle.GetCharacter1()
	character_2 := bs.battle.GetCharacter2()

	// Cria barras de vida
	bs.healthBar1 = hud.NewHealthBar(20, 10, 200, character_1, bs.GuiRoot)
	bs.healthBar2 = hud.NewHealthBar(float32(width)-220, 10, 200, character_2, bs.GuiRoot)

	bs.actionsPanel = hud.NewActionsPanel(character_1, character_1.Actions, bs.battle)
	bs.GuiRoot.Add(bs.actionsPanel)

}

func (bs *BattleScene) Update(deltaTime float64) {
	// Atualiza barras de vida com dados atuais da batalha
	char1 := bs.battle.GetCharacter1()
	char2 := bs.battle.GetCharacter2()

	bs.healthBar1.Update(char1.Name, char1.Life, char1.MaxLife)
	bs.healthBar2.Update(char2.Name, char2.Life, char2.MaxLife)
}

func (bs *BattleScene) OnResize(width, height int) {
	bs.BaseScene.OnResize(width, height)

	// Reposiciona barra 2 (direita)
	bs.healthBar2.Reposition(float32(width)-220, 10)
}

