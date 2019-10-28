package game

type Player struct {
	name	string
	mark 	string
	score	int
}

func NewPlayer(name string, mark string) (*Player, error) {
	p := new(Player)
	if(mark != "x" && mark != "o"){
		return nil, &MarkError{mark}
	} else {
		p.name = name
		p.mark = mark
		p.score = 0
	}
	return p, nil
}	

func (p *Player) GetScore() int {
	return p.score
}