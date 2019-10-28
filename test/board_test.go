package test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/aurelienspnll/tic-tac-toe/src/game"
	"github.com/stretchr/testify/assert"
)

var b *game.Board

func TestMain(m *testing.M) {
	b = game.NewBoard(3, 3)
	code := m.Run()
	os.Exit(code)
}

func TestWinByLaneRightToLeft_1(t *testing.T) {
	var s string
	var err error
	b = game.NewBoard(3, 3)
	boardSquare := []string{"x", "x", "x", " ", " ", " ", " ", " ", " "}
	err = b.Fill(boardSquare)
	if err != nil {
		fmt.Println(err)
	}
	// INIT OK

	s, err = b.ToString()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(s)
	}

	assert.True(t, b.IsWinByLineRightToLeft(), "Should be true")

}

func TestWinByLaneRightToLeft_2(t *testing.T) {
	var numberOfWin int = 0
	file, err := os.Open("./combination/all_combinations.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []string
		line = strings.Split(scanner.Text(), "")

		b = game.NewBoard(3, 3)
		err = b.Fill(line)
		if err != nil {
			fmt.Println(err)
		}
		// INIT OK

		if b.IsWinByLineRightToLeft() {
			numberOfWin += 1
		}
		// s, err = b.ToString()
		// if err != nil {
		// 	fmt.Println(err)
		// } else {
		// 	fmt.Println(s)
		// }
	}
	fmt.Println("Number of win : ", numberOfWin)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func TestWinByLaneTopToBot_1(t *testing.T) {
	var s string
	var err error
	b = game.NewBoard(3, 3)
	boardSquare := []string{"x", "-", "-", "x", "-", "-", "x", "-", "-"}
	err = b.Fill(boardSquare)
	if err != nil {
		fmt.Println(err)
	}
	// INIT OK

	s, err = b.ToString()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(s)
	}

	assert.True(t, b.IsWinByLineTopToBot(), "Should be true")

}

func TestWinByLaneTopToBot_2(t *testing.T) {
	var numberOfWin int = 0
	file, err := os.Open("./combination/all_combinations.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []string
		line = strings.Split(scanner.Text(), "")

		b = game.NewBoard(3, 3)
		err = b.Fill(line)
		if err != nil {
			fmt.Println(err)
		}
		// INIT OK

		if b.IsWinByLineTopToBot() {
			numberOfWin += 1
		}
		// s, err = b.ToString()
		// if err != nil {
		// 	fmt.Println(err)
		// } else {
		// 	fmt.Println(s)
		// }
	}
	fmt.Println("Number of win : ", numberOfWin)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
