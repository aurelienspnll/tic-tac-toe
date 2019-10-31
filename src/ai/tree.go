package ai

import (
	"fmt"
)

func GetOpennantMark(mark string) string {
	if mark == "o" {
		return "x"
	} else {
		return "o"
	}
}

func BuildTree(n *Node, depth int, mark string) {
	if depth == 0 || n.IsTerminal() {
		return
	} else {
		n.GenerateAllNextMove(mark)
		var scores []int
		for _, element := range n.children {
			BuildTree(element, depth-1, GetOpennantMark(mark))
			scores = append(scores, element.score) // All in the same loop = less complexity
		} // the tree below the node n was built
		// Now we evaluate the score with minimax
		if n.playerTurn {
			n.score = Min(scores)
		} else {
			n.score = Max(scores)
		}
	}
}

func Max(list []int) int {
	res := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] > res {
			res = list[i]
		}
	}
	return res
}

func Min(list []int) int {
	res := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] < res {
			res = list[i]
		}
	}
	return res
}

func IndexMax(list []int) int {
	score := list[0]
	res := 0
	for i := 1; i < len(list); i++ {
		if list[i] > score {
			score = list[i]
			res = i
		}
	}
	return res
}

// Assume that the tree has been generated
func PrintTree(n *Node) {
	var q []*Node    // Create a queue
	q = append(q, n) // Enqueue
	for len(q) > 0 {
		node := q[0]
		if node == nil {
			break
		}
		fmt.Println(node.ToString())
		for _, element := range node.children {
			q = append(q, element)
		}
		q = q[1:] // That don't remove but replace by nil...
	}
}

func GetPosBestMove(currentBoard []string, nextBoard []string, size int) (int, int) {
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if currentBoard[x*size+y] != nextBoard[x*size+y] {
				return x, y
			}
		}
	}
	return -1, -1 // Error
}

func BestMove(n *Node) (int, int) { // The tree was built so now we choose the best move
	// = The children with score = 1
	if n.GetCurrentBoard().IsEmpty() { // if board size = 3
		fmt.Println("The best move is : ", 1, 1)
		return 1, 1
	}
	var scores []int
	for _, element := range n.children {
		scores = append(scores, element.score)
	}
	best := n.children[IndexMax(scores)]
	posx, posy := GetPosBestMove(n.GetCurrentBoard().GetSquare(), best.GetCurrentBoard().GetSquare(), n.GetCurrentBoard().GetSize())
	fmt.Println("The best move is : ", posx, posy)
	return posx, posy
}
