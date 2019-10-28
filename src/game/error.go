package game

import (
    "fmt"
)

type PositionError struct {
	posx	int
 	posy	int
}

func (e *PositionError) Error() string {
    return fmt.Sprintf("Wrong position : %d - %d", e.posx, e.posy)
}

type MarkError struct {
	mark	string
}

func (e *MarkError) Error() string {
    return fmt.Sprintf("Wrong mark : %s, should be x or o", e.mark)
}
