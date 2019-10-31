package ai

import (
	"errors"
	"log"
	"strconv"

	"github.com/aurelienspnll/tic-tac-toe/src/game"
)

type Node struct {
	score        int
	playerTurn   bool
	currentBoard *game.Board
	children     []*Node
}

func NewNode(b *game.Board, score int, playerTurn bool) *Node {
	n := new(Node)
	n.score = score // will have value when eval()
	n.playerTurn = playerTurn
	n.currentBoard = b
	if !n.IsTerminal() {
		n.children = make([]*Node, n.currentBoard.RemainingMoves()) // check if it's win before allocate memory
	}
	return n
}

func (n *Node) ScoreNextMove(b *game.Board) int {
	if b.IsWin() && !n.playerTurn { // IA Wins
		return 1
	} else if b.IsWin() && n.playerTurn { // Player wins
		return -1
	} else if b.IsFull() { // No one wins
		return 0
	} else if n.playerTurn { // Non-terminal case
		return 100 // int(math.Inf(+1)) : overflow, 100 is enough
	} else { // Non-terminal case
		return -100 // int(math.Inf(-1))
	}
}

func (n *Node) GenerateAllNextMove(mark string) {
	remainingMoves := n.currentBoard.RemainingMoves()
	// res := make([]*string, remainingMoves)
	var b = make([]string, len(n.currentBoard.GetSquare()))
	for i := 0; i < remainingMoves; i++ {
		blank, err := n.currentBoard.GetBlankPos(i)
		if err != nil {
			log.Fatal(err)
		}
		err = CopyList(n.currentBoard.GetSquare(), b)
		if err != nil {
			log.Fatal(err)
		}
		b[blank] = mark
		// fmt.Println(b) // OK
		newBoard := game.NewBoard(n.currentBoard.GetSize(), n.currentBoard.GetNumberToWin())
		newBoard.Fill(b)
		score := n.ScoreNextMove(newBoard)
		// fmt.Println(score) // OK
		n.children[i] = NewNode(newBoard, score, !n.playerTurn)
	}
}

func (n *Node) IsTerminal() bool {
	return n.score == 1 || n.score == -1 || n.score == 0
}

func (n *Node) ToString() string {
	var res string
	res += "------ Node Description ------\n"
	res += "Score : " + strconv.Itoa(n.score) + "\n"
	res += "Player's turn ? " + strconv.FormatBool(n.playerTurn) + "\n"
	s, err := n.currentBoard.ToString()
	if err != nil {
		log.Fatal(err)
	}
	res += s + "------------------------------"
	return res
}

// Copy a into b
func CopyList(a []string, b []string) error {
	if len(a) != len(b) {
		return errors.New("Both lists must be the same length")
	}
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}
	return nil
}

func (n *Node) GetCurrentBoard() *game.Board {
	return n.currentBoard
}
