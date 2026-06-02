package engine

import "fmt"

type Action struct {
	Name string
}

func (actor Character) Execute(action Action, target *Character) {
	fmt.Printf("%s executa %s em %s\n", actor.Name, action.Name, target.Name)
}

func GetActions() []Action {
	ActionAttack := Action{Name: "Attack"}
	ActionDefend := Action{Name: "Defend"}
	ActionHeal := Action{Name: "Heal"}
	ActionMove := Action{Name: "Move"}
	return []Action{ActionAttack, ActionDefend, ActionHeal, ActionMove}
}
