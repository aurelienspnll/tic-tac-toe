package game

import (
	"errors"
	"fmt"
	"log"
)

type Board struct {
	square            []string
	size              int
	numberToWin       int
	extraLineWinnable int
}

func NewBoard(size int, numberToWin int) *Board {
	b := new(Board)
	b.size = size
	b.numberToWin = numberToWin
	b.extraLineWinnable = b.size % b.numberToWin
	b.square = make([]string, size*size)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			b.square[x*size+y] = "-"
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
		b.square[posx*b.size+posy] = mark
	}
	return nil
}

func (b *Board) GetMark(posx int, posy int) (string, error) {
	if posx > b.size || posy > b.size {
		return "error", &PositionError{posx, posy}
	} else {
		return b.square[posx*b.size+posy], nil
	}
}

// Must be call in game engine only after 5 turns
func (b *Board) IsWin() bool {
	//extraLineWinnable := b.size % b.numberToWin
	return false
}

//Split by 4 directions to be able to implement the algo : see paper
// TODO : Algo to implemented : see on my paper

func (b *Board) IsWinByLineRightToLeft() bool {
	for x := 0; x < b.size; x++ {
		for y := 0; y <= b.extraLineWinnable; y++ {
			var m, err = b.GetMark(x, y)
			if err != nil {
				log.Fatal(err)
			}
			if m == "x" || m == "o" { //avoid "-" case
				for i := 1; i < b.numberToWin; i++ {
					var mTmp, err = b.GetMark(x, y+i)
					fmt.Println(mTmp)
					if err != nil {
						log.Fatal(err)
					}
					if m == mTmp && i == b.numberToWin-1 {
						return true
					} else if m == mTmp {
						m = mTmp
					} else {
						break
					}
				}
			}
		}
	}
	return false
}

func (b *Board) IsWinByLineTopToBot() bool {
	for y := 0; y < b.size; y++ {
		for x := 0; x <= b.extraLineWinnable; x++ {
			var m, err = b.GetMark(x, y)
			if err != nil {
				log.Fatal(err)
			}
			if m == "x" || m == "o" { //avoid "-" case
				for i := 1; i < b.numberToWin; i++ {
					var mTmp, err = b.GetMark(x+i, y)
					fmt.Println(mTmp)
					if err != nil {
						log.Fatal(err)
					}
					if m == mTmp && i == b.numberToWin-1 {
						return true
					} else if m == mTmp {
						m = mTmp
					} else {
						break
					}
				}
			}
		}
	}
	return false
}

func (b *Board) IsWinByDiagonal() bool {
	//numberOfPossibility := b.size + b.size * b.extraLineWinnable // same for diagonals?
	return false
}

func (b *Board) IsFull() bool {
	for x := 0; x < b.size*b.size; x++ {
		if b.square[x] == "-" {
			return false
		}
	}
	return true
}

func (b *Board) Fill(square []string) error {
	if len(square) != b.size*b.size {
		return errors.New("Your list have not the good size")
	}
	copy(b.square, square)
	return nil
}

func (b *Board) ToString() (string, error) {
	var res string
	for x := 0; x < b.size; x++ {
		for y := 0; y < b.size; y++ {
			m, err := b.GetMark(x, y)
			if err != nil {
				return "error", err
			} else if y == 0 {
				res = res + "| " + m + " | "
			} else {
				res = res + m + " | "
			}
		}
		res = res + "\n"
	}
	return res, nil
}
