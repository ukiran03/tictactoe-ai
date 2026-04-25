package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ukiran.com/tictactoe/internal/board"
)

func main() {
	startGame()
}

func startGame() {
	scanner := bufio.NewScanner(os.Stdin)
	gameBoard := board.NewBoard()
	currPlayer := board.XState // Start with X

	for {
		gameBoard.Render()
		fmt.Printf("Player %s's turn...\n", currPlayer)

		quit, move, err := getMove(scanner)
		if quit {
			fmt.Println("Game terminated by user. Goodbye!")
			return // or break
		}

		if err != nil {
			fmt.Printf("Input error: %v. Please try again.\n", err)
			continue
		}

		newBoard, err := makeMove(gameBoard, move, currPlayer)
		if err != nil {
			fmt.Printf("Move error: %v\n", err)
			continue
		}

		gameBoard = newBoard

		if found, winner := getWinner(gameBoard); found {
			gameBoard.Render()
			fmt.Printf("GAME OVER: THE WINNER IS %q!\n", winner)
			break
		}

		if isBoardFull(gameBoard) {
			gameBoard.Render()
			fmt.Println("GAME OVER: IT'S A DRAW!")
			break
		}
		currPlayer = currPlayer.Next()
	}
}

func getMove(scanner *bufio.Scanner) (bool, *board.Vertex, error) {
	fmt.Print("Enter move (x,y) or 'q' to quit: ")

	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())

		// 1. Check for quit
		if strings.ToLower(input) == "q" {
			return true, nil, nil
		}

		// 2. Parse coordinates
		parts := strings.Split(input, ",")
		if len(parts) != 2 {
			fmt.Println("Invalid format. Please use x,y (e.g., 1,2):")
			continue // Keep looping until we get valid input or 'q'
		}

		x, errX := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, errY := strconv.Atoi(strings.TrimSpace(parts[1]))

		if errX != nil || errY != nil {
			fmt.Print("Invalid numbers. Please try again: ")
			continue
		}

		// 3. Return the vertex
		return false, &board.Vertex{X: x, Y: y}, nil
	}

	return true, nil, scanner.Err()
}
