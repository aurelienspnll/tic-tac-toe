package main

import (
	"fmt"
	"log"

	"github.com/aurelienspnll/tic-tac-toe/src/ai"
	"github.com/aurelienspnll/tic-tac-toe/src/game"
)

// func main() {
// 	fmt.Println("New game")
// 	var b = game.NewBoard(3, 3)
// 	var p, err = game.NewPlayer("Aurélien", "x")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		b.Mark(0, 0, "o")
// 		b.Mark(0, 1, "o")
// 		b.Mark(0, 2, "x")
// 		b.Mark(1, 0, "x")
// 		b.Mark(1, 1, "o")
// 		b.Mark(1, 2, "x")
// 		b.Mark(2, 0, "x")
// 		b.Mark(2, 1, "o")
// 		b.Mark(2, 2, "o")
// 		var s, err = b.GetMark(0, 0)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		s, err = b.ToString()
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Print(s)
// 		fmt.Println("IsFull :", b.IsFull())
// 		fmt.Println("IsWinByLineRightToLeft :", b.IsWinByLineRightToLeft())
// 		fmt.Println("IsWinByLineTopToBot :", b.IsWinByLineTopToBot())
// 		fmt.Println(p.GetScore())
// 	}
// }

func main() {
	var e = game.NewEngine(3, 3, "NOS", "ADEMO", true)
	//PLay VS AI
	//e.Play()
	var p *game.Player
	fmt.Println("----------------------")
	fmt.Println("Welcome to TicTacToe :")
	fmt.Println("----------------------")
	for !e.GetBoard().IsWin() && !e.GetBoard().IsFull() {
		s, err := e.GetBoard().ToString()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(s)
		p = e.WhoTurn()
		fmt.Println("It's", p.GetName(), "'s turn...")
		if p.IsAI() {
			var n = ai.NewNode(e.GetBoard(), 100, false)
			ai.BuildTree(n, 9, p.GetMark())
			//ai.PrintTree(n)
			posx, posy := ai.BestMove(n)
			err = e.GetBoard().Mark(posx, posy, p.GetMark())
			if err != nil {
				log.Fatal("AI failed")
			}
		} else {
			posx, posy := e.WherePlay()
			err = e.GetBoard().Mark(posx, posy, p.GetMark())
			for err != nil {
				fmt.Println("Wrong position !")
				posx, posy = e.WherePlay()
				err = e.GetBoard().Mark(posx, posy, p.GetMark())
			}
		}
		e.NextTurn()
	}
	if e.GetBoard().IsWin() {
		fmt.Println("----------------------")
		fmt.Println("    ", p.GetName(), " wins :")
		fmt.Println("----------------------")
	} else {
		fmt.Println("----------------------")
		fmt.Println("    No one wins :")
		fmt.Println("----------------------")
	}
	s, _ := e.GetBoard().ToString()
	fmt.Println(s)
}

// func main() {
// 	var b = game.NewBoard(3, 3)
// 	//var ai = ai.NewAI(b)
// 	//b.Mark(0, 0, "o")
// 	b.Mark(0, 1, "o")
// 	b.Mark(0, 2, "x")
// 	b.Mark(1, 0, "x")
// 	//b.Mark(1, 1, "x")
// 	//b.Mark(1, 2, "o")
// 	b.Mark(2, 0, "x")
// 	b.Mark(2, 1, "o")
// 	//b.Mark(2, 2, "x")
// 	s, _ := b.ToString()
// 	fmt.Println(s)
// 	var n = ai.NewNode(b, 100, false)
// 	//n.GenerateAllNextMove("o")
// 	//fmt.Println(n.ToString())
// 	ai.BuildTree(n, 9, "o")
// 	ai.PrintTree(n)
// 	ai.BestMove(n)
// 	//n.GenerateAllNextMove("o") // OK
// }
