package hud

import (
	"fmt"
	"lvlassis/moon-world-turns/internal/engine"

	"github.com/g3n/engine/gui"
)

func NewActionButton(label string, character *engine.Character, action engine.Action,) *gui.Button {
	action_btn := gui.NewButton(label)
	action_btn.SetPosition(20, 80)
	action_btn.SetSize(100, 40)
	action_btn.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		fmt.Printf("Botão '%s' clicado! %s executa ação: %s\n", label, character.Name, action.Name)
	})
	return action_btn
}

