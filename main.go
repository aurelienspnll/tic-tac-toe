package main

import (
	"fmt"

	"github.com/aurelienspnll/tic-tac-toe/src/game"
)

func main()  {
	fmt.Println("New game")
	var b = game.NewBoard(3, 3)
	var p, err = game.NewPlayer("Aurélien", "x")
	if err != nil {
		fmt.Println(err)
	} else {
		b.Mark(0, 0, "o")
		b.Mark(0, 1, "o")
		b.Mark(0, 2, "x")
		b.Mark(1, 0, "x")
		b.Mark(1, 1, "o")
		b.Mark(1, 2, "x")
		b.Mark(2, 0, "x")
		b.Mark(2, 1, "o")
		b.Mark(2, 2, "o")
		var s, err = b.GetMark(0, 0)
		if err != nil {
			fmt.Println(err)
		}	
		s, err = b.ToString()
		if err != nil {
			fmt.Println(err)
		}		
		fmt.Print(s)
		fmt.Println("IsFull :", b.IsFull())
		fmt.Println("IsWinByLineRightToLeft :", b.IsWinByLineRightToLeft(0))
		fmt.Println("IsWinByLineTopToBot :", b.IsWinByLineTopToBot(0))
		fmt.Println(p.GetScore())
	}
}