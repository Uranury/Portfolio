package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

const settingsFile = "game_settings.txt"

func loadTries() int {
	data, err := os.ReadFile(settingsFile)
	if err != nil {
		return 0
	}
	tries, err := strconv.Atoi(string(data))
	if err != nil {
		return 0
	}
	return tries
}

func saveTries(tries int) error {
	return os.WriteFile(settingsFile, []byte(strconv.Itoa(tries)), 0644)
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "root",
		Short: "CLI-based number guessing game.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to the Number Guessing Game!")
			fmt.Println("I'm thinking of a number between 1 and 100.")
			fmt.Println("Use the 'select' command to choose a difficulty and 'start' to begin guessing!")
		},
	}

	selectCmd := &cobra.Command{
		Use:   "select",
		Short: "Selecting the diffulty of the game",
		Run: func(cmd *cobra.Command, args []string) {
			var NumberOfTries int
			switch args[0] {
			case "easy":
				NumberOfTries = 10
				fmt.Println("Difficulty set to EASY. You have 10 tries.")
			case "medium":
				NumberOfTries = 5
				fmt.Println("Difficulty set to MEDIUM. You have 5 tries.")
			case "hard":
				NumberOfTries = 3
				fmt.Println("Difficulty set to HARD. You have 3 tries.")
			default:
				fmt.Println("Invalid difficulty. Choose 'easy', 'medium', or 'hard'.")
			}

			if err := saveTries(NumberOfTries); err != nil {
				fmt.Printf("Error saving difficulty setting: %v\n", err)
				return
			}
		},
	}

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Starting the game.",
		Run: func(cmd *cobra.Command, args []string) {
			NumberOfTries := loadTries()

			if NumberOfTries == 0 {
				fmt.Print("Please select the difficulty using 'select' command first.")
				return
			}
			attempts := 0
			randomInt := rand.Intn(100) + 1

			for NumberOfTries > 0 {
				fmt.Printf("Enter your guess (tries left: %d): ", NumberOfTries)
				var input string
				fmt.Scanln(&input)

				val, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Invalid input. Please enter a number.")
					continue
				}

				if val < randomInt {
					fmt.Printf("Incorrect! The number is greater than %d.\n", val)
					NumberOfTries--
					attempts++
				} else if val > randomInt {
					fmt.Printf("Incorrect! The number is less than %d.\n", val)
					NumberOfTries--
					attempts++
				} else {
					fmt.Printf("Congratulations! You guessed the correct number in %d attempts.", attempts)
					os.Remove(settingsFile)
					return
				}
			}
			if NumberOfTries == 0 {
				fmt.Printf("Game over! The correct number was %d.\n", randomInt)
				os.Remove(settingsFile)
			}
		},
	}
	rootCmd.AddCommand(selectCmd)
	rootCmd.AddCommand(startCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
