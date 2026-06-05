package game

type Action struct {
	Id string
	Name string
	Type string
	Time int
	Effect func(actor *Character, target *Character)
}

func (ar *ActionRegistry) LoadActions() {
	ar.Register(&Action{
		Id:   "fireball",
		Name: "Fireball",
		Type: "magic",
		Time: 2,
		Effect: func(actor *Character, target *Character) {
			damage := 10
			target.Life -= damage
		},
	})

	ar.Register(&Action{
		Id:   "water_gun",
		Name: "Water Gun",
		Type: "magic",
		Time: 2,
		Effect: func(actor *Character, target *Character) {
			damage := 10
			target.Life -= damage
		},
	})
}


