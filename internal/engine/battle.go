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

// PlayerAttack executa um ataque de um personagem no outro
// attackerNum: 1 = character1 ataca, 2 = character2 ataca
func (b *Battle) PlayerAttack(attackerNum int) {
	var attacker, defender *Character

	// Define atacante e defensor
	if attackerNum == 1 {
		attacker = b.character1
		defender = b.character2
	} else {
		attacker = b.character2
		defender = b.character1
	}

	// Verifica se é o turno do atacante
	if b.choosing != attacker {
		fmt.Printf("Não é o turno de %s!\n", attacker.Name)
		return
	}

	// Calcula dano (por enquanto usa força do atacante)
	damage := attacker.Strength

	// Aplica dano
	defender.Life -= damage
	if defender.Life < 0 {
		defender.Life = 0
	}

	fmt.Printf("%s atacou %s causando %d de dano! (%d HP restante)\n",
		attacker.Name, defender.Name, damage, defender.Life)

	// Verifica se a batalha acabou
	if defender.Life <= 0 {
		fmt.Printf("%s venceu a batalha!\n", attacker.Name)
		return
	}

	// Passa o turno
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
