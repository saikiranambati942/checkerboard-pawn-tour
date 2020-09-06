package main

import (
	"fmt"
	"math/rand"
)

const (
	N          = 8
	NumOfMoves = 8
)

var (
	moveNumber = 0
	board      [][]int
	pos        []int
	// Move pattern of pawn on basis of the change of x coordinates and y coordinates respectively
	MOVES      = [][]int{{0, -3}, {2, -2}, {3, 0}, {2, 2}, {0, 3}, {-2, 2}, {-3, 0}, {-2, -2}}
)

// initialPosition ... is a function that initializes the board slice and sets the initial position of a pawn
func initialPosition() []int {
	pos = make([]int, 2)
	pos[0] = rand.Intn(N)
	pos[1] = rand.Intn(N)
	board = make([][]int, N)
	for i := range board {
		board[i] = make([]int, N)
	}
	moveNumber++
	board[pos[0]][pos[1]] = moveNumber
	return pos
}

// nextMove ... is a function that decides the next position of a pawn
func nextMove(pos []int) []int {
	xPos := pos[0]
	yPos := pos[1]
	accessibility := NumOfMoves
	for i := 0; i < NumOfMoves; i++ {
		newX := xPos + MOVES[i][0]
		newY := yPos + MOVES[i][1]
		newAccessibility := getAccessibility(newX, newY)
		if isValid(newX, newY) && newAccessibility < accessibility {
			pos[0] = newX
			pos[1] = newY
			accessibility = newAccessibility
		}
	}
	moveNumber++
	board[pos[0]][pos[1]] = moveNumber
	return pos
}

// getAccessibility ... is a function that returns the accessibility of a pawn position
func getAccessibility(x, y int) int {
	accessibility := 0
	for i := 0; i < NumOfMoves; i++ {
		if isValid(x+MOVES[i][0], y+MOVES[i][1]) {
			accessibility++
		}
	}
	return accessibility
}
// isValid ... is a function to check valid indexes for N*N checker board
func isValid(x, y int) bool {
	if x < N && x >= 0 && y < N && y >= 0 && board[x][y] == 0 {
		return true
	}
	return false
}

// printResult ... is a function that prints the result matrix result[N][N]
func printResult() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Print(board[i][j])
			fmt.Print("\t")
		}
		fmt.Println()
	}
}

func main() {
	pos := initialPosition()
	for i := 1; i < N*N; i++ {
		nextMove(pos)
	}
	fmt.Println("The result is: ")
	printResult()

}
