package game

type Player struct {
	name  string
	mark  string
	score int
	AI    bool
}

func NewPlayer(name string, mark string, AI bool) (*Player, error) {
	p := new(Player)
	if mark != "x" && mark != "o" {
		return nil, &MarkError{mark}
	} else {
		p.name = name
		p.mark = mark
		p.score = 0
		p.AI = AI
	}
	return p, nil
}

func (p *Player) GetScore() int {
	return p.score
}

func (p *Player) GetMark() string {
	return p.mark
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) IsAI() bool {
	return p.AI
}
