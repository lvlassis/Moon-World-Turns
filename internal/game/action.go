package game

type Action struct {
	Id string
	Name string
	Type string
	Time int
	Effect func(actor *Character, target *Character)
}

