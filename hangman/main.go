package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)
var dictionary = []string{
	"Kazakhstan",
	"Manchester United",
	"Almaty",
	"Football",
	"You Got To Learn",
	"Satbayev University",
}

func main() {
	targetWord := getRandomWord()
	guessedLetters := initializeGuessedWord(targetWord)
	hangmanState := 0

	for !GameOver(targetWord, guessedLetters, hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("Invalid input. Please use letters only...")
			continue
		}

		letter := rune(input[0])
		if isCorrectGuess(targetWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}

	}

	fmt.Println("Game Over!")
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You Win!")
	} else if isHangmanComplete(hangmanState) {
		fmt.Println("You lose...")
	} else {
		panic("invalid state. Game is over and there is no winner")
	}
}

func getRandomWord() string {
	targetword := dictionary[rand.Intn(len(dictionary))]
	return targetword
}

func initializeGuessedWord(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func GameOver(targetWord string, guessedLetters map[rune]bool, hangmanState int) bool {
	return isWordGuessed(targetWord, guessedLetters) ||
		isHangmanComplete(hangmanState)
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}
	return true
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
}

func printGameState(targetWord string, guessedLetter map[rune]bool, hangmanState int) {
	fmt.Println(getGuessingProgress(targetWord, guessedLetter))
	fmt.Println()

	fmt.Println(getHangmanDrawing(hangmanState))
}

func getGuessingProgress(targetWord string, guessedLetter map[rune]bool) string {
	result := ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
		} else if guessedLetter[unicode.ToLower(ch)] == true {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}
		result += " "
	}

	return result
}

func getHangmanDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("states/hangman%v", hangmanState))
	if err != nil {
		panic(err)
	}

	return string(data)
}

func readInput() string {
	fmt.Println()
	fmt.Print("> ")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}
