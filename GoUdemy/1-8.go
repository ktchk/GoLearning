package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func input() (iNumber int) {
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

func main() {
	fmt.Printf("Enter Number 1 : ")
	num1 := input()
	fmt.Printf("Enter Number 2 : ")
	num2 := input()
	fmt.Printf("%d + %d = %d", num1, num2, num1+num2)
}
