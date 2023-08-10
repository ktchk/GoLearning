package main

import "fmt"

func main() {
	slice := make([]int, 100)
	for i := range slice {
		if i <= 100 {
			slice[i] = i + 1
		}
	}

	sliceTenPlusTen := append(slice[:10], slice[89:]...)
	sliceTPTreversed := make([]int, 10)

	for i, z := 0, len(sliceTenPlusTen)-1; i < z; i, z = i+1, z-1 {
		sliceTenPlusTen[i], sliceTenPlusTen[z] = sliceTenPlusTen[z], sliceTenPlusTen[i]
	}

	fmt.Println(slice)
	fmt.Println(sliceTenPlusTen)
	fmt.Println(sliceTPTreversed)
}
