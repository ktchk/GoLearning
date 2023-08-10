package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("What is your age : ")
	age, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	age = strings.TrimSpace(age)
	iAge, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	}

	if iAge < 5 {
		pl("Too young for school")
	} else if iAge == 5 {
		pl("Go to Kindergarten")
	} else if (iAge > 5) && (iAge <= 17) {
		pl("Go to School")
	} else {
		pl("Go to college")
	}
}
