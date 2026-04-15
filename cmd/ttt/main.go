package main

import (
	"fmt"

	"ukiran.com/tictactoe/internal/board"
)

func main() {
	startGame()
}

func startGame() {
	players := []player{xPlayer, oPlayer}
	turns := 0
	gameBoard := board.NewBoard()

	for {
		currPlayer := players[turns%2]
		fmt.Println(playerState[currPlayer])
		gameBoard.Render()
		fmt.Printf("Player %v's turn...\n", playerState[currPlayer])

		move, err := GetMove()
		if err != nil {
			// TODO: Handle a "quit/q" input by GetMove
			fmt.Printf("Input error: %v. Please try again.\n", err)
			continue
		}

		gameBoard, err = makeMove(gameBoard, move, currPlayer)
		if err != nil {
			continue
		}

		if found, winner := getWinner(gameBoard); found {
			gameBoard.Render()
			fmt.Printf("GAME OVER: THE WINNER IS %q!\n", playerState[winner])
			break
		}

		if isBoardFull(gameBoard) {
			gameBoard.Render()
			fmt.Println("GAME OVER: IT'S A DRAW!")
			break
		}

		turns++
	}
}
