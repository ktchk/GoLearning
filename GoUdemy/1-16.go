package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
 +---+
 0   |
/|\  |
/ \  |
    ===

Secret Word : M_N___
Incorrect Guesses : A

Guess a Letter : Y

Sorry Your Dead! The word is MONKEY
YES the Secret Word is MONKEY

*/

var hmArr = [7]string{
	" +---+\n" +
		"     |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/    |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/    |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/ \\  |\n" +
		"    ===\n",
}

var wordArr = [7]string{
	"JAZZ",
	"ROBOCOP",
	"ZIPPER",
	"HOOKER",
	"PITCH",
	"BEER",
	"LETTER",
}

var randWord string
var correctLetters []string
var guessedLetters string
var wrongGuesses string
var encryptedWord string
var letter string

func getLetter() string {
	for letter = ""; true; {
		reader := bufio.NewReader(os.Stdin)
		inputRune, size, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		} else if size != 1 || inputRune == ' ' || inputRune == 13 {
			fmt.Println("WE NEED ONE LETTER!")
		} else {
			inputString := string(inputRune)
			input := strings.ToUpper(inputString)
			//fmt.Println(inputRune)
			letter = input
			break
		}
	}
	return letter
}

func main() {

	// Initialisation, Define random Word from array
	rand.Seed(time.Now().Unix())
	randNumber := rand.Intn(len(wordArr)) - 1
	randWord = wordArr[randNumber]
	correctLetters = strings.Split(randWord, "")
	for range correctLetters {
		encryptedWord += "_"
	}

	//Loop
	for wrongGuessesNum := 0; true; {

		//check if WIN
		if strings.Contains(encryptedWord, "_") != true {
			fmt.Println("YES the Secret Word is ", encryptedWord)
			break
		}

		//check if DEAD
		if wrongGuessesNum == 6 {
			fmt.Println(hmArr[wrongGuessesNum])
			fmt.Println("Secret Word : ", encryptedWord)
			fmt.Println("Incorrect Guesses : ", wrongGuesses)
			fmt.Println("Sorry Your Dead! The word is ", randWord)
			break
		}

		//Show Game Board
		fmt.Println(hmArr[wrongGuessesNum])
		fmt.Println("Secret Word : ", encryptedWord)
		fmt.Println("Incorrect Guesses : ", wrongGuesses)
		fmt.Printf("Guess a Letter : ")

		//Letter input

		letter = getLetter()

		//encrypt Word to ____ with Guessed letters
		encryptedWord = ""
		for _, x := range correctLetters {
			if x == letter || strings.Contains(guessedLetters, x) {
				encryptedWord += x
			} else {
				encryptedWord += "_"
			}
		}

		//Main logic
		if strings.Contains(randWord, letter) {
			guessedLetters += letter
		} else {
			wrongGuessesNum += 1
			wrongGuesses += " " + letter
		}

	}

}
