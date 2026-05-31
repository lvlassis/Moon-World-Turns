package view

import (
	"fmt"
	"lvlassis/moon-world-turns/internal/engine"
	"lvlassis/moon-world-turns/internal/events"

	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
)

// HealthBar representa uma barra de vida parametrizada
type HealthBar struct {
	background   *gui.Panel  // Fundo cinza
	bar          *gui.Panel  // Barra verde/amarela/vermelha
	nameLabel    *gui.Label  // Nome do personagem
	attackButton *gui.Button // Botão de ataque
	maxWidth     float32     // Largura máxima da barra
}

// newHealthBar cria uma nova barra de vida
func newHealthBar(x, y, maxWidth float32, character *engine.Character, parent *gui.Panel) *HealthBar {
	hb := &HealthBar{
		maxWidth: maxWidth,
	}

	// Nome do personagem
	hb.nameLabel = gui.NewLabel(character.Name)
	hb.nameLabel.SetPosition(x, y)
	hb.nameLabel.SetColor(math32.NewColor("white"))
	hb.nameLabel.SetFontSize(16)
	parent.Add(hb.nameLabel)

	// Fundo da barra (cinza)
	hb.background = gui.NewPanel(maxWidth, 20)
	hb.background.SetPosition(x, y+25)
	hb.background.SetColor(math32.NewColor("darkgray"))
	hb.background.SetBorders(2, 2, 2, 2)
	hb.background.SetBordersColor(math32.NewColor("white"))
	parent.Add(hb.background)

	// Barra de vida (verde)
	hb.bar = gui.NewPanel(maxWidth, 20)
	hb.bar.SetPosition(x, y+25)
	hb.bar.SetColor(math32.NewColor("green"))
	parent.Add(hb.bar)

	return hb
}

// update atualiza a barra com novos valores
func (hb *HealthBar) update(name string, hp, maxHP int) {
	// Atualiza nome
	hb.nameLabel.SetText(name)

	// Calcula porcentagem
	percent := float32(hp) / float32(maxHP)

	// Atualiza largura da barra
	hb.bar.SetWidth(hb.maxWidth * percent)

	// Atualiza cor (verde -> amarelo -> vermelho)
	if percent > 0.5 {
		hb.bar.SetColor(math32.NewColor("green"))
	} else if percent > 0.25 {
		hb.bar.SetColor(math32.NewColor("yellow"))
	} else {
		hb.bar.SetColor(math32.NewColor("red"))
	}
}

// reposition reposiciona todos os elementos da barra
func (hb *HealthBar) reposition(x, y float32) {
	hb.nameLabel.SetPosition(x, y)
	hb.background.SetPosition(x, y+25)
	hb.bar.SetPosition(x, y+25)
	// attackButton foi removido do HealthBar
}

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
	healthBar1 *HealthBar
	healthBar2 *HealthBar
	atackBtn1	 *gui.Button
	atackBtn2	 *gui.Button
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
	bs.healthBar1 = newHealthBar(20, 10, 200, character_1, bs.guiRoot)
	bs.healthBar2 = newHealthBar(float32(width)-220, 10, 200, character_2, bs.guiRoot)

	// Cria botões de ataque
	bs.atackBtn1 = gui.NewButton("Attack")
	bs.atackBtn1.SetPosition(20, 80)
	bs.atackBtn1.SetSize(100, 40)
	bs.atackBtn1.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		fmt.Printf("=== BOTÃO DE ATAQUE '%s' CLICADO! ===\n", character_1.Name)
		bs.battle.Attack(character_1, character_2)
	})
	bs.guiRoot.Add(bs.atackBtn1)  // ← Corrigido: parent.Add(child)

	// Cria botões de ataque
	bs.atackBtn2 = gui.NewButton("Attack")
	bs.atackBtn2.SetPosition(200, 80)
	bs.atackBtn2.SetSize(100, 40)
	bs.atackBtn2.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		fmt.Printf("=== BOTÃO DE ATAQUE '%s' CLICADO! ===\n", character_2.Name)
		bs.battle.Attack(character_2, character_1)
	})
	bs.guiRoot.Add(bs.atackBtn2)  // ← Corrigido: parent.Add(child)

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

	bs.healthBar1.update(char1.Name, char1.Life, char1.MaxLife)
	bs.healthBar2.update(char2.Name, char2.Life, char2.MaxLife)
}

func (bs *BattleScene) OnResize(width, height int) {
	// Ajusta câmera 3D
	bs.cam.SetAspect(float32(width) / float32(height))

	// Ajusta câmera GUI
	bs.guiCam.SetAspect(float32(width) / float32(height))

	// Ajusta tamanho do painel GUI
	bs.guiRoot.SetSize(float32(width), float32(height))

	// Reposiciona barra 2 (direita)
	bs.healthBar2.reposition(float32(width)-220, 10)
}

func (bs *BattleScene) GetNode() *core.Node {
	return bs.Node
}
