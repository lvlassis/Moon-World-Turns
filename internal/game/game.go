package game

type Game struct {
	Character1 *Character
	Character2 *Character
	Stage      *Stage
}

func NewGame(Character1, Character2 *Character, Stage *Stage) *Game {
	return &Game{}
}

func (g *Game) Load() {
	
}
