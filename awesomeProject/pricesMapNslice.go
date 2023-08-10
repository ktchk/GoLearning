package main

import "fmt"

func main() {

	price := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}

	for d := range price {
		if price[d] >= 500 {
			fmt.Println(d)
		}
	}

	purchase := []string{"хлеб", "буженина", "сыр", "огурцы"}
	summ := 0
	for _, x := range purchase {
		a, b := price[x]
		if b == true {
			summ += a
		}
	}

	fmt.Println(summ)

}
