package game

import "lvlassis/moon-world-turns/events"

type Battle struct {
	Character1 *Character 
	Character2 *Character 
	Stage 	*Stage
	Turno int
	Vez *Character

	broker *events.Broker
}

func NewBattle(char1, char2 *Character, stage *Stage, broker *events.Broker) *Battle {
	// Estado Inicial
	battle := &Battle{
		Character1: char1,
		Character2: char2,
		Stage: stage,
		Turno: 0,
		Vez: char1, 
		broker: broker,
	}

	broker.Subscribe(events.CharacterAction, battle.OnCharacterAction)

	return battle
}


func (battle *Battle) OnCharacterAction(evname string, ev interface{}) {
	if ev.(*events.CharacterActionData).CharacterID != battle.Vez.ID {
		return
	}
}
