package hud

import (
	"lvlassis/moon-world-turns/internal/game"

	"github.com/g3n/engine/gui"
)

func NewActionsPanel(character *game.Character, actions []*game.Action, battle *game.Battle) *gui.Panel {
	actionsPanel := gui.NewPanel(120, float32(len(actions)*50))
	actionsPanel.SetPosition(20, 80)
	
	for i, action := range actions {
		actionBtn := NewActionButton(action, character, battle) // Passa o personagem correto aqui
		actionBtn.SetPosition(0, float32(i*50))
		actionsPanel.Add(actionBtn)
	}

	return actionsPanel
}
