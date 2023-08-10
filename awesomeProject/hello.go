package main

import "fmt"

func RemoveDuplicates(input []string) []string {

	output := make([]string, 0)
	inputSet := make(map[string]struct{}, len(input))
	for _, b := range input {
		if _, ok := inputSet[b]; !ok {
			output = append(output, b)
		}
		inputSet[b] = struct{}{}
	}
	return output
}

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	fmt.Println(RemoveDuplicates(input))
}
