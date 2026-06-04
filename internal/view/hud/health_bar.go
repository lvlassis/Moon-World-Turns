package hud

import (
	"lvlassis/moon-world-turns/internal/game"

	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
)

// HealthBar representa uma barra de vida parametrizada
type HealthBar struct {
	background   *gui.Panel  // Fundo cinza
	bar          *gui.Panel  // Barra verde/amarela/vermelha
	nameLabel    *gui.Label  // Nome do personagem
	maxWidth     float32     // Largura máxima da barra
}

// newHealthBar cria uma nova barra de vida
func NewHealthBar(x, y, maxWidth float32, character *game.Character, parent *gui.Panel) *HealthBar {
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
func (hb *HealthBar) Update(name string, hp, maxHP int) {
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
func (hb *HealthBar) Reposition(x, y float32) {
	hb.nameLabel.SetPosition(x, y)
	hb.background.SetPosition(x, y+25)
	hb.bar.SetPosition(x, y+25)
}

