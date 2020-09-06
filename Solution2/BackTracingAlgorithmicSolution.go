package main

import "fmt"

const N = 8

var result [][]int

// isValid ... is a function to check valid indexes for N*N checker board
func isValid(x, y int) bool {
	if x >= 0 && x < N && y >= 0 && y < N && result[x][y] == -1 {
		return true
	}
	return false
}

// printResult ... is a function that prints the result matrix result[N][N]
func printResult() {
	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			fmt.Print(result[x][y])
			fmt.Print("\t")
		}
		fmt.Println()
	}
}

// solvePawnUtil ... is a recursive utility function to solve Pawn Tour problem
func solvePawnUtil(x int, y int, movei int, xMove []int, yMove []int) bool {
	var k, nextX, nextY int
	if movei == N*N {
		return true
	}
	// Try all next moves from the current coordinate x,y
	for k = 0; k < N; k++ {
		nextX = x + xMove[k]
		nextY = y + yMove[k]
		if isValid(nextX, nextY) {
			result[nextX][nextY] = movei
			if solvePawnUtil(nextX, nextY, movei+1, xMove, yMove) {
				return true
			}
			result[nextX][nextY] = -1
		}
	}
	return false
}

/* solvePawnTour ... This function solves the Pawn Tour problem using Backtracking.
 This function mainly uses solvePawnUtil() to solve the problem. It returns false if no
 complete tour is possible, otherwise return true and prints the tour */
func solvePawnTour() bool {
	result = make([][]int, N)
	for i := range result {
		result[i] = make([]int, N)
	}

	// Initialization of result matrix
	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			result[x][y] = -1
		}
	}
	/* xMove and yMove define next move of Pawn.
	   xMove is for next value of x coordinate
	   yMove is for next value of y coordinate */
	xMove := []int{0, 2, 3, 2, 0, -2, -3, -2}
	yMove := []int{-3, -2, 0, 2, 3, 2, 0, -2}

	// Since the Pawn is initially at the first block
	result[0][0] = 0
	// Start from 0,0 and explore all tours using solvePawnUtil()
	if !solvePawnUtil(0, 0, 1, xMove, yMove) {
		fmt.Println("The solution doesn't exist")
		return false
	}
	printResult()
	return true
}

func main() {
	fmt.Println("The result is: ")
	solvePawnTour()
}
