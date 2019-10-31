package game

import (
	"errors"
	"log"
	"strconv"
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
	if posx >= b.size || posy >= b.size {
		return "error", &PositionError{posx, posy}
	} else {
		return b.square[posx*b.size+posy], nil
	}
}

// Must be call in game engine only after 5 turns
func (b *Board) IsWin() bool {
	return b.IsWinByLineRightToLeft() || b.IsWinByLineTopToBot() || b.IsWinByDiagonalRightToLeft() || b.IsWinByDiagonalLeftToRight() // || b.IsFull()
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
					//fmt.Println(mTmp)
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

func (b *Board) IsWinByDiagonalRightToLeft() bool {
	for x := 0; x <= b.extraLineWinnable; x++ {
		for y := 0; y <= b.extraLineWinnable; y++ {
			var m, err = b.GetMark(x, y)
			if err != nil {
				log.Fatal(err)
			}
			if m == "x" || m == "o" { //avoid "-" case
				for i := 1; i < b.numberToWin; i++ {
					var mTmp, err = b.GetMark(x+i, y+i)
					//fmt.Println(mTmp)
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

func (b *Board) IsWinByDiagonalLeftToRight() bool {
	bound := b.size - b.extraLineWinnable - 1 //-1 because array start to 0
	for x := 0; x <= b.extraLineWinnable; x++ {
		for y := (b.size - 1); y >= bound; y-- {
			var m, err = b.GetMark(x, y)
			if err != nil {
				log.Fatal(err)
			}
			if m == "x" || m == "o" { //avoid "-" case
				for i := 1; i < b.numberToWin; i++ {
					var mTmp, err = b.GetMark(x+i, y-i)
					//fmt.Println(mTmp)
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

func (b *Board) RemainingMoves() int {
	var res = 0
	for x := 0; x < b.size*b.size; x++ {
		if b.square[x] == "-" {
			res += 1
		}
	}
	return res
}

func (b *Board) GetBlankPos(i int) (int, error) {
	var count = 0
	for x := 0; x < b.size*b.size; x++ {
		if b.square[x] == "-" && count == i {
			return x, nil
		} else if b.square[x] == "-" {
			count += 1
		}
	}
	return 0, errors.New("Bad position of blank")
}

func (b *Board) IsFull() bool {
	for x := 0; x < b.size*b.size; x++ {
		if b.square[x] == "-" {
			return false
		}
	}
	return true
}

func (b *Board) IsEmpty() bool {
	for x := 0; x < b.size*b.size; x++ {
		if b.square[x] == "x" || b.square[x] == "o" {
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

func (b *Board) LineToString(line int) (string, error) {
	if line >= b.size {
		return "error", errors.New("Wrong line")
	}
	var res string
	for y := 0; y < b.size; y++ {
		m, err := b.GetMark(line, y)
		if err != nil {
			return "error", err
		} else if y == 0 {
			res = "| " + strconv.Itoa(line) + " | " + m + " | "
		} else {
			res = res + m + " | "
		}
	}
	res = res + "\n"
	return res, nil
}

func (b *Board) LinePosToString() string {
	var res string
	for y := 0; y < b.size; y++ {
		if y == 0 {
			res = "|x\\y| " + strconv.Itoa(y) + " | "
		} else {
			res = res + strconv.Itoa(y) + " | "
		}
	}
	res = res + "\n"
	return res
}

func (b *Board) ToString() (string, error) {
	var res string
	for x := 0; x < b.size; x++ {
		if x == 0 {
			res += b.LinePosToString()
		}
		tmp, err := b.LineToString(x)
		if err != nil {
			return tmp, err
		}
		res += tmp
	}
	return res, nil
}

func (b *Board) GetSquare() []string {
	return b.square
}

func (b *Board) GetSize() int {
	return b.size
}

func (b *Board) GetNumberToWin() int {
	return b.numberToWin
}
