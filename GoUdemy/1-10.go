package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func inputNum() (iNumber int) {
	reader := bufio.NewReader(os.Stdin)
	number, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	number = strings.TrimSpace(number)
	iNumber, err = strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func getRandomInt() (randomNumber int) {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randomNumber = rand.Intn(50) + 1
	return
}

func main() {

	solutionNum := getRandomInt()
	fmt.Println("Guess a number between 0 and 50")
	for userNum := inputNum(); ; userNum = inputNum() {
		if userNum < solutionNum {
			fmt.Println("Guess Something Higher")
		} else if userNum > solutionNum {
			fmt.Println("Guess Something Lower")
		} else {
			fmt.Println("Bingo!")
			break
		}
	}

}
