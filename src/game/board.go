package game

import "log"
import "fmt"

const size = 3
const numberToWin = 3 

type Board struct {
	square	[]string
}

func NewBoard() *Board {
	b := new(Board)
	b.square = make([]string, size * size)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			b.square[x * size + y] = " "
		}
	}
	return b
}

func (b *Board) Mark(posx int, posy int, mark string) error {
	m, err := b.GetMark(posx, posy)
	if err != nil {
		return err
	} else if m == "x" || m == "o" {
		return &PositionError{posx, posy} // Square already played
	} else {
		b.square[posx * size + posy] = mark
	}
	return nil
}

func (b *Board) GetMark(posx int, posy int) (string, error) {
	if posx > size || posy > size {
		return "error", &PositionError{posx, posy}
	} else {
		return b.square[posx * size + posy], nil
	}
}

// Must be call in game engine only after 5 turns
func (b *Board) IsWin() bool {
	//extraLineWinnable := size % numberToWin
	return false
}

//Split by 4 directions to be able to implement the algo : see paper
// TODO : Algo to implemented : see on my paper 

func (b *Board) IsWinByLineRightToLeft(extraLineWinnable int) bool {
	for x := 0; x < size; x++ {
		for y := 0; y <= extraLineWinnable; y++ {
			var m, err = b.GetMark(x, y)
			if err != nil {
				log.Fatal(err)
			}
			for i := 1; i < numberToWin; i++ {
				var mTmp, err = b.GetMark(x, y + i)
				fmt.Println(mTmp)
				if err != nil {
					log.Fatal(err)
				}
				if m == mTmp && i == numberToWin - 1 {
					return true
				} else if m == mTmp {
					m = mTmp
				} else { 
					break
				}
			}
		}
	}
	return false
}

func (b *Board) IsWinByLineTopToBot(extraLineWinnable int) bool {
	for y := 0; y < size; y++ {
		for x := 0; x <= extraLineWinnable; x++ {
			var m, err = b.GetMark(x, y)
			if err != nil {
				log.Fatal(err)
			}
			for i := 1; i < numberToWin; i++ {
				var mTmp, err = b.GetMark(x + i, y)
				fmt.Println(mTmp)
				if err != nil {
					log.Fatal(err)
				}
				if m == mTmp && i == numberToWin - 1 {
					return true
				} else if m == mTmp {
					m = mTmp
				} else { 
					break
				}
			}
		}
	}
	return false
}


func (b *Board) IsWinByDiagonal(extraLineWinnable int) bool {
	//numberOfPossibility := size + size * extraLineWinnable // same for diagonals?
	return false
}

func (b *Board) IsFull() bool {
	for x := 0; x < size * size ; x++ {
		if b.square[x] == " " {
			return false
		}
	}
	return true
}

func (b *Board) ToString() (string, error) {
	var res string
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			m, err := b.GetMark(x, y)
			if err != nil {
				return "error", err
			} else if y == 0{
				res = res + "| " + m + " | "
			} else {
				res = res + m + " | "
			}
		}
		res = res + "\n"
	}
	return res, nil
}

