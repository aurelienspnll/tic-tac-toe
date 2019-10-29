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

const size = 3
const numberToWin = 3

func TestMain(m *testing.M) {
	b = game.NewBoard(size, numberToWin)
	code := m.Run()
	os.Exit(code)
}

func TestGetMark(t *testing.T) {
	var s string
	var err error
	b = game.NewBoard(size, numberToWin)
	boardSquare := []string{"x", "o", "x", "-", "x", "-", "o", "-", "o"}
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
	var res string
	res, _ = b.GetMark(0, 0)
	assert.Equal(t, "x", res, "Should be x")
	res, _ = b.GetMark(0, 1)
	assert.Equal(t, "o", res, "Should be o")
	res, _ = b.GetMark(0, 2)
	assert.Equal(t, "x", res, "Should be x")
	res, _ = b.GetMark(1, 0)
	assert.Equal(t, "-", res, "Should be -")
	res, _ = b.GetMark(1, 1)
	assert.Equal(t, "x", res, "Should be x")
	res, _ = b.GetMark(1, 2)
	assert.Equal(t, "-", res, "Should be -")
	res, _ = b.GetMark(2, 0)
	assert.Equal(t, "o", res, "Should be o")
	res, _ = b.GetMark(2, 1)
	assert.Equal(t, "-", res, "Should be -")
	res, _ = b.GetMark(2, 2)
	assert.Equal(t, "o", res, "Should be o")
	res, err = b.GetMark(0, 3)
	assert.NotNil(t, err, "err should be not nil")
	assert.Equal(t, "error", res, "Should be error")
}

func TestIsFull(t *testing.T) {
	var s string
	var err error
	b = game.NewBoard(size, numberToWin)
	boardSquare := []string{"x", "x", "x", "o", "-", "o", "o", "o", "x"}
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

	assert.False(t, b.IsFull(), "Should be false")
	b.Mark(1, 1, "x")
	s, err = b.ToString()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(s)
	}
	assert.True(t, b.IsFull(), "Should be true")
}

func TestWinByLineRightToLeft_1(t *testing.T) {
	var s string
	var err error
	b = game.NewBoard(size, numberToWin)
	boardSquare := []string{"x", "x", "x", "-", "-", "-", "-", "-", "-"}
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

func TestWinByLineRightToLeft_2(t *testing.T) {
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

		b = game.NewBoard(size, numberToWin)
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

func TestWinByLineTopToBot_1(t *testing.T) {
	var s string
	var err error
	b = game.NewBoard(size, numberToWin)
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

func TestWinByLineTopToBot_2(t *testing.T) {
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

		b = game.NewBoard(size, numberToWin)
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

func TestWinByDiagonalRightToLeft_1(t *testing.T) {
	var s string
	var err error
	b = game.NewBoard(size, numberToWin)
	boardSquare := []string{"x", "-", "-", "-", "x", "-", "-", "-", "x"}
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

	assert.True(t, b.IsWinByDiagonalRightToLeft(), "Should be true")

}

func TestWinByDiagonalRightToLeft_2(t *testing.T) {
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

		b = game.NewBoard(size, numberToWin)
		err = b.Fill(line)
		if err != nil {
			fmt.Println(err)
		}
		// INIT OK

		if b.IsWinByDiagonalRightToLeft() {
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

func TestWinByDiagonalLeftToRight_1(t *testing.T) {
	var s string
	var err error
	b = game.NewBoard(size, numberToWin)
	boardSquare := []string{"-", "-", "x", "-", "x", "-", "x", "-", "-"}
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

	assert.True(t, b.IsWinByDiagonalLeftToRight(), "Should be true")

}

func TestWinByDiagonalLeftToRight_2(t *testing.T) {
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

		b = game.NewBoard(size, numberToWin)
		err = b.Fill(line)
		if err != nil {
			fmt.Println(err)
		}
		// INIT OK

		if b.IsWinByDiagonalLeftToRight() {
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
