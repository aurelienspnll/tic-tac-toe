package test

import (
	"testing"
	"github.com/aurelienspnll/tic-tac-toe/src/game"
)

func TestMain(m *testing.M) {
	var b = game.NewBoard()
	b.Mark(0, 0, "o")
}

func TestSum(t *testing.T) {
    total := 5 + 5 
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}