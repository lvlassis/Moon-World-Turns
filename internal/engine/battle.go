package engine

import (
	"fmt"
	"lvlassis/moon-world-turns/internal/events"
)

type Battle struct {
	character1 	*Character
	character2 	*Character
	choosing	*Character
	turn		int
}

func NewBattle(character1, character2 *Character, bus *events.Bus) *Battle {
	b := &Battle{
		character1: character1,
		character2: character2,
	}
	bus.Subscribe(events.CharacterAction, b.OnAction)
	return b
}

func (b *Battle) Start() {
	// Lógica para iniciar a batalha entre character1 e character2
	fmt.Printf("A batalha começou entre %s e %s!\n", b.character1.Name, b.character2.Name)

	// Define quem começa com base na velocidade dos personagens
	fastest := b.character1.Speed > b.character2.Speed
	if fastest {
		b.choosing = b.character1
	} else {
		b.choosing = b.character2
	}

	fmt.Printf("%s começa a batalha!\n", b.choosing.Name)
	fmt.Printf("Aguardando a ação de %s...\n", b.choosing.Name)
}

func (b *Battle) OnAction(eventName string, eventData interface{}) {
	data, ok := eventData.(events.CharacterActionData)
	if !ok {
		return
	}

	// Verifica se o personagem que realizou a ação é o que está escolhendo
	if data.CharacterID != b.choosing.ID {
		return
	}

	// Mostra a ação realizada
	fmt.Printf("Character %s fez uma ação!\n", data.CharacterID)

	// Passa o turno
	b.turn++
	if b.choosing == b.character1 {
		b.choosing = b.character2
	} else {
		b.choosing = b.character1
	}
}

// GetCharacter1 retorna o primeiro personagem da batalha
func (b *Battle) GetCharacter1() *Character {
	return b.character1
}

// GetCharacter2 retorna o segundo personagem da batalha
func (b *Battle) GetCharacter2() *Character {
	return b.character2
}

func (b *Battle) Attack(attacker *Character, defender *Character) {
	if attacker != b.choosing {
		fmt.Printf("Não é a vez de %s atacar!\n", attacker.Name)
		return
	}
	attacker.Attack(defender)
	b.nextTurn()
}

// nextTurn passa o turno para o próximo personagem
func (b *Battle) nextTurn() {
	b.turn++

	if b.choosing == b.character1 {
		b.choosing = b.character2
	} else {
		b.choosing = b.character1
	}

	fmt.Printf("Turno %d: É a vez de %s\n", b.turn, b.choosing.Name)
}
