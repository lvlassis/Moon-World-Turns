package view

import (
	"fmt"
	"lvlassis/moon-world-turns/internal/engine"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
)

type BattleHUD struct {
	root *gui.Panel

	// Barras de vida
	char1HealthBg   *gui.Panel  // Fundo da barra (cinza)
	char1HealthBar  *gui.Panel  // Barra de vida (verde)
	char1NameLabel  *gui.Label  // Nome do personagem 1

	char2HealthBg   *gui.Panel  // Fundo da barra (cinza)
	char2HealthBar  *gui.Panel  // Barra de vida (verde)
	char2NameLabel  *gui.Label  // Nome do personagem 2

	// Botões de ação
	char1AttackButton *gui.Button
	char2AttackButton *gui.Button

	maxBarWidth float32

	// Referência para a Battle
	battle *engine.Battle
}

func NewBattleHUD(width, height int, battle *engine.Battle) *BattleHUD {
	hud := &BattleHUD{
		root:        gui.NewPanel(float32(width), float32(height)),
		maxBarWidth: 200,
		battle:      battle,
	}

	// Configura estilo do root (transparente para não cobrir o mundo 3D)
	hud.root.SetColor4(&math32.Color4{0, 0, 0, 0})

	// === Personagem 1 (esquerda superior) ===

	// Nome do personagem 1
	hud.char1NameLabel = gui.NewLabel("Character 1")
	hud.char1NameLabel.SetPosition(20, 10)
	hud.char1NameLabel.SetColor(math32.NewColor("white"))
	hud.char1NameLabel.SetFontSize(16)
	hud.root.Add(hud.char1NameLabel)

	// Fundo da barra de vida 1 (cinza)
	hud.char1HealthBg = gui.NewPanel(hud.maxBarWidth, 20)
	hud.char1HealthBg.SetPosition(20, 35)
	hud.char1HealthBg.SetColor(math32.NewColor("darkgray"))
	hud.char1HealthBg.SetBorders(2, 2, 2, 2)
	hud.char1HealthBg.SetBordersColor(math32.NewColor("white"))
	hud.root.Add(hud.char1HealthBg)

	// Barra de vida 1 (verde, vai encolher conforme dano)
	hud.char1HealthBar = gui.NewPanel(hud.maxBarWidth, 20)
	hud.char1HealthBar.SetPosition(20, 35)
	hud.char1HealthBar.SetColor(math32.NewColor("green"))
	hud.root.Add(hud.char1HealthBar)

	// Botão de ataque Character 1
	hud.char1AttackButton = gui.NewButton("Attack")
	hud.char1AttackButton.SetPosition(20, 60)
	hud.char1AttackButton.SetSize(100, 40)  // Define tamanho explícito
	hud.char1AttackButton.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		fmt.Println("=== CHARACTER 1 ATTACK BUTTON CLICKED! ===")
		hud.onChar1AttackClick()
	})
	hud.char1AttackButton.Subscribe(gui.OnMouseDown, func(name string, ev interface{}) {
		fmt.Println("=== Mouse DOWN no botão 1 ===")
	})
	hud.root.Add(hud.char1AttackButton)

	// === Personagem 2 (direita superior) ===

	char2X := float32(width) - hud.maxBarWidth - 20

	// Nome do personagem 2
	hud.char2NameLabel = gui.NewLabel("Character 2")
	hud.char2NameLabel.SetPosition(char2X, 10)
	hud.char2NameLabel.SetColor(math32.NewColor("white"))
	hud.char2NameLabel.SetFontSize(16)
	hud.root.Add(hud.char2NameLabel)

	// Fundo da barra de vida 2 (cinza)
	hud.char2HealthBg = gui.NewPanel(hud.maxBarWidth, 20)
	hud.char2HealthBg.SetPosition(char2X, 35)
	hud.char2HealthBg.SetColor(math32.NewColor("darkgray"))
	hud.char2HealthBg.SetBorders(2, 2, 2, 2)
	hud.char2HealthBg.SetBordersColor(math32.NewColor("white"))
	hud.root.Add(hud.char2HealthBg)

	// Barra de vida 2 (verde)
	hud.char2HealthBar = gui.NewPanel(hud.maxBarWidth, 20)
	hud.char2HealthBar.SetPosition(char2X, 35)
	hud.char2HealthBar.SetColor(math32.NewColor("green"))
	hud.root.Add(hud.char2HealthBar)

	// Botão de ataque Character 2
	hud.char2AttackButton = gui.NewButton("Attack")
	hud.char2AttackButton.SetPosition(char2X, 60)
	hud.char2AttackButton.SetSize(100, 40)  // Define tamanho explícito
	hud.char2AttackButton.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		fmt.Println("=== CHARACTER 2 ATTACK BUTTON CLICKED! ===")
		hud.onChar2AttackClick()
	})
	hud.char2AttackButton.Subscribe(gui.OnMouseDown, func(name string, ev interface{}) {
		fmt.Println("=== Mouse DOWN no botão 2 ===")
	})
	hud.root.Add(hud.char2AttackButton)

	return hud
}

// HandleClick processa um clique nas coordenadas x, y
// Retorna true se algum botão foi clicado
func (h *BattleHUD) HandleClick(x, y float32) bool {
	fmt.Printf("HandleClick: x=%.2f, y=%.2f\n", x, y)

	// Verifica botão 1
	btn1X := h.char1AttackButton.Position().X
	btn1Y := h.char1AttackButton.Position().Y
	btn1W := h.char1AttackButton.Width()
	btn1H := h.char1AttackButton.Height()

	fmt.Printf("Botão 1: x=%.2f, y=%.2f, w=%.2f, h=%.2f\n", btn1X, btn1Y, btn1W, btn1H)

	if x >= btn1X && x <= btn1X+btn1W && y >= btn1Y && y <= btn1Y+btn1H {
		fmt.Println("=== CLIQUE NO BOTÃO 1 DETECTADO! ===")
		h.onChar1AttackClick()
		return true
	}

	// Verifica botão 2
	btn2X := h.char2AttackButton.Position().X
	btn2Y := h.char2AttackButton.Position().Y
	btn2W := h.char2AttackButton.Width()
	btn2H := h.char2AttackButton.Height()

	fmt.Printf("Botão 2: x=%.2f, y=%.2f, w=%.2f, h=%.2f\n", btn2X, btn2Y, btn2W, btn2H)

	if x >= btn2X && x <= btn2X+btn2W && y >= btn2Y && y <= btn2Y+btn2H {
		fmt.Println("=== CLIQUE NO BOTÃO 2 DETECTADO! ===")
		h.onChar2AttackClick()
		return true
	}

	return false
}

// onChar1AttackClick é chamado quando o botão de ataque do char1 é clicado
func (h *BattleHUD) onChar1AttackClick() {
	fmt.Println("Character 1 Attack button clicked!")

	// Chama método de ataque da Battle
	h.battle.PlayerAttack(1) // 1 = character1 ataca
}

// onChar2AttackClick é chamado quando o botão de ataque do char2 é clicado
func (h *BattleHUD) onChar2AttackClick() {
	fmt.Println("Character 2 Attack button clicked!")

	// Chama método de ataque da Battle
	h.battle.PlayerAttack(2) // 2 = character2 ataca
}

// UpdateHealthBars atualiza as barras de vida com valores de 0 a 100
func (h *BattleHUD) UpdateHealthBars(char1Name string, char1HP, char1MaxHP int, char2Name string, char2HP, char2MaxHP int) {
	// Calcula porcentagem
	char1Percent := float32(char1HP) / float32(char1MaxHP)
	char2Percent := float32(char2HP) / float32(char2MaxHP)

	// Atualiza largura das barras
	h.char1HealthBar.SetWidth(h.maxBarWidth * char1Percent)
	h.char2HealthBar.SetWidth(h.maxBarWidth * char2Percent)

	// Atualiza cor (verde -> amarelo -> vermelho)
	h.char1HealthBar.SetColor(h.getHealthColor(char1Percent))
	h.char2HealthBar.SetColor(h.getHealthColor(char2Percent))

	// Atualiza nomes
	h.char1NameLabel.SetText(char1Name)
	h.char2NameLabel.SetText(char2Name)
}

// getHealthColor retorna cor baseada na porcentagem de vida
func (h *BattleHUD) getHealthColor(percent float32) *math32.Color {
	if percent > 0.5 {
		return math32.NewColor("green")
	} else if percent > 0.25 {
		return math32.NewColor("yellow")
	} else {
		return math32.NewColor("red")
	}
}

// OnResize reposiciona elementos quando a janela é redimensionada
func (h *BattleHUD) OnResize(width, height float32) {
	h.root.SetSize(width, height)

	// Reposiciona barra da direita
	char2X := width - h.maxBarWidth - 20
	h.char2NameLabel.SetPosition(char2X, 10)
	h.char2HealthBg.SetPosition(char2X, 35)
	h.char2HealthBar.SetPosition(char2X, 35)
}

// GetRoot retorna o root panel para ser renderizado
func (h *BattleHUD) GetRoot() *gui.Panel {
	return h.root
}
