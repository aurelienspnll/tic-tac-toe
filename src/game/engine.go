package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Engine struct {
	board      *Board
	player_one *Player
	player_two *Player
	turn       int
}

func NewEngine(size int, numberToWin int, name_one string, name_two string) *Engine {
	var err error
	e := new(Engine)
	e.board = NewBoard(size, numberToWin)
	e.player_one, err = NewPlayer(name_one, "x")
	e.turn = 0
	if err != nil {
		log.Fatal(err)
	}
	e.player_two, err = NewPlayer(name_two, "o")
	if err != nil {
		log.Fatal(err)
	}
	return e
}

func (e *Engine) WhoTurn() *Player {
	// TODO : in one line ?
	if e.turn == 0 {
		return e.player_one
	} else {
		return e.player_two
	}
}

func (e *Engine) NextTurn() {
	if e.turn == 0 {
		e.turn = 1
	} else {
		e.turn = 0
	}
}

func (e *Engine) WherePlay() (int, int) {
	//TODO : Check if the player type integer on input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Where do you want to play (type : x y) ? ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	text = strings.ReplaceAll(text, "\n", "")
	var line []string
	line = strings.Split(text, " ")
	posx, err := strconv.Atoi(line[0])
	if err != nil {
		log.Fatal(err)
	}
	posy, err := strconv.Atoi(line[1])
	if err != nil {
		log.Fatal(err)
	}
	return posx, posy
}

func (e *Engine) Play() {
	//STDIN to play turn by turn
	var p *Player
	fmt.Println("----------------------")
	fmt.Println("Welcome to TicTacToe :")
	fmt.Println("----------------------")
	for !e.board.IsWin() && !e.board.IsFull() {
		s, err := e.board.ToString()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(s)
		p = e.WhoTurn()
		fmt.Println("It's", p.GetName(), "'s turn...")
		posx, posy := e.WherePlay()
		err = e.board.Mark(posx, posy, p.GetMark())
		for err != nil {
			fmt.Println("Wrong position !")
			posx, posy = e.WherePlay()
			err = e.board.Mark(posx, posy, p.GetMark())
		}
		e.NextTurn()
	}
	if e.board.IsWin() {
		fmt.Println("----------------------")
		fmt.Println("    ", p.GetName(), " wins :")
		fmt.Println("----------------------")
	} else {
		fmt.Println("----------------------")
		fmt.Println("    No one wins :")
		fmt.Println("----------------------")
	}
	s, _ := e.board.ToString()
	fmt.Println(s)
}
