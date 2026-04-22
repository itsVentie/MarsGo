package main

import (
	"bufio"
	"fmt"
	"marsgo/pkg/engine"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	game := engine.NewMarseilleGame()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		clearScreen()
		fmt.Println("=== MarsGo: Marseille Chess ===")
		fmt.Println(game.DrawBoard())

		fmt.Printf("Status: %s\n", game.GetStatus())
		fmt.Printf("Turn: %s | SubTurn: %v\n", game.Turn(), game.SubTurn())

		history := game.History()
		if len(history) > 0 {
			fmt.Printf("History: %s\n", strings.Join(history, " "))
		}

		fmt.Printf("Valid moves: %v\n", game.GetValidMoves())
		fmt.Print("\nEnter move or 'undo' (e.g. e2e4): ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(strings.ToLower(scanner.Text()))

		if input == "undo" {
			if err := game.Undo(); err != nil {
				fmt.Printf("\nError: %v. Press Enter...", err)
				scanner.Scan()
			}
			continue
		}

		if err := game.PushMove(input); err != nil {
			fmt.Printf("\nError: %v. Press Enter...", err)
			scanner.Scan()
			continue
		}

		if game.IsGameOver() {
			clearScreen()
			fmt.Println("=== FINAL POSITION ===")
			fmt.Println(game.DrawBoard())
			fmt.Printf("RESULT: %s\n", game.GetStatus())
			fmt.Println("Press Enter to exit...")
			scanner.Scan()
			break
		}
	}
}
