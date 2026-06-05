package hud

import (
	"fmt"
	"lvlassis/moon-world-turns/internal/game"

	"github.com/g3n/engine/gui"
)

func NewActionButton(action *game.Action, character *game.Character, battle *game.Battle) *gui.Button {
	label := action.Name
	action_btn := gui.NewButton(label)
	action_btn.SetPosition(20, 80)
	action_btn.SetSize(100, 40)
	action_btn.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		fmt.Printf("Botão '%s' clicado! %s executa ação: %s\n", label, character.Name, action.Name)
		char2 := battle.GetCharacter2()
		battle.DoAction(character, action, char2)
	})
	return action_btn
}

