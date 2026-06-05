package game

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
	return b
}


// Inicia a batalha, definindo quem começa e aguardando as ações dos personagens
func (b *Battle) Start() {
	// Lógica para iniciar a batalha entre character1 e character2
	fmt.Printf("A batalha começou entre %s e %s!\n", b.character1.Name, b.character2.Name)

	// Define quem começa com base na velocidade dos personagens
	b.choosing = b.firstCharacter(b.character1, b.character2)

	fmt.Printf("%s começa a batalha!\n", b.choosing.Name)
	fmt.Printf("Aguardando a ação de %s...\n", b.choosing.Name)
}


// GetCharacter1 retorna o primeiro personagem da batalha
func (b *Battle) GetCharacter1() *Character {
	return b.character1
}


// GetCharacter2 retorna o segundo personagem da batalha
func (b *Battle) GetCharacter2() *Character {
	return b.character2
}


func (b *Battle) GetActions(character *Character) []*Action {
	// Retorna as ações disponíveis para o personagem neste turno
	return character.Actions
}


func (b *Battle) DoAction(character *Character, action *Action, target *Character) {
	if character != b.choosing {
		fmt.Printf("Não é a vez de %s agir!\n", character.Name)
		return
	}

	fmt.Printf("%s executa a ação: %s\n", character.Name, action.Name)
	action.Effect(character, target)

	b.nextTurn()
}


// Define o personagem quem começa
func (b *Battle) firstCharacter(char1, char2 *Character) *Character {
	if char1.Speed > char2.Speed {
		return char1
	} else {
		return char2
	}
}


func (b *Battle) attack(attacker *Character, defender *Character) {
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
